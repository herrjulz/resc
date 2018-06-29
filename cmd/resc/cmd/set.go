package cmd

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "set base user and uri",
	Run:   set,
}

func set(cmd *cobra.Command, args []string) {
	userRepo := strings.Split(args[0], "/")
	config := Config{
		User: userRepo[0],
		Repo: userRepo[1],
	}

	conf, err := yaml.Marshal(config)
	exitWithError(err)

	err = ioutil.WriteFile(filepath.Join(os.Getenv("HOME"), ".rsrc"), conf, 0644)
}
