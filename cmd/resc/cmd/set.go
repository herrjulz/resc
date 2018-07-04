package cmd

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
)

var setCmd = &cobra.Command{
	Use:   "set [owner/repository]",
	Short: "set a default resc repository",
	Run:   set,
}

func set(cmd *cobra.Command, args []string) {
	rsrc := filepath.Join(os.Getenv("HOME"), ".rsrc")
	owner, repo, branch := parseFlags(cmd)

	validateInput(args, owner, repo, branch)

	if len(args) != 0 {
		userRepo := strings.Split(args[0], "/")
		owner = userRepo[0]
		repo = userRepo[1]
	}

	config := setConfig(owner, repo, branch, rsrc)

	conf, err := yaml.Marshal(config)
	exitWithError(err)

	err = ioutil.WriteFile(rsrc, conf, 0644)
	exitWithError(err)
}

func initSet() {
	setCmd.Flags().StringP("branch", "b", "master", "set the default branch of the repository containing the scripts. Default ' master'")
	setCmd.Flags().StringP("owner", "o", "", "set the default owner of a repository")
	setCmd.Flags().StringP("repo", "r", "", "set the default repository to be used")
}

func setConfig(owner, repo, branch, rsrc string) Config {
	var config Config
	if _, err := os.Stat(rsrc); err == nil {
		file, err := ioutil.ReadFile(rsrc)
		exitWithError(err)
		err = yaml.Unmarshal(file, &config)
		exitWithError(err)
		if owner != "" {
			config.User = owner
		}
		if repo != "" {
			config.Repo = repo
		}
		if branch != "" {
			config.Branch = branch
		}
	} else {
		config = Config{
			User:   owner,
			Repo:   repo,
			Branch: branch,
		}
	}
	return config
}

func parseFlags(cmd *cobra.Command) (string, string, string) {
	owner, err := cmd.Flags().GetString("owner")
	exitWithError(err)
	repo, err := cmd.Flags().GetString("repo")
	exitWithError(err)
	branch, err := cmd.Flags().GetString("branch")
	exitWithError(err)
	return owner, repo, branch
}

func validateInput(args []string, owner, repo, branch string) {
	if len(args) == 0 && owner == "" && repo == "" && branch == "" {
		exitWithError(errors.New("No owner, repository, or branch specified. Use 'resc help set' for more info"))
	}
}
