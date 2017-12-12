package ec2

import (
	"fmt"

	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
)

type keyPairClient interface {
	DescribeKeyPairs(*awsec2.DescribeKeyPairsInput) (*awsec2.DescribeKeyPairsOutput, error)
	DeleteKeyPair(*awsec2.DeleteKeyPairInput) (*awsec2.DeleteKeyPairOutput, error)
}

type KeyPairs struct {
	client keyPairClient
	logger logger
}

func NewKeyPairs(client keyPairClient, logger logger) KeyPairs {
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

		_, err := a.client.DeleteKeyPair(&awsec2.DeleteKeyPairInput{KeyName: t.KeyName})
		if err == nil {
			a.logger.Printf("SUCCESS deleting key pair %s\n", n)
		} else {
			a.logger.Printf("ERROR deleting key pair %s: %s\n", n, err)
		}
	}

	return nil
}
