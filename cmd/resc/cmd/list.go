package cmd

import (
	"fmt"
	"strings"

	"github.com/JulzDiverse/resc/scriptmanager"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list remote script of a resc repository",
	Run:   list,
}

func list(cmd *cobra.Command, args []string) {
	var user, repo string
	if len(args) == 0 {
		user, repo = configFromFile()
	} else {
		sl := strings.Split(args[0], "/")
		user = sl[0]
		repo = sl[1]
	}

	scriptManager := scriptmanager.New(
		"https://api.github.com",
		user,
		repo,
	)

	list, err := scriptManager.List()
	exitWithError(err)

	listScripts(fmt.Sprintf("%s/%s", user, repo), list)
}

func listScripts(repo string, scripts []string) {
	fmt.Println("ReSc Repository:", repo)
	fmt.Println("\nAvailable Scripts:")
	for _, script := range scripts {
		fmt.Println("  -", script)
	}
}
