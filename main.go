package main

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
	awsiam "github.com/aws/aws-sdk-go/service/iam"
	"github.com/genevievelesperance/leftovers/app"
	"github.com/genevievelesperance/leftovers/aws/ec2"
	"github.com/genevievelesperance/leftovers/aws/iam"
	flags "github.com/jessevdk/go-flags"
)

type opts struct {
	NoConfirm bool `short:"n"  long:"no-confirm"`

	AWSAccessKeyID     string `           long:"aws-access-key-id"     env:"AWS_ACCESS_KEY_ID"`
	AWSSecretAccessKey string `           long:"aws-secret-access-key" env:"AWS_SECRET_ACCESS_KEY"`
	AWSRegion          string `           long:"aws-region"            env:"AWS_REGION"`
}

type resource interface {
	Delete() error
}

func main() {
	log.SetFlags(0)

	var c opts
	parser := flags.NewParser(&c, flags.HelpFlag|flags.PrintErrors)
	_, err := parser.ParseArgs(os.Args)
	if err != nil {
		os.Exit(0)
	}

	logger := app.NewLogger(os.Stdout, os.Stdin, c.NoConfirm)

	if c.AWSAccessKeyID == "" {
		log.Fatal("Missing AWS_ACCESS_KEY_ID.")
	}

	if c.AWSSecretAccessKey == "" {
		log.Fatal("Missing AWS_SECRET_ACCESS_KEY.")
	}

	if c.AWSRegion == "" {
		log.Fatal("Missing AWS_REGION.")
	}

	config := &aws.Config{
		Credentials: credentials.NewStaticCredentials(c.AWSAccessKeyID, c.AWSSecretAccessKey, ""),
		Region:      aws.String(c.AWSRegion),
	}

	iamClient := awsiam.New(session.New(config))
	ec2Client := awsec2.New(session.New(config))

	ir := iam.NewRoles(iamClient, logger)
	ip := iam.NewInstanceProfiles(iamClient, logger)
	sc := iam.NewServerCertificates(iamClient, logger)
	vo := ec2.NewVolumes(ec2Client, logger)

	resources := []resource{ir, ip, sc, vo}
	for _, r := range resources {
		if err = r.Delete(); err != nil {
			log.Fatalf("\n\n%s\n", err)
		}
	}
}
