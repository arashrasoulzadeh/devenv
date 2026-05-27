package commands

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/arashrasoulzadeh/devenv/src/app"
	"github.com/arashrasoulzadeh/devenv/src/consts"
	"github.com/arashrasoulzadeh/devenv/src/log"
)

func EnvCommand(args []string, r *app.Runner) {
	osEnv := make(map[string]string)

	for _, entry := range os.Environ() {
		parts := strings.SplitN(entry, "=", 2)
		if len(parts) == 2 {
			osEnv[parts[0]] = parts[1]
		}
	}

	if err := r.Run([]string{"env", args[2], consts.DontCommitFlag}); err != nil {
		log.Fatal(err)
	}

	for key, value := range r.FinalCfg {
		if _, exists := osEnv[key]; !exists {
			if slices.Contains(args, consts.NoValuesFlag) {
				log.Info(fmt.Sprintf("environment does not have %s", key))
			} else {
				log.Info(fmt.Sprintf("environment does not have %s=%v", key, value))
			}
		}
	}
}
