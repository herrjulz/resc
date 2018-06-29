package scriptmanager

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

type Content struct {
	Name        string `json:"name"`
	ContentType string `json:"type"`
}

func (s *ScriptManager) List() ([]string, error) {
	list, err := getter(fmt.Sprintf("%s/repos/%s/%s/contents", s.url, s.user, s.repo))
	if err != nil {
		return nil, err
	}

	var contents []Content
	err = json.Unmarshal(list, &contents)
	if err != nil {
		return nil, err
	}

	validScripts := []string{}
	for _, content := range contents {
		if content.ContentType == "dir" {
			if ok, err := exists(fmt.Sprintf("%s/repos/%s/%s/contents/%s/.resc", s.url, s.user, s.repo, content.Name)); ok {
				validScripts = append(validScripts, content.Name)
			} else {
				if err != nil {
					return nil, err
				}
			}
		}
	}
	return validScripts, nil
}

func exists(url string) (bool, error) {
	resp, err := http.Get(url)
	if err != nil {
		return false, errors.Wrap(err, "failed to perform http GET")
	}

	if resp.StatusCode == http.StatusOK {
		return true, nil
	} else if resp.StatusCode == http.StatusNotFound {
		return false, nil
	} else {
		return false, errors.New(fmt.Sprintf("an error ocurred while checking repository: %s", resp.Status))
	}
}
