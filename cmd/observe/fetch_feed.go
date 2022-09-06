package main

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func FetchOldFeed() (Feed, error) {
	endpoint := os.Getenv("STORAGE_API_ENDPOINT")
	accessKeyID := os.Getenv("STORAGE_ACCESS_KEY_ID")
	accessKeySecret := os.Getenv("STORAGE_SECRET_ACCESS_KEY")
	bucket := os.Getenv("GCS_BUCKET_NAME")
	object := os.Getenv("GCS_OBJECT_NAME")

	sess := session.Must(session.NewSession(&aws.Config{
		Region:           aws.String("auto"),
		Endpoint:         aws.String(endpoint),
		Credentials:      credentials.NewStaticCredentials(accessKeyID, accessKeySecret, ""),
		S3ForcePathStyle: aws.Bool(true),
	}))

	client := s3.New(sess)
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	output, err := client.GetObjectWithContext(ctx, &s3.GetObjectInput{Bucket: &bucket, Key: &object})
	if err != nil {
		return Feed{}, err
	}

	defer output.Body.Close()
	body, err := io.ReadAll(output.Body)
	if err != nil {
		return Feed{}, err
	}

	var feed Feed
	err = json.Unmarshal(body, &feed)
	if err != nil && err != io.EOF {
		return Feed{}, err
	}

	return feed, nil
}

func FetchLatestFeed() (Feed, error) {
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
