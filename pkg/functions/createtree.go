package functions

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
)

type TreeObject struct {
	BaseTree string `json:"base_tree"`
	Tree     []Tree  `json:"tree"`
}

type Tree struct {
	Path string `json:"path"`
	Mode string `json:"mode"`
	Type string `json:"type"`
	Sha  string `json:"sha"`
}

type TreeResp struct {
	Sha  string `json:"sha"`
	Url  string `json:"url"`
}

func CreateTree(owner string, repo string, gittoken string, basetreesha string, blobsha string) (string, error) {

	client := resty.New()

	tree := Tree{
		Path: "test",
		Mode: "100644",
		Type: "blob",
		Sha:  blobsha,
	}

	// this is for one file
	treeObj := TreeObject{
		BaseTree: basetreesha,
		Tree: []Tree{tree},
	}

	// https://api.github.com/repos/owner/repo/git/trees
	resp, err := client.R().
		SetHeaders(map[string]string{
			"Authorization": "token " + gittoken,
			"Accept":        "application/vnd.github.v3+json",
		}).
		SetPathParams(map[string]string{
			"owner": owner,
			"repo":  repo,
		}).
		SetBody(treeObj).
		Post("https://api.github.com/repos/{owner}/{repo}/git/trees")

	if err != nil {
		fmt.Println(err)
	}

	treeResp := TreeResp{}

	json.Unmarshal(resp.Body(), &treeResp)

	//todo remove
	fmt.Println("tree: " + treeResp.Sha)

	return treeResp.Sha, nil
}
