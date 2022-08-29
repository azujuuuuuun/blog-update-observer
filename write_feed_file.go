package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func WriteFeedFile(feed Feed) {
	b, err := json.MarshalIndent(feed, "", "  ")
	if err != nil {
		fmt.Printf("%#v", err)
		return
	}
	f, err := os.Create("output/feed.json")
	if err != nil {
		fmt.Printf("%#v", err)
		return
	}
	c, err := f.Write(b)
	if err != nil {
		fmt.Printf("%#v", err)
		return
	}
	fmt.Printf("%#v", c)
}
