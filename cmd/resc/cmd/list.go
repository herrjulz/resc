package cmd

import (
	"fmt"

	"github.com/JulzDiverse/resc/models/github"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list [owner/repository]",
	Short: "list remote scripts of a repository",
	Run:   list,
}

func list(cmd *cobra.Command, args []string) {
	userRepo, err := cmd.Flags().GetString("repo")
	exitWithError(err)

	branch, err := cmd.Flags().GetString("branch")
	exitWithError(err)

	scriptManager, user, repo, branch := initScriptManager(github.ApiUrl, userRepo, branch)

	list, err := scriptManager.List()
	exitWithError(err)

	listScripts(fmt.Sprintf("%s/%s on branch %s", user, repo, branch), list)
}

func initList() {
	listCmd.Flags().StringP("branch", "b", "", "branch of the repository containing the scripts. Default: master")
	listCmd.Flags().StringP("repo", "r", "", "name of the repository containing the script. Pattern: <user|org>/<repo>")
}

func listScripts(repo string, scripts []string) {
	fmt.Println("ReSc Repository:", repo)
	fmt.Println("\nAvailable Scripts:")
	for _, script := range scripts {
		fmt.Println("  -", script)
	}
}
