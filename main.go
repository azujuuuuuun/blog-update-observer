package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	env := GetEnv()

	feedRepository := NewFeedRepository(env)
	oldFeed, err := feedRepository.FetchOldFeed()
	if err != nil {
		log.Printf("%#v", err)
		return
	}

	latestFeed, err := feedRepository.FetchLatestFeed()
	if err != nil {
		log.Printf("%#v", err)
		return
	}

	blog := Blog{}
	if !blog.IsUpdated(oldFeed, latestFeed) {
		log.Println("Blog is not updated.")
		return
	}

	if err := feedRepository.UploadFeedFile(latestFeed); err != nil {
		log.Printf("%#v", err)
		return
	}
	log.Println("Uploading feed file succeeded.")

	githubApi := NewGitHubApi(env)
	if err := githubApi.CreateRepositoryDispatchEvent(); err != nil {
		log.Printf("%#v", err)
	}
	log.Println("Creating repository dispatch event succeeded.")
}
