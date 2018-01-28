package ec2

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
)

type keyPairsClient interface {
	DescribeKeyPairs(*awsec2.DescribeKeyPairsInput) (*awsec2.DescribeKeyPairsOutput, error)
	DeleteKeyPair(*awsec2.DeleteKeyPairInput) (*awsec2.DeleteKeyPairOutput, error)
}

type KeyPairs struct {
	client keyPairsClient
	logger logger
}

func NewKeyPairs(client keyPairsClient, logger logger) KeyPairs {
	return KeyPairs{
		client: client,
		logger: logger,
	}
}

func (k KeyPairs) List(filter string) (map[string]string, error) {
	keyPairs, err := k.list(filter)
	if err != nil {
		return nil, err
	}

	delete := map[string]string{}
	for _, key := range keyPairs {
		delete[*key.name] = ""
	}

	return delete, nil
}

func (k KeyPairs) list(filter string) ([]KeyPair, error) {
	keyPairs, err := k.client.DescribeKeyPairs(&awsec2.DescribeKeyPairsInput{})
	if err != nil {
		return nil, fmt.Errorf("Describing key pairs: %s", err)
	}

	var resources []KeyPair
	for _, key := range keyPairs.KeyPairs {
		resource := NewKeyPair(k.client, key.KeyName)

		if !strings.Contains(resource.identifier, filter) {
			continue
		}

		proceed := k.logger.Prompt(fmt.Sprintf("Are you sure you want to delete key pair %s?", resource.identifier))
		if !proceed {
			continue
		}

		resources = append(resources, resource)
	}

	return resources, nil
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
