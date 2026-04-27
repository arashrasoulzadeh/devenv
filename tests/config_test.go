package tests

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/arashrasoulzadeh/devenv/src/config"
)

func writeFile(t *testing.T, path, content string) {
	t.Helper()

	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write file: %v", err)
	}
}

func mustGetSection(t *testing.T, cfg map[string]map[string]any, name string) map[string]any {
	t.Helper()

	section, ok := cfg[name]
	if !ok {
		t.Fatalf("missing section [%s]", name)
	}
	return section
}

func TestReadConfigFile(t *testing.T) {
	content := `
[base]
HOST = "localhost"
PORT = 8080

[output]
name = ".env.test"
type = "dotenv"

[development]
DEBUG = true

[production]
DEBUG = false
PORT = 80
`

	tmpFile := filepath.Join(t.TempDir(), "config.toml")
	writeFile(t, tmpFile, content)

	config.Parse(tmpFile)
	cfg := config.Get()

	// ---- base ----
	base := mustGetSection(t, cfg, "base")

	if base["HOST"] != "localhost" {
		t.Fatalf("base.HOST = %v, want localhost", base["HOST"])
	}
	if base["PORT"] != int64(8080) {
		t.Fatalf("base.PORT = %v, want 8080", base["PORT"])
	}

	// ---- output ----
	output := mustGetSection(t, cfg, "output")

	if output["name"] != ".env.test" {
		t.Fatalf("output.name = %v, want .env.test", output["name"])
	}
	if output["type"] != "dotenv" {
		t.Fatalf("output.type = %v, want dotenv", output["type"])
	}

	// ---- development ----
	dev := mustGetSection(t, cfg, "development")

	if dev["DEBUG"] != true {
		t.Fatalf("development.DEBUG = %v, want true", dev["DEBUG"])
	}

	// ---- production ----
	prod := mustGetSection(t, cfg, "production")

	if prod["DEBUG"] != false {
		t.Fatalf("production.DEBUG = %v, want false", prod["DEBUG"])
	}
	if prod["PORT"] != int64(80) {
		t.Fatalf("production.PORT = %v, want 80", prod["PORT"])
	}
}

func TestMergeConfigOverlay(t *testing.T) {
	base := map[string]any{
		"HOST":  "localhost",
		"PORT":  int64(8000),
		"DEBUG": false,
	}

	variant := map[string]any{
		"PORT":  int64(8081),
		"DEBUG": true,
	}

	merged := make(map[string]any, len(base))

	for k, v := range base {
		merged[k] = v
	}
	for k, v := range variant {
		merged[k] = v
	}

	if merged["HOST"] != "localhost" {
		t.Fatalf("HOST mismatch: %v", merged["HOST"])
	}
	if merged["PORT"] != int64(8081) {
		t.Fatalf("PORT mismatch: %v", merged["PORT"])
	}
	if merged["DEBUG"] != true {
		t.Fatalf("DEBUG mismatch: %v", merged["DEBUG"])
	}
}
