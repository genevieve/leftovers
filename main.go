package main

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/genevievelesperance/leftovers/awsiam"
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
	ip := awsiam.NewInstanceProfiles(iamClient)
	ip.Delete()

	sc := awsiam.NewServerCertificates(iamClient)
	sc.Delete()
}
