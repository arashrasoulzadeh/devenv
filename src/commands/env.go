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
	envs := os.Environ()

	environementVariables := make(map[string]string)

	for _, env := range envs {
		// each env supposed to be in "KEY=value" format
		parts := strings.SplitN(env, "=", 2)
		key := parts[0]
		value := parts[1]
		environementVariables[key] = value

	}

	//replace requested env name with "env" in command args
	appArgs := []string{"env", args[2], consts.DontCommitFlag}

	if err := r.Run(appArgs); err != nil {
		log.Fatal(err)
	}

	//check non existing envs
	for key, value := range r.FinalCfg {
		if _, exists := environementVariables[key]; !exists {
			if slices.Contains(args, consts.NoValuesFlag) {
				log.Info(fmt.Sprintf("environment does not have %s", key))
			} else {
				log.Info(fmt.Sprintf("environment does not have %s=%v", key, value))
			}
		}
	}
}
