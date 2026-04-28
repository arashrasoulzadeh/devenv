package main

import (
	"os"

	"github.com/arashrasoulzadeh/devenv/src/log"
)

func main() {
	log.Start()

	cli := ParseCLI(os.Args)

	if HandleBuiltins(cli) {
		return
	}

	r := mustInitApp(cli.ConfigFile)

	if err := Dispatch(cli, r); err != nil {
		log.Fatal(err.Error())
	}
}