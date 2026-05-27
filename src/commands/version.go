package commands

import (
	"os"

	"github.com/arashrasoulzadeh/devenv/src/log"
)

var (
	commit    string
	platform  string
	buildDate string
)

func VersionCommand(args []string) {
	if commit == "" {
		commit = "unknown"
	}
	if platform == "" {
		platform = "unknown"
	}
	if buildDate == "" {
		buildDate = "unknown"
	}

	log.Print("devenv")
	log.Print(" commit:    ", commit)
	log.Print(" platform:  ", platform)
	log.Print(" built at:  ", buildDate)

	os.Exit(0)
}
