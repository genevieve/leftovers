package ec2

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
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

func (k KeyPairs) List(filter string) (map[string]string, error) {
	delete := map[string]string{}

	keyPairs, err := k.client.DescribeKeyPairs(&awsec2.DescribeKeyPairsInput{})
	if err != nil {
		return delete, fmt.Errorf("Describing key pairs: %s", err)
	}

	for _, key := range keyPairs.KeyPairs {
		n := *key.KeyName

		if !strings.Contains(n, filter) {
			continue
		}

		proceed := k.logger.Prompt(fmt.Sprintf("Are you sure you want to delete key pair %s?", n))
		if !proceed {
			continue
		}

		delete[n] = ""
	}

	return delete, nil
}

func (k KeyPairs) Delete(keyPairs map[string]string) error {
	for name, _ := range keyPairs {
		_, err := k.client.DeleteKeyPair(&awsec2.DeleteKeyPairInput{KeyName: aws.String(name)})

		if err == nil {
			k.logger.Printf("SUCCESS deleting key pair %s\n", name)
		} else {
			k.logger.Printf("ERROR deleting key pair %s: %s\n", name, err)
		}
	}

	return nil
}
