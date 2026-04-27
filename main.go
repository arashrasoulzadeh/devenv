package main

import (
	"os"

	"github.com/arashrasoulzadeh/devenv/src/app"
	"github.com/arashrasoulzadeh/devenv/src/config"
	"github.com/arashrasoulzadeh/devenv/src/log"
)

func main() {
	log.Start()

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
