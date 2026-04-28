package main

import (
	"os"

	"github.com/arashrasoulzadeh/devenv/src/app"
	"github.com/arashrasoulzadeh/devenv/src/commands"
	"github.com/arashrasoulzadeh/devenv/src/config"
	"github.com/arashrasoulzadeh/devenv/src/log"
)

func main() {
	log.Start()

	// help command
	if len(os.Args) > 1 && os.Args[1] == "help" || len(os.Args) == 1 {
		commands.HelpCommand(os.Args)
	}

	// version command
	if len(os.Args) > 1 && os.Args[1] == "version" {
		commands.VersionCommand(os.Args)
	}

	// ---- load config instance ----
	c := config.New()

	if err := c.Load("config.toml"); err != nil {
		log.Fatal(err.Error())
	}

	// ---- create app ----
	r := app.New(c)
	r.OutputDir = "."

	// ---- run ----
	if err := r.Run(os.Args); err != nil {
		log.Fatal(err.Error())
	}
}
