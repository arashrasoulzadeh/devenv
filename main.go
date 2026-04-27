package main

import (
	"fmt"
	"os"

	"github.com/arashrasoulzadeh/devenv/src/config"
	"github.com/arashrasoulzadeh/devenv/src/io"
	"github.com/arashrasoulzadeh/devenv/src/log"
	renderer "github.com/arashrasoulzadeh/devenv/src/renderer"
)

func main() {
	log.Start()

	config.Parse("config.toml")

	rawConfigs := config.Get()
	if rawConfigs == nil {
		log.Fatal("Config is nil after parsing config.toml")
	}

	baseConfig := make(map[string]any)
	outputConfigs := make(map[string]any)
	variantConfigs := make(map[string]map[string]any)

	for k, v := range rawConfigs {
		switch k {
		case "base":
			for k2, v2 := range v {
				baseConfig[k2] = v2
			}
		case "output":
			for k2, v2 := range v {
				outputConfigs[k2] = v2
			}
		default:
			variantConfigs[k] = v
		}
	}
	env := ""
	if len(os.Args) > 1 && os.Args[1] != "" {
		env = os.Args[1]
		if _, ok := variantConfigs[env]; !ok {
			log.Fatal(fmt.Sprintf("environment '%s' not found in config\n", env))
		}
		log.Info("switching to", env)
	}

	// Start output as a copy of baseConfig
	output := make(map[string]any, len(baseConfig))
	for k, v := range baseConfig {
		output[k] = v
	}
	// Overlay variant config if applicable
	if env != "" {
		for k, v := range variantConfigs[env] {
			output[k] = v
		}
	}

	outputName, outputType := ".env", "dotenv"
	if name, ok := outputConfigs["name"].(string); ok && name != "" {
		outputName = name
	} else {
		log.Info("output name is missing or not a string, defaulting to .env")
	}
	if typ, ok := outputConfigs["type"].(string); ok && typ != "" {
		outputType = typ
	} else {
		log.Info("output type is missing or not a string, defaulting to dotenv")
	}

	if outputType == "dotenv" {
		formatted := renderer.ParseDotEnv(output)
		io.SaveToFile(outputName, formatted)

	}

}
