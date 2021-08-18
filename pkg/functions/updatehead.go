package functions

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

type UpdateHeadObject struct {
	Sha string `json:"sha"`
}

func UpdateHead(owner string, repo string, branch string, gittoken string, newcommitsha string) (string, error) {

	client := resty.New()

	updateHead := UpdateHeadObject{Sha: newcommitsha}

	resp, err := client.R().
		SetHeaders(map[string]string{
			"Authorization": "token " + gittoken,
			"Accept":        "application/vnd.github.v3+json",
		}).
		SetPathParams(map[string]string{
			"owner":  owner,
			"repo":   repo,
			"branch": branch,
		}).
		SetBody(updateHead).
		Patch("https://api.github.com/repos/{owner}/{repo}/git/refs/heads/{branch}")

	if err != nil {
		fmt.Println(err)
	}

	return resp.Status(), nil
}
