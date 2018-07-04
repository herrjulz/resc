package scriptmanager

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

type Content struct {
	Name        string `json:"name"`
	ContentType string `json:"type"`
}

func (s *ScriptManager) List() ([]string, error) {
	contents, err := getContents(fmt.Sprintf("%s/repos/%s/%s/contents", s.url, s.user, s.repo), s.branch)
	if err != nil {
		return nil, err
	}

	validScripts := []string{}
	for _, content := range contents {
		if content.ContentType == "dir" {
			if ok, err := exists(fmt.Sprintf("%s/repos/%s/%s/contents/%s/.resc", s.url, s.user, s.repo, content.Name), s.branch); ok {
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

func doQuery(url string, branch string) (*http.Response, error) {
	client := http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to build request")
	}

	q := req.URL.Query()
	q.Add("ref", branch)
	req.URL.RawQuery = q.Encode()

	return client.Do(req)
}

func getContents(url, branch string) ([]Content, error) {
	resp, err := doQuery(url, branch)
	if err != nil {
		return nil, errors.Wrap(err, "failed to perform http GET")
	}

	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var contents []Content
	err = json.Unmarshal(raw, &contents)
	if err != nil {
		return nil, err
	}

	return contents, nil
}

func exists(url string, branch string) (bool, error) {
	resp, err := doQuery(url, branch)
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
