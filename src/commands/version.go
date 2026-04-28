package commands

import (
	"fmt"
	"os"
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

	fmt.Println("devenv")
	fmt.Println(" commit:    ", commit)
	fmt.Println(" platform:  ", platform)
	fmt.Println(" built at:  ", buildDate)

	os.Exit(0)
}
