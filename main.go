package main

import (
	"fmt"
	"os"

	"github.com/arashrasoulzadeh/devenv/src/app"
	"github.com/arashrasoulzadeh/devenv/src/config"
	"github.com/arashrasoulzadeh/devenv/src/log"
)

var (
	commit    string
	platform  string
	buildDate string
)

func main() {
	log.Start()

	// version command
	if len(os.Args) > 1 && os.Args[1] == "version" {
		if commit == "" {
			commit = "unknown"
		}
		if platform == "" {
			platform = "unknown"
		}
		if buildDate == "" {
			buildDate = "unknown"
		}

		fmt.Println("devenv")
		fmt.Println(" commit:    ", commit)
		fmt.Println(" platform:  ", platform)
		fmt.Println(" built at:  ", buildDate)

		os.Exit(0)
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
