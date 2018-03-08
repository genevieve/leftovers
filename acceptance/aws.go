package acceptance

import (
	"os"

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

func NewAWSAcceptance() AWSAcceptance {
	accessKeyId := os.Getenv("BBL_AWS_ACCESS_KEY_ID")
	Expect(accessKeyId).NotTo(Equal(""))

	secretAccessKey := os.Getenv("BBL_AWS_SECRET_ACCESS_KEY")
	Expect(secretAccessKey).NotTo(Equal(""))

	region := os.Getenv("BBL_AWS_REGION")
	Expect(region).NotTo(Equal(""))

	return AWSAcceptance{
		AccessKeyId:     accessKeyId,
		SecretAccessKey: secretAccessKey,
		Region:          region,
		Logger:          app.NewLogger(os.Stdin, os.Stdout, true),
	}
}

func (a AWSAcceptance) CreateKeyPair(name string) {
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
