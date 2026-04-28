package commands

import (
	"fmt"
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

Examples:
  devenv dev         # Generate config using the 'dev' environment
  devenv help        # Display this help message
  devenv version     # Print version information

For additional documentation and examples, visit:
  https://github.com/arashrasoulzadeh/devenv
`
	fmt.Println(helpText)
	os.Exit(0)
}
