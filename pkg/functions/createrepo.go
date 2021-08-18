package functions

import (
	"fmt"
	"github-connector/pkg/model"
	"github.com/go-resty/resty/v2"
)

func CreateRepo(gittoken string) (string, error) {

	//remember to change json of the model
	repoData := model.Createrepodata{
		Name:              "connector-generated-repo",
		Autoinit:          true,
		Private:           true,
		Gitignoretemplate: "nanoc",
	}

	client := resty.New()
	resp, err := client.R().
		SetHeader("Authorization", "token " + gittoken).
		SetBody(repoData).
		Post("https://api.github.com/user/repos")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Status())

	return resp.Status(), nil
}
