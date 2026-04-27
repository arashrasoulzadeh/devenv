package tests

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/arashrasoulzadeh/devenv/src/io"
)

func TestSaveToFile_SavesContentToFile(t *testing.T) {
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "testfile.txt")
	content := "hello, io test!"

	err := io.SaveToFile(testFile, content)
	if err != nil {
		t.Fatalf("SaveToFile failed: %v", err)
	}

	data, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	if string(data) != content {
		t.Errorf("File contents mismatch: got %q, want %q", string(data), content)
	}
}

func TestSaveToFile_HandlesInvalidPath(t *testing.T) {
	// Use an invalid path (directory does not exist).
	invalidPath := filepath.Join("does_not_exist_dir", "file.txt")
	err := io.SaveToFile(invalidPath, "data")
	if err == nil {
		t.Error("Expected error for invalid path, got nil")
	}
}
