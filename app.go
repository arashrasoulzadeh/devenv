package main

import (
	"github.com/arashrasoulzadeh/devenv/src/app"
	"github.com/arashrasoulzadeh/devenv/src/config"
	"github.com/arashrasoulzadeh/devenv/src/log"
)

func mustInitApp(configFile string) *app.Runner {
	c := config.New()

	if err := c.Load(configFile); err != nil {
		log.Fatal(err.Error())
	}

	r := app.New(c)
	r.OutputDir = "."

	return r
}
