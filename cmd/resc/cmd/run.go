package cmd

import (
	"strings"

	"github.com/JulzDiverse/resc/runner"
	"github.com/JulzDiverse/resc/scriptmanager"
	"github.com/spf13/cobra"
)

type Config struct {
	User string `yaml:"user"`
	Repo string `yaml:"repo"`
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "execute the desired script",
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	userRepo, err := cmd.Flags().GetString("repo")
	exitWithError(err)

	argsString, err := cmd.Flags().GetString("args")
	exitWithError(err)

	runArgs := strings.Split(argsString, " ")

	var user, repo string
	if userRepo == "" {
		user, repo = configFromFile()
	} else {
		sl := strings.Split(userRepo, "/")
		user = sl[0]
		repo = sl[1]
	}

	runner := runner.New(runner.Default)
	scriptManager := scriptmanager.New(
		"https://raw.githubusercontent.com",
		user,
		repo,
	)

	script, err := scriptManager.GetScript(args[0])
	exitWithError(err)

	_, err = runner.Run(string(script), runArgs...)
	exitWithError(err)
}

func initRun() {
	runCmd.Flags().StringP("repo", "r", "", "name of the repository containing the script. Pattern: <user|org>/<repo>")
	runCmd.Flags().StringP("args", "a", "", "space separated list of arguments: eg '-c config -v var'")
}
