package acceptance

import (
	"os"
	"strings"

	awslib "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
	"github.com/genevieve/leftovers/app"
	. "github.com/onsi/gomega"
)

type AWSAcceptance struct {
	AccessKeyId     string
	SecretAccessKey string
	Region          string
	Logger          *app.Logger
}

func NewAWSAcceptance() *AWSAcceptance {
	return &AWSAcceptance{}
}

func (a *AWSAcceptance) ReadyToTest() bool {
	iaas := os.Getenv("LEFTOVERS_ACCEPTANCE")
	if iaas == "" {
		return false
	}

	if strings.ToLower(iaas) != "aws" {
		return false
	}

	a.AccessKeyId = os.Getenv("BBL_AWS_ACCESS_KEY_ID")
	a.SecretAccessKey = os.Getenv("BBL_AWS_SECRET_ACCESS_KEY")
	a.Region = os.Getenv("BBL_AWS_REGION")

	logger := app.NewLogger(os.Stdin, os.Stdout, true)
	a.Logger = logger

	return true
}

func (a *AWSAcceptance) CreateKeyPair(name string) {
	config := &awslib.Config{
		Credentials: credentials.NewStaticCredentials(a.AccessKeyId, a.SecretAccessKey, ""),
		Region:      awslib.String(a.Region),
	}

	client := awsec2.New(session.New(config))

	_, err := client.CreateKeyPair(&awsec2.CreateKeyPairInput{KeyName: awslib.String(name)})
	if cast, ok := err.(awserr.Error); ok {
		if cast.Code() == "InvalidKeyPair.Duplicate" {
			return
		}
	}
	Expect(err).NotTo(HaveOccurred())
}
