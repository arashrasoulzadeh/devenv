package tests

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/arashrasoulzadeh/devenv/src/app"
	"github.com/arashrasoulzadeh/devenv/src/commands"
	"github.com/arashrasoulzadeh/devenv/src/config"
	"github.com/arashrasoulzadeh/devenv/src/log"
)

// Keys are prefixed with DEVENV_XTEST_ to avoid colliding with real OS env vars.
const envTestConfig = `
[base]
DEVENV_XTEST_HOST = "localhost"
DEVENV_XTEST_PORT = 8080

[output]
name = ".env"
type = "dotenv"

[development]
DEVENV_XTEST_DEBUG = true
`

func newEnvTestRunner(t *testing.T) *app.Runner {
	t.Helper()

	dir := t.TempDir()
	configPath := filepath.Join(dir, "config.toml")
	writeFile(t, configPath, envTestConfig)

	c := config.New()
	if err := c.Load(configPath); err != nil {
		t.Fatalf("config load failed: %v", err)
	}

	r := app.New(c)
	r.OutputDir = dir
	return r
}

func captureLog(t *testing.T) *bytes.Buffer {
	t.Helper()
	var buf bytes.Buffer
	log.SetOutput(&buf)
	t.Cleanup(func() { log.SetOutput(os.Stdout) })
	return &buf
}

func TestEnvCommand_MissingKeysAreReported(t *testing.T) {
	buf := captureLog(t)
	r := newEnvTestRunner(t)

	commands.EnvCommand([]string{"devenv", "env", "development"}, r)

	out := buf.String()
	for _, key := range []string{"DEVENV_XTEST_HOST", "DEVENV_XTEST_PORT", "DEVENV_XTEST_DEBUG"} {
		if !strings.Contains(out, key) {
			t.Errorf("expected %s to be reported as missing, got:\n%s", key, out)
		}
	}
}

func TestEnvCommand_MissingKeyShowsValue(t *testing.T) {
	buf := captureLog(t)
	r := newEnvTestRunner(t)

	commands.EnvCommand([]string{"devenv", "env", "development"}, r)

	out := buf.String()
	if !strings.Contains(out, "localhost") {
		t.Errorf("expected value 'localhost' in output, got:\n%s", out)
	}
}

func TestEnvCommand_NoValues_HidesValues(t *testing.T) {
	buf := captureLog(t)
	r := newEnvTestRunner(t)

	commands.EnvCommand([]string{"devenv", "env", "development", "--no-values"}, r)

	out := buf.String()
	if strings.Contains(out, "localhost") {
		t.Errorf("expected value 'localhost' to be hidden with --no-values, got:\n%s", out)
	}
	if !strings.Contains(out, "DEVENV_XTEST_HOST") {
		t.Errorf("expected key name to still appear with --no-values, got:\n%s", out)
	}
}

func TestEnvCommand_PresentKey_NotReported(t *testing.T) {
	t.Setenv("DEVENV_XTEST_HOST", "some-value")
	buf := captureLog(t)
	r := newEnvTestRunner(t)

	commands.EnvCommand([]string{"devenv", "env", "development"}, r)

	out := buf.String()
	if strings.Contains(out, "DEVENV_XTEST_HOST") {
		t.Errorf("expected present key DEVENV_XTEST_HOST to not be reported, got:\n%s", out)
	}
}
