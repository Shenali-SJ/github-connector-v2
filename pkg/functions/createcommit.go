package functions

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
)

type Commit struct {
	Message string   `json:"message"`
	Parents []string `json:"parents"`
	Tree    string   `json:"tree"`
}

type CommitResp struct {
	Sha     string `json:"sha"`
	Url     string `json:"url"`
}

func CreateCommit(owner string, repo string, gittoken string, currentcommitsha string, newtreesha string) (string, error) {

	client := resty.New()

	newCommit := Commit{
		Message: "New commit",
		Parents: []string{currentcommitsha},
		Tree:    newtreesha,
	}

	resp, err := client.R().
		SetHeaders(map[string]string{
			"Authorization": "token " + gittoken,
			"Accept":        "application/vnd.github.v3+json",
		}).
		SetPathParams(map[string]string{
			"owner": owner,
			"repo":  repo,
		}).
		SetBody(newCommit).
		Post("https://api.github.com/repos/{owner}/{repo}/git/commits")

	if err != nil {
		fmt.Println(err)
	}

	commitResp := CommitResp{}

	json.Unmarshal(resp.Body(), &commitResp)

	//todo remove
	fmt.Println("commit: " + commitResp.Sha)

	return commitResp.Sha, nil

}
