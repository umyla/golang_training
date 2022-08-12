package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type repoRequest struct {
	RepoName        string `json:"name"`
	RepoDescription string `json:"description"`
}
type RepoResponse struct {
	RepoId    int    `json:"id"`
	RepoName  string `json:"name"`
	RepoOwner `json:"owner"`
}

type RepoOwner struct {
	Id    int    `json:"id"`
	Login string `json:"login"`
	Url   string `json:"url"`
}
type ErrGithub struct {
	Resource string `json:"resource"`
	Message  string `json:"message"`
	Code     string `json:"code"`
}
type ErrorGithub struct {
	Message string      `json:"message"`
	Err     []ErrGithub `json:"errors"`
}

func main() {
	repo := repoRequest{
		RepoName:        "test406",
		RepoDescription: "test-repo",
	}
	fmt.Println(createRepo(repo))
}

func createRepo(repo repoRequest) error {

	jsonData, err := json.Marshal(repo)
	if err != nil {
		return err
	}
	//construction of the request
	req, err := http.NewRequest(http.MethodPost, "https://api.github.com/user/repos", bytes.NewReader(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "token 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")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req = req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode > 399 {
		var githubError ErrorGithub
		err = json.Unmarshal(data, &githubError)
		if err != nil {
			return err
		}
		fmt.Printf("%+v\n", githubError)
		return errors.New("repo creation failed")
	}

	var result RepoResponse
	err = json.Unmarshal(data, &result)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", result)
	return nil

}
