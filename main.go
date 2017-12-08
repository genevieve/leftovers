package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
)

func main() {
	stdout := log.New(os.Stdout, "", 0)
	// stderr := log.New(os.Stderr, "", 0)

	accessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	if accessKeyID == "" {
		stdout.Fatal("Missing AWS_ACCESS_KEY_ID.")
	}

	secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	if secretAccessKey == "" {
		stdout.Fatal("Missing AWS_SECRET_ACCESS_KEY.")
	}

	region := os.Getenv("AWS_REGION")
	if region == "" {
		stdout.Fatal("Missing AWS_REGION.")
	}

	config := &aws.Config{
		Credentials: credentials.NewStaticCredentials(accessKeyID, secretAccessKey, ""),
		Region:      aws.String(region),
	}

	iamClient := iam.New(session.New(config))

	profiles, err := iamClient.ListInstanceProfiles(&iam.ListInstanceProfilesInput{})
	if err != nil {
		fmt.Printf("ERROR listing instance profiles: %s", err)
	}

	for _, p := range profiles.InstanceProfiles {
		n := p.InstanceProfileName
		_, err := iamClient.DeleteInstanceProfile(&iam.DeleteInstanceProfileInput{InstanceProfileName: n})
		if err == nil {
			fmt.Printf("SUCCESS deleting instance profile %s\n", &n)
		} else {
			fmt.Printf("ERROR deleting instance profile %s: %s\n", &n, err)
		}
	}
}
