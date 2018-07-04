package scriptmanager

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

type ScriptManager struct {
	url    string
	user   string
	repo   string
	branch string
}

func New(url, user, repo, branch string) *ScriptManager {
	return &ScriptManager{
		url:    url,
		user:   user,
		repo:   repo,
		branch: branch,
	}
}

func (s *ScriptManager) Get(scriptName string) ([]byte, error) {
	return getter(fmt.Sprintf("%s/%s/%s/%s/%s/run.sh", s.url, s.user, s.repo, s.branch, scriptName))
}

func (s *ScriptManager) GetScript(scriptPath string) ([]byte, error) {
	return getter(fmt.Sprintf("%s/%s/%s/%s/%s", s.url, s.user, s.repo, s.branch, scriptPath))
}

func (s *ScriptManager) GetReadmeForScript(scriptName string) ([]byte, error) {
	return getter(fmt.Sprintf("%s/%s/%s/%s/%s/README.md", s.url, s.user, s.repo, s.branch, scriptName))
}

func getter(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "failed to perform http GET")
	}

	var file []byte
	if resp.StatusCode == http.StatusOK {
		file, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.Wrap(err, "could not read response body")
		}
	} else {
		return nil, errors.New(fmt.Sprintf("requesting file failed: %s", resp.Status))
	}

	return file, nil
}
