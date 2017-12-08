package main

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/genevievelesperance/leftovers/awsiam"
	flags "github.com/jessevdk/go-flags"
)

type infraCreds struct {
	AWSAccessKeyID     string `long:"aws-access-key-id" env:"AWS_ACCESS_KEY_ID"`
	AWSSecretAccessKey string `long:"aws-secret-access-key" env:"AWS_SECRET_ACCESS_KEY"`
	AWSRegion          string `long:"aws-region" env:"AWS_REGION"`
}

func main() {
	stdout := log.New(os.Stdout, "", 0)

	var c infraCreds
	parser := flags.NewParser(&c, flags.IgnoreUnknown)
	parser.ParseArgs(os.Args)

	if c.AWSAccessKeyID == "" {
		stdout.Fatal("Missing AWS_ACCESS_KEY_ID.")
	}

	if c.AWSSecretAccessKey == "" {
		stdout.Fatal("Missing AWS_SECRET_ACCESS_KEY.")
	}

	if c.AWSRegion == "" {
		stdout.Fatal("Missing AWS_REGION.")
	}

	config := &aws.Config{
		Credentials: credentials.NewStaticCredentials(c.AWSAccessKeyID, c.AWSSecretAccessKey, ""),
		Region:      aws.String(c.AWSRegion),
	}

	iamClient := iam.New(session.New(config))

	ir := awsiam.NewRoles(iamClient)
	ir.Delete()

	ip := awsiam.NewInstanceProfiles(iamClient)
	ip.Delete()

	sc := awsiam.NewServerCertificates(iamClient)
	sc.Delete()
}
