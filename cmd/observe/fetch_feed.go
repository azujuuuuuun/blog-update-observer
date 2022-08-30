package main

import (
	"encoding/xml"
	"io"
	"net/http"
)

func FetchFeed() (Feed, error) {
	resp, err := http.Get("https://azujuuuuuun.hatenablog.com/feed")
	if err != nil {
		return Feed{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Feed{}, err
	}
	var feed Feed
	err = xml.Unmarshal(body, &feed)
	if err != nil {
		return Feed{}, err
	}
	return feed, nil
}
