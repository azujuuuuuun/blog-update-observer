package main

import (
	"fmt"
	"os"
	"strings"
)

type GCS struct {
	Endpoint        string
	AccessKeyID     string
	AccessKeySecret string
	Bucket          string
	Object          string
}

type GitHub struct {
	AccessToken string
}
type Env struct {
	Gcs    GCS
	GitHub GitHub
}

func GetEnv() (Env, error) {
	var env Env
	var missing []string

	for k, v := range map[string]*string{
		"STORAGE_API_ENDPOINT":      &env.Gcs.Endpoint,
		"STORAGE_ACCESS_KEY_ID":     &env.Gcs.AccessKeyID,
		"STORAGE_SECRET_ACCESS_KEY": &env.Gcs.AccessKeySecret,
		"GCS_BUCKET_NAME":           &env.Gcs.Bucket,
		"GCS_OBJECT_NAME":           &env.Gcs.Object,
		"GITHUB_ACCESS_TOKEN":       &env.GitHub.AccessToken,
	} {
		*v = os.Getenv(k)

		if *v == "" {
			missing = append(missing, k)
		}
	}

	if len(missing) > 0 {
		return env, fmt.Errorf("missing env(s): " + strings.Join(missing, ", "))
	}

	return env, nil
}
