package ec2

import (
	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
)

type KeyPair struct {
	client     keyPairsClient
	name       *string
	identifier string
}

func NewKeyPair(client keyPairsClient, name *string) KeyPair {
	return KeyPair{
		client:     client,
		name:       name,
		identifier: *name,
	}
}

func (k KeyPair) Delete() error {
	_, err := k.client.DeleteKeyPair(&awsec2.DeleteKeyPairInput{KeyName: k.name})
	return err
}
