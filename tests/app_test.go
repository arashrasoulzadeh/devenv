package tests

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/arashrasoulzadeh/devenv/src/app"
	"github.com/arashrasoulzadeh/devenv/src/config"
)

func readFile(t *testing.T, path string) string {
	t.Helper()

	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read file failed: %v", err)
	}

	return string(b)
}

func TestApp_Run_DefaultEnv(t *testing.T) {
	dir := t.TempDir()

	configPath := filepath.Join(dir, "config.toml")
	outputPath := filepath.Join(dir, ".env")

	writeFile(t, configPath, `
[base]
HOST = "localhost"
PORT = 8080

[output]
name = ".env"
type = "dotenv"

[development]
DEBUG = true
`)

	// 🔥 NEW: instance-based config
	c := config.New()
	if err := c.Load(configPath); err != nil {
		t.Fatalf("config load failed: %v", err)
	}

	r := app.New(c)
	r.OutputDir = dir

	r.Run([]string{"cmd"})

	got := readFile(t, outputPath)

	if !strings.Contains(got, "HOST=") {
		t.Fatalf("missing HOST, got:\n%s", got)
	}

	if !strings.Contains(got, "PORT=8080") {
		t.Fatalf("missing PORT=8080, got:\n%s", got)
	}
}

func TestApp_Run_WithEnvOverlay(t *testing.T) {
	dir := t.TempDir()

	configPath := filepath.Join(dir, "config.toml")
	outputPath := filepath.Join(dir, "prod.env")

	writeFile(t, configPath, `
[base]
HOST = "localhost"
PORT = 8080

[output]
name = "prod.env"
type = "dotenv"

[production]
PORT = 80
DEBUG = false
`)

	// 🔥 NEW: instance-based config
	c := config.New()
	if err := c.Load(configPath); err != nil {
		t.Fatalf("config load failed: %v", err)
	}

	r := app.New(c)
	r.OutputDir = dir

	r.Run([]string{"cmd", "production"})

	got := readFile(t, outputPath)

	if !strings.Contains(got, "PORT=80") {
		t.Fatalf("expected PORT=80, got:\n%s", got)
	}

	if !strings.Contains(got, "DEBUG=false") {
		t.Fatalf("expected DEBUG=false, got:\n%s", got)
	}
}
