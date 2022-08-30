package main

import "encoding/xml"

type Link struct {
	XMLName xml.Name `xml:"link" json:"-"`
	Href    string   `xml:"href,attr" json:"href"`
}

type Author struct {
	Name string `xml:"name" json:"name"`
}

type Feed struct {
	Title    string `xml:"title" json:"title"`
	Subtitle string `xml:"subtitle" json:"subTitle"`
	Link     Link   `xml:"link" json:"link"`
	Updated  string `xml:"updated" json:"updated"`
	Author   Author `xml:"author" json:"author"`
	Id       string `xml:"id" json:"id"`
	Entry    []struct {
		Title     string `xml:"title" json:"title"`
		Link      []Link `xml:"link" json:"link"`
		Id        string `xml:"id" json:"id"`
		Published string `xml:"published" json:"published"`
		Updated   string `xml:"updated" json:"updated"`
		Summary   string `xml:"summary" json:"summary"`
		Content   string `xml:"content" json:"content"`
		Category  struct {
			XMLName xml.Name `xml:"category" json:"-"`
			Term    string   `xml:"term,attr" json:"term"`
			Label   string   `xml:"label,attr" json:"label"`
		} `xml:"category" json:"category"`
		Author Author `xml:"author" json:"author"`
	} `xml:"entry" json:"entry"`
}
