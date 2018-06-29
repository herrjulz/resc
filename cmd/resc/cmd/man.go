package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"strings"

	"github.com/JulzDiverse/resc/processor"
	"github.com/JulzDiverse/resc/scriptmanager"
	"github.com/spf13/cobra"
)

var manCmd = &cobra.Command{
	Use:   "man",
	Short: "print the manual of a remote script",
	Run:   man,
}

func man(cmd *cobra.Command, args []string) {
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
	manCmd.Flags().StringP("repo", "r", "", "name of the repository containing the script. Pattern: <user|org>/<repo>")
}
