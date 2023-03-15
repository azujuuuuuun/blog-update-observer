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

func NewGitHubApi(env Env) *GitHubApi {
	return &GitHubApi{accessToken: env.GitHub.AccessToken}
}

func (api *GitHubApi) CreateRepositoryDispatchEvent() error {
	requestBody := RequestBody{EventType: "blog-updated"}
	b, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %w", err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://api.github.com/repos/azujuuuuuun/azujuuuuuun.github.io/dispatches", bytes.NewBuffer(b))
	if err != nil {
		return fmt.Errorf("failed to construct a request: %w", err)
	}
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", "Bearer "+api.accessToken)
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to request repository dispatch: %w", err)
	}
	defer resp.Body.Close()

	return nil
}
