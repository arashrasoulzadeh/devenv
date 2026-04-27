package app

import (
	"fmt"
	"path/filepath"

	"github.com/arashrasoulzadeh/devenv/src/config"
	"github.com/arashrasoulzadeh/devenv/src/io"
	"github.com/arashrasoulzadeh/devenv/src/log"
	"github.com/arashrasoulzadeh/devenv/src/renderer"
)

type Runner struct {
	ConfigPath string
	OutputDir  string
}

func New(configPath string) *Runner {
	return &Runner{
		ConfigPath: configPath,
		OutputDir:  ".",
	}
}

func (r *Runner) Run(args []string) {
	log.Start()

	config.Parse(r.ConfigPath)

	raw := config.Get()
	if raw == nil {
		log.Fatal("config is nil after parsing")
	}

	base := map[string]any{}
	outputMeta := map[string]any{}
	variants := map[string]map[string]any{}

	// ---- split config ----
	for k, v := range raw {
		switch k {
		case "base":
			for k2, v2 := range v {
				base[k2] = v2
			}
		case "output":
			for k2, v2 := range v {
				outputMeta[k2] = v2
			}
		default:
			variants[k] = v
		}
	}

	// ---- env selection ----
	env := ""
	if len(args) > 1 && args[1] != "" {
		env = args[1]

		if _, ok := variants[env]; !ok {
			log.Fatal(fmt.Sprintf("environment '%s' not found", env))
		}

		log.Info("switching to", env)
	}

	// ---- merge config ----
	final := make(map[string]any, len(base))
	for k, v := range base {
		final[k] = v
	}

	if env != "" {
		for k, v := range variants[env] {
			final[k] = v
		}
	}

	// ---- output config ----
	outputName := ".env"
	outputType := "dotenv"

	if v, ok := outputMeta["name"].(string); ok && v != "" {
		outputName = v
	}

	if v, ok := outputMeta["type"].(string); ok && v != "" {
		outputType = v
	}

	if outputType != "dotenv" {
		log.Fatal("only dotenv output type is supported")
	}

	// ---- render ----
	formatted := renderer.ParseDotEnv(final)

	fullPath := filepath.Join(r.OutputDir, outputName)

	if err := io.SaveToFile(fullPath, formatted); err != nil {
		log.Fatal(fmt.Sprintf("failed to write output file: %v", err))
	}

	log.Info("generated file:", fullPath)
}
