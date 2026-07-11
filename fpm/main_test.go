//go:build !windows

package main

import (
	"context"
	"os"
	"path/filepath"
	"reflect"
	"sync"
	"syscall"
	"testing"
	"time"

	"github.com/pelletier/go-toml/v2"
)

func TestNormalizePoolSocketEndpoint(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		wantEndpoint string
		wantPath     string
		wantUnix     bool
		wantErr      bool
	}{
		{
			name:         "unix endpoint with prefix",
			input:        "unix:/tmp/client.sock",
			wantEndpoint: "unix:/tmp/client.sock",
			wantPath:     "/tmp/client.sock",
			wantUnix:     true,
		},
		{
			name:         "absolute unix path without prefix",
			input:        "/tmp/client.sock",
			wantEndpoint: "unix:/tmp/client.sock",
			wantPath:     "/tmp/client.sock",
			wantUnix:     true,
		},
		{
			name:         "relative unix path without prefix",
			input:        "./run/client.sock",
			wantEndpoint: "unix:./run/client.sock",
			wantPath:     "./run/client.sock",
			wantUnix:     true,
		},
		{
			name:         "tcp endpoint untouched",
			input:        "127.0.0.1:9100",
			wantEndpoint: "127.0.0.1:9100",
			wantPath:     "",
			wantUnix:     false,
		},
		{
			name:    "empty endpoint returns error",
			input:   "",
			wantErr: true,
		},
		{
			name:    "invalid unix prefix path returns error",
			input:   "unix:",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEndpoint, gotPath, gotUnix, err := normalizePoolSocketEndpoint(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error, got endpoint=%q path=%q isUnix=%t", gotEndpoint, gotPath, gotUnix)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if gotEndpoint != tt.wantEndpoint {
				t.Fatalf("endpoint mismatch: got %q want %q", gotEndpoint, tt.wantEndpoint)
			}
			if gotPath != tt.wantPath {
				t.Fatalf("socket path mismatch: got %q want %q", gotPath, tt.wantPath)
			}
			if gotUnix != tt.wantUnix {
				t.Fatalf("isUnix mismatch: got %t want %t", gotUnix, tt.wantUnix)
			}
		})
	}
}

func TestPoolConfigParsesGlobalASA(t *testing.T) {
	raw := []byte(`
site_name = "tenant-a"
uid = 1001
gid = 1001
socket = "unix:/var/run/axonasp/tenant-a.sock"
config_file = "/opt/axonasp/config/tenant-a.toml"
global_asa = "/srv/tenant-a"
app_path = "/srv/tenant-a"
memory_limit_mb = 256
max_restarts = 5
tmp_dir = "/opt/axonasp/temp"
`)

	var conf PoolConfig
	if err := toml.Unmarshal(raw, &conf); err != nil {
		t.Fatalf("failed to parse pool TOML: %v", err)
	}

	if conf.GlobalAsa != "/srv/tenant-a" {
		t.Fatalf("global_asa mismatch: got %q want %q", conf.GlobalAsa, "/srv/tenant-a")
	}
}

func TestSIGUSR2ReloadsOnlyModifiedPools(t *testing.T) {
	tmpDir := t.TempDir()

	poolAPath := filepath.Join(tmpDir, "site-a.conf")
	poolBPath := filepath.Join(tmpDir, "site-b.conf")

	if err := os.WriteFile(poolAPath, []byte("site_name='a'\nuid=1001\ngid=1001\nsocket='9001'\nconfig_file='/tmp/a.toml'\napp_path='/'\nmemory_limit_mb=64\nmax_restarts=0\n"), 0644); err != nil {
		t.Fatalf("failed to write pool A: %v", err)
	}
	if err := os.WriteFile(poolBPath, []byte("site_name='b'\nuid=1001\ngid=1001\nsocket='9002'\nconfig_file='/tmp/b.toml'\napp_path='/'\nmemory_limit_mb=64\nmax_restarts=0\n"), 0644); err != nil {
		t.Fatalf("failed to write pool B: %v", err)
	}

	originalConfigDir := configDir
	originalActivePools := activePools
	originalLauncher := launchPoolSupervisor

	configDir = tmpDir
	activePools = make(map[string]poolHandle)

	var mu sync.Mutex
	starts := make(map[string]int)
	cancels := make(map[string]int)

	launchPoolSupervisor = func(ctx context.Context, configPath string, done chan struct{}) {
		mu.Lock()
		starts[configPath]++
		mu.Unlock()

		go func() {
			<-ctx.Done()
			mu.Lock()
			cancels[configPath]++
			mu.Unlock()
			close(done)
		}()
	}

	t.Cleanup(func() {
		shutdownAllPools()
		configDir = originalConfigDir
		activePools = originalActivePools
		launchPoolSupervisor = originalLauncher
	})

	scanAndLoadConfigs()

	mu.Lock()
	if starts[poolAPath] != 1 || starts[poolBPath] != 1 {
		mu.Unlock()
		t.Fatalf("expected both pools to start once, got startsA=%d startsB=%d", starts[poolAPath], starts[poolBPath])
	}
	mu.Unlock()

	time.Sleep(2 * time.Millisecond)
	if err := os.WriteFile(poolAPath, []byte("site_name='a'\nuid=1001\ngid=1001\nsocket='9001'\nconfig_file='/tmp/a.toml'\napp_path='/'\nglobal_asa='/srv/a'\nmemory_limit_mb=64\nmax_restarts=0\n"), 0644); err != nil {
		t.Fatalf("failed to update pool A: %v", err)
	}

	if shouldStop := handleManagerSignal(syscall.SIGUSR2); shouldStop {
		t.Fatal("SIGUSR2 should not request manager shutdown")
	}

	mu.Lock()
	startsA := starts[poolAPath]
	startsB := starts[poolBPath]
	cancelsA := cancels[poolAPath]
	cancelsB := cancels[poolBPath]
	mu.Unlock()

	if startsA != 2 {
		t.Fatalf("expected modified pool A to restart once, got starts=%d", startsA)
	}
	if startsB != 1 {
		t.Fatalf("expected unmodified pool B to remain running, got starts=%d", startsB)
	}
	if cancelsA != 1 {
		t.Fatalf("expected modified pool A to receive one cancellation, got %d", cancelsA)
	}
	if cancelsB != 0 {
		t.Fatalf("expected unmodified pool B to not be cancelled, got %d", cancelsB)
	}
}

func TestBuildWorkerArgsIncludesOptionalFlags(t *testing.T) {
	tests := []struct {
		name         string
		conf         PoolConfig
		listen       string
		expectedArgs []string
	}{
		{
			name: "includes global_asa and pool name when set",
			conf: PoolConfig{
				SiteName:   "tenant-a",
				ConfigFile: "/opt/axonasp/config/tenant-a.toml",
				TmpDir:     "/opt/axonasp/temp",
				GlobalAsa:  "/srv/tenant-a",
			},
			listen: "unix:/var/run/axonasp/tenant-a.sock",
			expectedArgs: []string{
				"--fastcgi.server_port", "unix:/var/run/axonasp/tenant-a.sock",
				"--config.config_file", "/opt/axonasp/config/tenant-a.toml",
				"--global.temp_dir", "/opt/axonasp/temp",
				"--config.global_asa", "/srv/tenant-a",
				"--pool.name", "tenant-a",
			},
		},
		{
			name: "omits optional flags when values are empty",
			conf: PoolConfig{
				SiteName:   "",
				ConfigFile: "/opt/axonasp/config/default.toml",
				TmpDir:     "/opt/axonasp/temp",
				GlobalAsa:  "",
			},
			listen: "127.0.0.1:9100",
			expectedArgs: []string{
				"--fastcgi.server_port", "127.0.0.1:9100",
				"--config.config_file", "/opt/axonasp/config/default.toml",
				"--global.temp_dir", "/opt/axonasp/temp",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildWorkerArgs(tt.conf, tt.listen)
			if !reflect.DeepEqual(got, tt.expectedArgs) {
				t.Fatalf("args mismatch\n got:  %#v\n want: %#v", got, tt.expectedArgs)
			}
		})
	}
}
