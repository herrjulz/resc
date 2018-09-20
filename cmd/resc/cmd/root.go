package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/JulzDiverse/resc/scriptmanager"
	"github.com/ghodss/yaml"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "resc",
	Short:   "execute remote scripts",
	Long:    `This tool is executing scripts located on github`,
	Version: "0.4.0",
}

func init() {
	initRun()
	initPrint()
	initMan()
	initSet()
	initList()

	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(setCmd)
	rootCmd.AddCommand(printCmd)
	rootCmd.AddCommand(manCmd)
	rootCmd.AddCommand(listCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initScriptManager(url, userRepoString, branchFromFlag string) (*scriptmanager.ScriptManager, string, string, string) {
	var user, repo, branch string
	if branchFromFlag == "" {
		branch = "master"
	}

	if userRepoString == "" {
		user, repo, branch = configFromFile()
	} else {
		sl := strings.Split(userRepoString, "/")
		user = sl[0]
		repo = sl[1]
	}

	if branchFromFlag != "" {
		branch = branchFromFlag
	}

	return scriptmanager.New(
		url,
		user,
		repo,
		branch,
	), user, repo, branch
}

func exitWithError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Exit: %s", err.Error())
		os.Exit(1)
	}
}

func configFromFile() (string, string, string) {
	rc := filepath.Join(os.Getenv("HOME"), ".rsrc")
	checkIfConfigFileExists(rc)

	configFile, err := ioutil.ReadFile(rc)
	exitWithError(err)

	config := Config{}
	err = yaml.Unmarshal(configFile, &config)
	exitWithError(err)
	return config.User, config.Repo, config.Branch
}

func checkIfConfigFileExists(rc string) {
	if _, err := os.Stat(rc); os.IsNotExist(err) {
		msg := `No script repository set!

Either use the set command to set a script repository:

$ screm set <github-user|github-org>/<github-repo>

or

use the run -r/--repo option:

$ screm run <script> -r <github-user|github-org>/<github-repo>
`
		fmt.Println(msg)
		exitWithError(errors.New("no repository provided"))
	}
}
