package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/JulzDiverse/resc/scriptmanager"
	"github.com/spf13/cobra"
)

var printCmd = &cobra.Command{
	Use:   "print",
	Short: "prints the desired script",
	Run:   print,
}

func print(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		exitWithError(errors.New("No script specified"))
	}

	userRepo, err := cmd.Flags().GetString("repo")
	exitWithError(err)

	var user, repo string
	if userRepo == "" {
		user, repo = configFromFile()
	} else {
		sl := strings.Split(userRepo, "/")
		user = sl[0]
		repo = sl[1]
	}

	scriptManager := scriptmanager.New(
		"https://raw.githubusercontent.com",
		user,
		repo,
	)

	script, err := scriptManager.GetScript(args[0])
	exitWithError(err)

	fmt.Println(string(script))
}

func initPrint() {
	printCmd.Flags().StringP("repo", "r", "", "name of the repository containing the script. Pattern: <user|org>/<repo>")
}
