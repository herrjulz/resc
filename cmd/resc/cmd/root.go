package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ghodss/yaml"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "resc",
	Short:   "execute remote scripts",
	Long:    `This tool is executing scripts located on github`,
	Version: "0.2.0",
}

func init() {
	initRun()
	initPrint()
	initMan()

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

func exitWithError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Exit: %s", err.Error())
		os.Exit(1)
	}
}

func configFromFile() (string, string) {
	rc := filepath.Join(os.Getenv("HOME"), ".rsrc")
	checkIfConfigFileExists(rc)

	configFile, err := ioutil.ReadFile(rc)
	exitWithError(err)

	config := Config{}
	err = yaml.Unmarshal(configFile, &config)
	exitWithError(err)
	return config.User, config.Repo
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
