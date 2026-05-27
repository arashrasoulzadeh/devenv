package tests

import (
	"path/filepath"
	"testing"

	"github.com/arashrasoulzadeh/devenv/src/app"
	"github.com/arashrasoulzadeh/devenv/src/commands"
	"github.com/arashrasoulzadeh/devenv/src/config"
	"github.com/arashrasoulzadeh/devenv/src/consts"
)

func TestApp_Run_Env_Command(t *testing.T) {
	dir := t.TempDir()

	configPath := filepath.Join(dir, "config.toml")

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

	c := config.New()
	if err := c.Load(configPath); err != nil {
		t.Fatalf("config load failed: %v", err)
	}

	r := app.New(c)
	r.OutputDir = dir

	r.Run([]string{consts.DontCommitFlag})

	// read os ENV
	// env := os.Getenv()

	// run command
	commands.EnvCommand([]string{}, r)
}
