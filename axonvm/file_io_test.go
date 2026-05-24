package axonvm

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestFileIOBasic(t *testing.T) {
	tempDir := t.TempDir()
	testFile1 := filepath.Join(tempDir, "test1.txt")
	testFile2 := filepath.Join(tempDir, "test2.txt")

	code := `
		Open "` + filepath.ToSlash(testFile1) + `" For Output As #1
		Print #1, "Hello World"
		Close #1
		
		Dim s1, s2
		Open "` + filepath.ToSlash(testFile1) + `" For Input As #1
		Line Input #1, s1
		Close #1
		
		Open "` + filepath.ToSlash(testFile2) + `" For Output As #2
		Print #2, "Another File"
		Close #2
		
		Open "` + filepath.ToSlash(testFile2) + `" For Input As #2
		Line Input #2, s2
		Close #2
	`

	compiler := NewCompiler(code)
	err := compiler.Compile()
	if err != nil {
		t.Fatalf("Compilation failed: %v", err)
	}

	s1Idx, _ := compiler.Globals.Get("s1")
	s2Idx, _ := compiler.Globals.Get("s2")

	vm := NewVM(compiler.Bytecode(), compiler.Constants(), compiler.GlobalsCount())
	vm.SetExecutionMode(ExecutionModeCLI)

	err = vm.Run()
	if err != nil {
		t.Fatalf("VM run failed: %v", err)
	}

	s1 := vm.Globals[s1Idx]
	s2 := vm.Globals[s2Idx]

	if s1.String() != "Hello World" {
		t.Errorf("Expected 'Hello World', got '%s'", s1.String())
	}
	if s2.String() != "Another File" {
		t.Errorf("Expected 'Another File', got '%s'", s2.String())
	}
}

func TestFileIORESTRICTION(t *testing.T) {
	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "test_res.txt")

	code := `
		Open "` + filepath.ToSlash(testFile) + `" For Output As #1
		Print #1, "Denied"
		Close #1
	`

	compiler := NewCompiler(code)
	err := compiler.Compile()
	if err != nil {
		t.Fatalf("Compilation failed: %v", err)
	}
	vm := NewVM(compiler.Bytecode(), compiler.Constants(), compiler.GlobalsCount())
	vm.SetExecutionMode(ExecutionModeServer) // Should fail

	err = vm.Run()
	if err == nil {
		t.Fatal("Expected error for File I/O in Server mode, but got nil")
	}
	if !strings.Contains(err.Error(), "restricted to CLI environment only") {
		t.Errorf("Expected restriction error message, got: %v", err)
	}
}

func TestFileIOCleanup(t *testing.T) {
	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "test_cleanup.txt")

	code := `
		Open "` + filepath.ToSlash(testFile) + `" For Output As #1
		Print #1, "Cleanup Test"
		' No Close
	`

	compiler := NewCompiler(code)
	err := compiler.Compile()
	if err != nil {
		t.Fatalf("Compilation failed: %v", err)
	}
	vm := NewVM(compiler.Bytecode(), compiler.Constants(), compiler.GlobalsCount())
	vm.SetExecutionMode(ExecutionModeCLI)

	err = vm.Run()
	if err != nil {
		t.Fatalf("VM run failed: %v", err)
	}

	if len(vm.fileIOItems) != 0 {
		t.Errorf("Expected 0 open files after Run, got %d", len(vm.fileIOItems))
	}

	// Check if file exists and can be opened by OS (meaning it was closed)
	_, err = os.Stat(testFile)
	if err != nil {
		t.Errorf("File should exist: %v", err)
	}
}
