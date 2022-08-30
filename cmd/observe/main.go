package main

import "fmt"

func main() {
	latestFeed, err := FetchLatestFeed()
	if err != nil {
		fmt.Printf("%#v", err)
	}
	WriteFeedFile(latestFeed)
}
