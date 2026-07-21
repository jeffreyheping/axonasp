/*
 * AxonASP Server
 * Copyright (C) 2026 G3pix Ltda. All rights reserved.
 *
 * Developed by Lucas Guimarães - G3pix Ltda
 * Contact: https://g3pix.com.br
 * Project URL: https://g3pix.com.br/axonasp
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 *
 * Attribution Notice:
 * If this software is used in other projects, the name "AxonASP Server"
 * must be cited in the documentation or "About" section.
 *
 * Contribution Policy:
 * Modifications to the core source code of AxonASP Server must be
 * made available under this same license terms.
 */

// This file contains platform-agnostic unit tests for the FPM configuration
// types and helper functions. These tests are safe to run on all platforms,
// including Windows, since they only exercise pure data transformation logic.
package main

import (
	"reflect"
	"testing"

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

func TestPoolConfigParsesAllFields(t *testing.T) {
	raw := []byte(`
site_name = "tenant-a"
uid = 1001
gid = 1001
socket = "unix:/var/run/axonasp/tenant-a.sock"
config_file = "/opt/axonasp/config/tenant-a.toml"
global_asa_path = "/srv/tenant-a"
app_path = "/srv/tenant-a/wwwroot"
memory_limit_mb = 256
max_restarts = 5
tmp_dir = "/opt/axonasp/temp"
`)

	var conf PoolConfig
	if err := toml.Unmarshal(raw, &conf); err != nil {
		t.Fatalf("failed to parse pool TOML: %v", err)
	}

	if conf.SiteName != "tenant-a" {
		t.Fatalf("site_name mismatch: got %q want %q", conf.SiteName, "tenant-a")
	}
	if conf.UID != 1001 {
		t.Fatalf("uid mismatch: got %d want %d", conf.UID, 1001)
	}
	if conf.GID != 1001 {
		t.Fatalf("gid mismatch: got %d want %d", conf.GID, 1001)
	}
	if conf.GlobalAsa != "/srv/tenant-a" {
		t.Fatalf("global_asa_path mismatch: got %q want %q", conf.GlobalAsa, "/srv/tenant-a")
	}
	if conf.AppPath != "/srv/tenant-a/wwwroot" {
		t.Fatalf("app_path mismatch: got %q want %q", conf.AppPath, "/srv/tenant-a/wwwroot")
	}
	if conf.MemoryLimitMB != 256 {
		t.Fatalf("memory_limit_mb mismatch: got %d want %d", conf.MemoryLimitMB, 256)
	}
	if conf.MaxRestarts != 5 {
		t.Fatalf("max_restarts mismatch: got %d want %d", conf.MaxRestarts, 5)
	}
	if conf.TmpDir != "/opt/axonasp/temp" {
		t.Fatalf("tmp_dir mismatch: got %q want %q", conf.TmpDir, "/opt/axonasp/temp")
	}
}

func TestBuildWorkerArgsIncludesServerWebRoot(t *testing.T) {
	tests := []struct {
		name         string
		conf         PoolConfig
		listen       string
		expectedArgs []string
	}{
		{
			name: "all fields populated including app_path",
			conf: PoolConfig{
				SiteName:   "tenant-a",
				ConfigFile: "/opt/axonasp/config/tenant-a.toml",
				TmpDir:     "/opt/axonasp/temp",
				GlobalAsa:  "/srv/tenant-a",
				AppPath:    "/srv/tenant-a/wwwroot",
			},
			listen: "unix:/var/run/axonasp/tenant-a.sock",
			expectedArgs: []string{
				"--fastcgi.server_port", "unix:/var/run/axonasp/tenant-a.sock",
				"--config.config_file", "/opt/axonasp/config/tenant-a.toml",
				"--global.temp_dir", "/opt/axonasp/temp",
				"--config.global_asa", "/srv/tenant-a",
				"--pool.name", "tenant-a",
				"--server.web_root", "/srv/tenant-a/wwwroot",
			},
		},
		{
			name: "omits optional flags when values are empty but app_path is set",
			conf: PoolConfig{
				SiteName:   "",
				ConfigFile: "/opt/axonasp/config/default.toml",
				TmpDir:     "/opt/axonasp/temp",
				GlobalAsa:  "",
				AppPath:    "/var/www/default/wwwroot",
			},
			listen: "127.0.0.1:9100",
			expectedArgs: []string{
				"--fastcgi.server_port", "127.0.0.1:9100",
				"--config.config_file", "/opt/axonasp/config/default.toml",
				"--global.temp_dir", "/opt/axonasp/temp",
				"--server.web_root", "/var/www/default/wwwroot",
			},
		},
		{
			name: "app_path is always included when set even without other optionals",
			conf: PoolConfig{
				SiteName:   "",
				ConfigFile: "/opt/axonasp/config/minimal.toml",
				TmpDir:     "/opt/axonasp/temp",
				GlobalAsa:  "",
				AppPath:    "/home/user/site",
			},
			listen: "9000",
			expectedArgs: []string{
				"--fastcgi.server_port", "9000",
				"--config.config_file", "/opt/axonasp/config/minimal.toml",
				"--global.temp_dir", "/opt/axonasp/temp",
				"--server.web_root", "/home/user/site",
			},
		},
		{
			name: "app_path empty results in no --server.web_root flag",
			conf: PoolConfig{
				SiteName:   "no-webroot",
				ConfigFile: "/opt/axonasp/config/no-webroot.toml",
				TmpDir:     "/opt/axonasp/temp",
				GlobalAsa:  "",
				AppPath:    "",
			},
			listen: "9100",
			expectedArgs: []string{
				"--fastcgi.server_port", "9100",
				"--config.config_file", "/opt/axonasp/config/no-webroot.toml",
				"--global.temp_dir", "/opt/axonasp/temp",
				"--pool.name", "no-webroot",
			},
		},
		{
			name: "app_path with trailing whitespace is trimmed",
			conf: PoolConfig{
				SiteName:   "trimmed",
				ConfigFile: "/opt/axonasp/config/trimmed.toml",
				TmpDir:     "/opt/axonasp/temp",
				GlobalAsa:  "   ",
				AppPath:    "  /var/www/trimmed-site  ",
			},
			listen: "127.0.0.1:9100",
			expectedArgs: []string{
				"--fastcgi.server_port", "127.0.0.1:9100",
				"--config.config_file", "/opt/axonasp/config/trimmed.toml",
				"--global.temp_dir", "/opt/axonasp/temp",
				"--pool.name", "trimmed",
				"--server.web_root", "/var/www/trimmed-site",
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

// TestBuildWorkerArgsOrdering verifies that required flags always come first,
// followed by optional flags in a deterministic order (global_asa, pool.name,
// server.web_root).
func TestBuildWorkerArgsOrdering(t *testing.T) {
	conf := PoolConfig{
		SiteName:   "order-test",
		ConfigFile: "/cfg/order.toml",
		TmpDir:     "/tmp/order",
		GlobalAsa:  "/asa/order",
		AppPath:    "/app/order",
	}
	args := buildWorkerArgs(conf, "127.0.0.1:9000")

	// Required flags must come first, in order.
	required := []string{
		"--fastcgi.server_port",
		"--config.config_file",
		"--global.temp_dir",
	}
	optional := []string{
		"--config.global_asa",
		"--pool.name",
		"--server.web_root",
	}

	requiredIdx := 0
	optionalIdx := 0
	seenOptional := false

	for i := 0; i < len(args); i++ {
		if requiredIdx < len(required) && args[i] == required[requiredIdx] {
			if seenOptional {
				t.Fatalf("required flag %s appears after optional flags at position %d", required[requiredIdx], i)
			}
			i++ // skip value
			requiredIdx++
			continue
		}
		if optionalIdx < len(optional) && args[i] == optional[optionalIdx] {
			seenOptional = true
			i++ // skip value
			optionalIdx++
			continue
		}
	}

	if requiredIdx != len(required) {
		t.Fatalf("missing required flags: %v", required[requiredIdx:])
	}
	if optionalIdx != len(optional) {
		t.Fatalf("missing optional flags: %v", optional[optionalIdx:])
	}
}
