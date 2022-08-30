package main

import (
	"log"

	"github.com/azujuuuuuun/blog-update-observer/internal"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	internal.CreateRepositoryDispatchEvent()
}
