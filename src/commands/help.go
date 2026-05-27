package commands

import (
	"github.com/arashrasoulzadeh/devenv/src/log"
	"os"
)

func HelpCommand(args []string) {
	helpText := `devenv - Deterministic Environment Configuration Manager

Usage:
  devenv [environment]

Description:
  devenv generates environment configuration files in multiple formats (dotenv, YAML, TOML)
  using a simple, layered TOML configuration.

Available Commands:
  help        Display this help message
  version     Show version and build information
  env         Check which config keys are missing from your OS environment

  Flags:
    --config FILE   Specify custom config file path
    --dont-commit   Dry-run: merge without writing the output file
    --no-values     Hide values in the missing-keys report (use with env)

Examples:
  devenv dev                    # Generate config using the 'dev' environment
  devenv env prod               # Report keys missing from OS environment
  devenv env prod --no-values   # Same, but hide values
  devenv help                   # Display this help message
  devenv version                # Print version information

For additional documentation and examples, visit:
  https://github.com/arashrasoulzadeh/devenv
`
	log.Print(helpText)
	os.Exit(0)
}
