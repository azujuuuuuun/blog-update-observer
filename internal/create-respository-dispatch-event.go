package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type RequestBody struct {
	EventType string `json:"event_type"`
}

func CreateRepositoryDispatchEvent() {
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
	accessToken := os.Getenv("GITHUB_ACCESS_TOKEN")
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", "Bearer "+accessToken)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("%#v", err)
		return
	}
	defer resp.Body.Close()
	fmt.Printf("%#v", resp)
}
