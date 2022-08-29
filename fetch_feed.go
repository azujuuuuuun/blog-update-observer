package main

import (
	"encoding/xml"
	"io"
	"net/http"
)

type Link struct {
	XMLName xml.Name `xml:"link"`
	Href    string   `xml:"href,attr"`
}

type Author struct {
	Name string `xml:"name"`
}

type Feed struct {
	Title    string `xml:"title"`
	Subtitle string `xml:"subtitle"`
	Link     Link   `xml:"link"`
	Updated  string `xml:"updated"`
	Author   Author `xml:"author"`
	Id       string `xml:"id"`
	Entry    []struct {
		Title     string `xml:"title"`
		Link      []Link `xml:"link"`
		Id        string `xml:"id"`
		Published string `xml:"published"`
		Updated   string `xml:"updated"`
		Summary   string `xml:"summary"`
		Content   string `xml:"content"`
		Category  struct {
			XMLName xml.Name `xml:"category"`
			Term    string   `xml:"term,attr"`
			Label   string   `xml:"label,attr"`
		} `xml:"category"`
		Author Author `xml:"author"`
	} `xml:"entry"`
}

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
