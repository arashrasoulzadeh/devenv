package main

import "strings"

type CLI struct {
	Args       []string
	ConfigFile string
}

func ParseCLI(argv []string) CLI {
	args := append([]string{}, argv[1:]...)
	configFile := "config.toml"

	for i := 0; i < len(args); {
		switch {
		case args[i] == "--config" && i+1 < len(args):
			configFile = args[i+1]
			args = append(args[:i], args[i+2:]...)

		case strings.HasPrefix(args[i], "--config="):
			configFile = strings.TrimPrefix(args[i], "--config=")
			args = append(args[:i], args[i+1:]...)

		default:
			i++
		}
	}

	return CLI{
		Args:       args,
		ConfigFile: configFile,
	}
}
