package axonvm

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/shirou/gopsutil/v3/disk"
)

// TestFSODiskUsageMatchesGopsutil verifies the FSO disk usage adapter preserves
// the underlying filesystem totals exposed by gopsutil.
func TestFSODiskUsageMatchesGopsutil(t *testing.T) {
	t.Parallel()

	rootPath := string(os.PathSeparator)
	if runtime.GOOS == "windows" {
		cwd, err := os.Getwd()
		if err != nil {
			t.Fatalf("get cwd: %v", err)
		}
		volume := filepath.VolumeName(cwd)
		if volume == "" {
			t.Skip("no drive volume available for disk usage test")
		}
		rootPath = volume + string(os.PathSeparator)
	}

	expected, err := disk.Usage(rootPath)
	if err != nil {
		t.Fatalf("disk usage for %q: %v", rootPath, err)
	}

	usage := (&VM{}).fsoDiskUsage(rootPath)
	if usage == nil {
		t.Fatalf("expected disk usage state for %q", rootPath)
	}

	if got := usage.Size(); got != expected.Total {
		t.Fatalf("total mismatch for %q: got %d want %d", rootPath, got, expected.Total)
	}
	if got := usage.Free(); got != expected.Free {
		t.Fatalf("free mismatch for %q: got %d want %d", rootPath, got, expected.Free)
	}
	if got := usage.Available(); got != expected.Free {
		t.Fatalf("available mismatch for %q: got %d want %d", rootPath, got, expected.Free)
	}
}
