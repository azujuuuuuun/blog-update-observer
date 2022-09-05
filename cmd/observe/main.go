package main

import "fmt"

func main() {
	oldFeed, err := FetchLocalOldFeed()
	if err != nil {
		fmt.Printf("%#v", err)
		return
	}

	latestFeed, err := FetchLatestFeed()
	if err != nil {
		fmt.Printf("%#v", err)
		return
	}

	if !IsBlogUpdated(oldFeed, latestFeed) {
		fmt.Printf("Blog is not updated.")
		return
	}

	WriteFeedFile(latestFeed)
}
