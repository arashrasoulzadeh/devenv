package commands

import (
	"fmt"
	"math"
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

	existsCount := 0
	notExistsCount := 0

	for key, value := range r.FinalCfg {
		if _, exists := osEnv[key]; !exists {
			notExistsCount++
			if slices.Contains(args, consts.NoValuesFlag) {
				log.Info(fmt.Sprintf("environment does not have %s", key))
			} else {
				log.Info(fmt.Sprintf("environment does not have %s=%v", key, value))
			}
		} else {
			existsCount++
		}
	}

	if notExistsCount > 0 {
		if slices.Contains(args, consts.PercentEnvFlag) {
			notExistsPercent := int(math.Round((float64(notExistsCount) / float64(len(r.FinalCfg))) * 100))
			log.Info(fmt.Sprintf("about %d percent of env in config does not exists in os env", notExistsPercent))
		}
	}
}
