package functions

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
)

type CurrentCommit struct {
	Sha     string `json:"sha"`
	Tree struct {
		Sha string `json:"sha"`
		Url string `json:"url"`
	} `json:"tree"`
}

func GetCurrentCommit(commiturl string, gittoken string) ([]string, error) {

	client := resty.New()

	resp, err := client.R().
		SetHeaders(map[string]string{
			"Authorization": "token " + gittoken,
			"Accept":        "application/vnd.github.v3+json",
		}).
		Get(commiturl)

	if err != nil {
		fmt.Println(err)
	}

	commitData := CurrentCommit{}

	json.Unmarshal(resp.Body(), &commitData)

	// todo remove
	fmt.Println(commitData.Sha, commitData.Tree.Sha)

	// todo check
	//commitSha, treeSha := &commitData.Sha, &commitData.Tree.Sha
	//commitShaPointer, treeShaPointer := &commitSha, &treeSha

	//commit := []*string{commitSha, treeSha}

	return []string{commitData.Sha, commitData.Tree.Sha}, nil

}
