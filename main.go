package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	oldFeed, err := FetchOldFeed()
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

	err = UploadFeedFile(latestFeed)
	if err != nil {
		fmt.Printf("%#v", err)
		return
	}
	fmt.Printf("Uploading feed file succeeded.")

	CreateRepositoryDispatchEvent()
}
