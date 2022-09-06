package main

import (
	"bytes"
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func UploadFeedFile(feed Feed) error {
	b, err := json.MarshalIndent(feed, "", "  ")
	if err != nil {
		return err
	}

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
	uploader := s3manager.NewUploader(sess)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	_, err = uploader.UploadWithContext(ctx, &s3manager.UploadInput{Bucket: &bucket, Key: &object, Body: bytes.NewReader(b)})
	if err != nil {
		return err
	}
	return nil
}
