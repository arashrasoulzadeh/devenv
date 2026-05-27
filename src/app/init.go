package app

import (
	"github.com/arashrasoulzadeh/devenv/src/config"
	"github.com/arashrasoulzadeh/devenv/src/log"
)

func MustInitApp(configFile string) *Runner {
	c := config.New()

	if err := c.Load(configFile); err != nil {
		log.Fatal(err.Error())
	}

	r := New(c)
	r.OutputDir = "."

	return r
}
