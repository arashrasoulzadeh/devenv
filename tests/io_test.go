package tests

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/arashrasoulzadeh/devenv/src/io"
)

func TestSaveToFile_SavesContentToFile(t *testing.T) {
	tmpDir := t.TempDir()

	testFile := filepath.Join(tmpDir, "testfile.txt")
	content := "hello, io test!"

	if err := io.SaveToFile(testFile, content); err != nil {
		t.Fatalf("SaveToFile failed: %v", err)
	}

	got, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("failed to read file: %v", err)
	}

	if string(got) != content {
		t.Fatalf("content mismatch:\n got:  %q\n want: %q", got, content)
	}
}

func TestSaveToFile_HandlesInvalidPath(t *testing.T) {
	var invalidPath string

	if runtime.GOOS == "windows" {
		// Windows invalid path
		invalidPath = `Z:\invalid_path_that_should_not_exist_\file.txt`
	} else {
		// Unix invalid path
		invalidPath = "/this_path_should_not_exist_123456/file.txt"
	}

	err := io.SaveToFile(invalidPath, "data")

	if err == nil {
		t.Fatalf("expected error for invalid path, got nil")
	}
}
