package main

import (
	"log"
)

func main() {
	env, err := GetEnv()
	if err != nil {
		log.Fatalf("Failed to load env: %v", err)
	}

	feedRepository := NewFeedRepository(env)
	oldFeed, err := feedRepository.FetchOldFeed()
	if err != nil {
		log.Printf("Failed to fetch old feed: %v", err)
		return
	}

	latestFeed, err := feedRepository.FetchLatestFeed()
	if err != nil {
		log.Printf("Failed to fetch latest feed: %v", err)
		return
	}

	blog := Blog{}
	if !blog.IsUpdated(oldFeed, latestFeed) {
		log.Println("Blog is not updated.")
		return
	}

	if err := feedRepository.UploadFeedFile(latestFeed); err != nil {
		log.Printf("Failed to upload feed file: %v", err)
		return
	}
	log.Println("Uploading feed file succeeded.")

	githubApi := NewGitHubApi(env)
	if err := githubApi.CreateRepositoryDispatchEvent(); err != nil {
		log.Printf("Failed to create repository dispatch event: %v", err)
	}
	log.Println("Creating repository dispatch event succeeded.")
}
