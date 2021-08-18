package functions

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
)

type Blob struct {
	Content  string `json:"content"`
	Encoding string `json:"encoding"`
}

type BlobResp struct {
	Sha string `json:"sha"`
	Url string `json:"url"`
}

func CreateBlob(owner string, repo string, gittoken string) (string, error) {

	client := resty.New()

	// fixme content
	body := Blob{
		Content:  "New content",
		Encoding: "utf-8",
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
		SetBody(body).
		Post("https://api.github.com/repos/{owner}/{repo}/git/blobs")

	if err != nil {
		fmt.Println(err)
	}

	blobData := BlobResp{}

	json.Unmarshal(resp.Body(), &blobData)

	return blobData.Sha, nil

}
