package main

import (
	"os"

	"github.com/arashrasoulzadeh/devenv/src/app"
)

func main() {
	app.New("config.toml").Run(os.Args)
}
