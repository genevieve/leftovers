package ec2

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
)

type KeyPairs struct {
	client ec2Client
	logger logger
}

func NewKeyPairs(client ec2Client, logger logger) KeyPairs {
	return KeyPairs{
		client: client,
		logger: logger,
	}
}

func (a KeyPairs) Delete() error {
	keyPairs, err := a.client.DescribeKeyPairs(&awsec2.DescribeKeyPairsInput{})
	if err != nil {
		return fmt.Errorf("Describing key pairs: %s", err)
	}

	for _, t := range keyPairs.KeyPairs {
		n := *t.KeyName
		proceed := a.logger.Prompt(fmt.Sprintf("Are you sure you want to delete key pair %s?", n))
		if !proceed {
			continue
		}

		_, err := a.client.DeleteKeyPair(&awsec2.DeleteKeyPairInput{KeyName: aws.String(n)})
		if err == nil {
			a.logger.Printf("SUCCESS deleting key pair %s\n", n)
		} else {
			a.logger.Printf("ERROR deleting key pair %s: %s\n", n, err)
		}
	}

	return nil
}
