package main

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"io"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type FeedRepository struct {
	endpoint        string
	accessKeyID     string
	accessKeySecret string
	bucket          string
	object          string
}

func NewFeedRepository(env Env) *FeedRepository {
	return &FeedRepository{
		endpoint:        env.Gcs.Endpoint,
		accessKeyID:     env.Gcs.AccessKeyID,
		accessKeySecret: env.Gcs.AccessKeySecret,
		bucket:          env.Gcs.Bucket,
		object:          env.Gcs.Object,
	}
}

func (fr *FeedRepository) FetchOldFeed() (Feed, error) {
	sess := session.Must(session.NewSession(&aws.Config{
		Region:           aws.String("auto"),
		Endpoint:         aws.String(fr.endpoint),
		Credentials:      credentials.NewStaticCredentials(fr.accessKeyID, fr.accessKeySecret, ""),
		S3ForcePathStyle: aws.Bool(true),
	}))

	client := s3.New(sess)
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	output, err := client.GetObjectWithContext(ctx, &s3.GetObjectInput{Bucket: &fr.bucket, Key: &fr.object})
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

func (fr *FeedRepository) FetchLatestFeed() (Feed, error) {
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

func (fr *FeedRepository) UploadFeedFile(feed Feed) error {
	b, err := json.MarshalIndent(feed, "", "  ")
	if err != nil {
		return err
	}

	sess := session.Must(session.NewSession(&aws.Config{
		Region:           aws.String("auto"),
		Endpoint:         aws.String(fr.endpoint),
		Credentials:      credentials.NewStaticCredentials(fr.accessKeyID, fr.accessKeySecret, ""),
		S3ForcePathStyle: aws.Bool(true),
	}))
	uploader := s3manager.NewUploader(sess)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	_, err = uploader.UploadWithContext(ctx, &s3manager.UploadInput{Bucket: &fr.bucket, Key: &fr.object, Body: bytes.NewReader(b)})
	if err != nil {
		return err
	}
	return nil
}
