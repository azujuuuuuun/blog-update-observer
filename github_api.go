package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type GitHubApi struct {
	accessToken string
}

type RequestBody struct {
	EventType string `json:"event_type"`
}

func (api *GitHubApi) CreateRepositoryDispatchEvent() {
	requestBody := RequestBody{EventType: "blog-updated"}
	b, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Printf("%#v", err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://api.github.com/repos/azujuuuuuun/azujuuuuuun.github.io/dispatches", bytes.NewBuffer(b))
	if err != nil {
		fmt.Printf("%#v", err)
		return
	}
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", "Bearer "+api.accessToken)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("%#v", err)
		return
	}
	defer resp.Body.Close()
	fmt.Printf("%#v", resp)
}

func NewGitHubApi(env Env) *GitHubApi {
	return &GitHubApi{accessToken: env.GitHub.AccessToken}
}
