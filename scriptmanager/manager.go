package scriptmanager

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

type ScriptManager struct {
	url  string
	user string
	repo string
}

func New(url, user, repo string) *ScriptManager {
	return &ScriptManager{
		url:  url,
		user: user,
		repo: repo,
	}
}

func (s *ScriptManager) GetScript(scriptName string) ([]byte, error) {
	return getter(fmt.Sprintf("%s/%s/%s/master/%s/run.sh", s.url, s.user, s.repo, scriptName))
}

func (s *ScriptManager) GetReadmeForScript(scriptName string) ([]byte, error) {
	return getter(fmt.Sprintf("%s/%s/%s/master/%s/README.md", s.url, s.user, s.repo, scriptName))
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
