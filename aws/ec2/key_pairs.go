package ec2

import (
	"fmt"
	awsec2 "github.com/aws/aws-sdk-go/service/ec2"

	"github.com/genevieve/leftovers/common"
)

//go:generate faux --interface keyPairsClient --output fakes/key_pairs_client.go
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

func (k KeyPairs) List(filter string, regex bool) ([]common.Deletable, error) {
	keyPairs, err := k.client.DescribeKeyPairs(&awsec2.DescribeKeyPairsInput{})
	if err != nil {
		return nil, fmt.Errorf("Describing EC2 Key Pairs: %s", err)
	}

	var resources []common.Deletable
	for _, key := range keyPairs.KeyPairs {
		r := NewKeyPair(k.client, key.KeyName)

		if !common.ResourceMatches(r.Name(),  filter, regex) {
			continue
		}

		proceed := k.logger.PromptWithDetails(r.Type(), r.Name())
		if !proceed {
			continue
		}

		resources = append(resources, r)
	}

	return resources, nil
}

func (k KeyPairs) Type() string {
	return "ec2-key-pair"
}
