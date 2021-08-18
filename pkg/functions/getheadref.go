package functions

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
)

type RefHead struct {
	Ref    string `json:"ref"`
	NodeId string `json:"node_id"`
	Url    string `json:"url"`
	Object struct {
		Sha  string `json:"sha"`
		Type string `json:"type"`
		Url  string `json:"url"`
	} `json:"object"`
}

func GetHeadRef(owner string, repo string, branch string, gittoken string) (string, error) {

	client := resty.New()

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
		Get("https://api.github.com/repos/{owner}/{repo}/git/ref/heads/{branch}")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)
	fmt.Println(resp.RawResponse)
	fmt.Println("String resp: " + resp.String())

	refData := RefHead{}

	//respString := resp.String()

	json.Unmarshal(resp.Body(), &refData)
	//json.Unmarshal([]byte(respString), refData)

	fmt.Println("URL: " + refData.Ref)

	return refData.Object.Url, nil
}
