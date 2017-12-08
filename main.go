package main

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/genevievelesperance/leftovers/app"
	"github.com/genevievelesperance/leftovers/awsec2"
	"github.com/genevievelesperance/leftovers/awsiam"
	flags "github.com/jessevdk/go-flags"
)

type opts struct {
	NoConfirm bool `short:"n"  long:"no-confirm"`

	AWSAccessKeyID     string `           long:"aws-access-key-id"     env:"AWS_ACCESS_KEY_ID"`
	AWSSecretAccessKey string `           long:"aws-secret-access-key" env:"AWS_SECRET_ACCESS_KEY"`
	AWSRegion          string `           long:"aws-region"            env:"AWS_REGION"`
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

	iamClient := iam.New(session.New(config))

	ir := awsiam.NewRoles(iamClient, logger)
	err = ir.Delete()
	if err != nil {
		log.Fatalf("\n\n%s\n", err)
	}

	ip := awsiam.NewInstanceProfiles(iamClient, logger)
	err = ip.Delete()
	if err != nil {
		log.Fatalf("\n\n%s\n", err)
	}

	sc := awsiam.NewServerCertificates(iamClient, logger)
	err = sc.Delete()
	if err != nil {
		log.Fatalf("\n\n%s\n", err)
	}

	ec2Client := ec2.New(session.New(config))

	vo := awsec2.NewVolumes(ec2Client, logger)
	err = vo.Delete()
	if err != nil {
		log.Fatalf("\n\n%s\n", err)
	}
}
