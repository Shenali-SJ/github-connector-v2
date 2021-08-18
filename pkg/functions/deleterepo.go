package functions

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

func DeleteRepo(gittoken string, repoowner string, reponame string) (string, error) {

	client := resty.New()

	response, err := client.R().
		SetHeaders(map[string]string{
			"Authorization": "token " + gittoken,
			"Accept": "application/vnd.github.v3+json",
		}).
		SetPathParams(map[string]string{
			"owner":      repoowner,
			"repository": reponame,
		}).
		Delete("https://api.github.com/repos/{owner}/{repository}")

	if err != nil {
		fmt.Println(err)
	}

	return response.Status(), nil
}
