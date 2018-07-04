package cmd

import (
	"errors"
	"fmt"

	"github.com/JulzDiverse/resc/models/github"
	"github.com/spf13/cobra"
)

var printCmd = &cobra.Command{
	Use:   "print <script-name>",
	Short: "prints the desired script",
	Run:   print,
}

func print(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		exitWithError(errors.New("No script specified"))
	}

	userRepo, err := cmd.Flags().GetString("repo")
	exitWithError(err)

	branch, err := cmd.Flags().GetString("branch")
	exitWithError(err)

	scriptPath, err := cmd.Flags().GetString("script")
	exitWithError(err)

	scriptManager, _, _, _ := initScriptManager(github.RawContentUrl, userRepo, branch)

	var script []byte
	if scriptPath == "" {
		script, err = scriptManager.Get(args[0])
		exitWithError(err)
	} else {
		script, err = scriptManager.GetScript(scriptPath)
	}

	fmt.Println(string(script))
}

func initPrint() {
	printCmd.Flags().StringP("repo", "r", "", "name of the repository containing the script. Pattern: <owner>/<repo>")
	printCmd.Flags().StringP("branch", "b", "", "branch of the repository containing the script. Default: master")
}
