package main

import (
	"os"

	"github.com/arashrasoulzadeh/devenv/src/app"
	"github.com/arashrasoulzadeh/devenv/src/commands"
)

func HandleBuiltins(cli CLI) bool {
	if len(cli.Args) == 0 || cli.Args[0] == "help" {
		commands.HelpCommand(os.Args)
		return true
	}

	if cli.Args[0] == "version" {
		commands.VersionCommand(os.Args)
		return true
	}

	return false
}

func Dispatch(cli CLI, r *app.Runner) error {
	switch cli.Args[0] {
	case "env":
		commands.EnvCommand(os.Args, r)
		return nil
	default:
		return r.Run(cli.Args)
	}
}
