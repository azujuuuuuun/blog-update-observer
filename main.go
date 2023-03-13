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

	env := GetEnv()

	feedRepository := NewFeedRepository(env)
	oldFeed, err := feedRepository.FetchOldFeed()
	if err != nil {
		fmt.Printf("%#v", err)
		return
	}

	latestFeed, err := feedRepository.FetchLatestFeed()
	if err != nil {
		fmt.Printf("%#v", err)
		return
	}

	blog := Blog{}
	if !blog.IsUpdated(oldFeed, latestFeed) {
		fmt.Printf("Blog is not updated.")
		return
	}

	err = feedRepository.UploadFeedFile(latestFeed)
	if err != nil {
		fmt.Printf("%#v", err)
		return
	}
	fmt.Printf("Uploading feed file succeeded.")

	githubApi := NewGitHubApi(env)
	githubApi.CreateRepositoryDispatchEvent()
}
