package main

import "os"

type GitHub struct {
	AccessToken string
}

type GCS struct {
	Endpoint        string
	AccessKeyID     string
	AccessKeySecret string
	Bucket          string
	Object          string
}

type Env struct {
	Gcs    GCS
	GitHub GitHub
}

func GetEnv() Env {
	gcs := GCS{
		Endpoint:        os.Getenv("STORAGE_API_ENDPOINT"),
		AccessKeyID:     os.Getenv("STORAGE_ACCESS_KEY_ID"),
		AccessKeySecret: os.Getenv("STORAGE_SECRET_ACCESS_KEY"),
		Bucket:          os.Getenv("GCS_BUCKET_NAME"),
		Object:          os.Getenv("GCS_OBJECT_NAME"),
	}

	github := GitHub{
		AccessToken: os.Getenv("GITHUB_ACCESS_TOKEN"),
	}

	return Env{
		Gcs:    gcs,
		GitHub: github,
	}
}
