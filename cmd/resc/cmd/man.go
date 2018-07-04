package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"strings"

	"github.com/JulzDiverse/resc/models/github"
	"github.com/JulzDiverse/resc/processor"
	"github.com/spf13/cobra"
)

var manCmd = &cobra.Command{
	Use:   "man <script-name>",
	Short: "print the manual of a remote script",
	Run:   man,
}

func man(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		exitWithError(errors.New("No script specified"))
	}

	userRepo, err := cmd.Flags().GetString("repo")
	exitWithError(err)

	branch, err := cmd.Flags().GetString("branch")
	exitWithError(err)

	scriptManager, _, _, _ := initScriptManager(github.RawContentUrl, userRepo, branch)

	readme, err := scriptManager.GetReadmeForScript(args[0])
	exitWithError(err)

	scanner := bufio.NewScanner(strings.NewReader(string(readme)))
	processor := processor.New()
	for scanner.Scan() {
		line := processor.Process(scanner.Text())
		fmt.Println(line)
	}
}

func initMan() {
	manCmd.Flags().StringP("repo", "r", "", "name of the repository containing the script. Pattern: <owner>/<repo>")
	manCmd.Flags().StringP("branch", "b", "", "branch of the repository containing the script. Default: master")
}
