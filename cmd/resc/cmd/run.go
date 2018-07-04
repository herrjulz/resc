package cmd

import (
	"errors"
	"strings"

	"github.com/JulzDiverse/resc/models/github"
	"github.com/JulzDiverse/resc/runner"
	"github.com/spf13/cobra"
)

type Config struct {
	User   string `yaml:"user"`
	Repo   string `yaml:"repo"`
	Branch string `yaml:"branch"`
}

var runCmd = &cobra.Command{
	Use:   "run <script-name>",
	Short: "run a remote script",
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	scriptPath, err := cmd.Flags().GetString("script")
	exitWithError(err)

	if len(args) == 0 && scriptPath == "" {
		exitWithError(errors.New("No script specified"))
	}

	userRepo, err := cmd.Flags().GetString("repo")
	exitWithError(err)

	argsString, err := cmd.Flags().GetString("args")
	exitWithError(err)

	branch, err := cmd.Flags().GetString("branch")
	exitWithError(err)

	scriptManager, _, _, _ := initScriptManager(github.RawContentUrl, userRepo, branch)
	runner := runner.New(runner.Default)

	var script []byte
	if scriptPath == "" {
		script, err = scriptManager.Get(args[0])
		exitWithError(err)
	} else {
		script, err = scriptManager.GetScript(scriptPath)
	}

	runArgs := strings.Split(argsString, " ")
	_, err = runner.Run(string(script), runArgs...)
	exitWithError(err)
}

func initRun() {
	runCmd.Flags().StringP("repo", "r", "", "name of the repository containing the script. Pattern: <owner>/<repo>")
	runCmd.Flags().StringP("branch", "b", "", "branch of the repository containing the script. Default: master")
	runCmd.Flags().StringP("script", "s", "", "path to a specific script file in the specified repository (eg topDir/subDir/script.sh)")
	runCmd.Flags().StringP("args", "a", "", "space separated list of arguments: eg '-c config -v var'")
}
