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
	cfg       *config.Config
	OutputDir string
}

func New(cfg *config.Config) *Runner {
	return &Runner{
		cfg:       cfg,
		OutputDir: ".",
	}
}

func (r *Runner) Run(args []string) error {
	log.Start()

	raw := r.cfg.Get()
	if raw == nil {
		return fmt.Errorf("config is nil")
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
			return fmt.Errorf("environment '%s' not found", env)
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
		return fmt.Errorf("only dotenv output type is supported")
	}

	// ---- render ----
	formatted := renderer.ParseDotEnv(final)

	fullPath := filepath.Join(r.OutputDir, outputName)

	if err := io.SaveToFile(fullPath, formatted); err != nil {
		return fmt.Errorf("failed to write output file: %w", err)
	}

	log.Info("generated file:", fullPath)

	return nil
}
