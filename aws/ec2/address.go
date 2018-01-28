package ec2

import (
	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
)

type Address struct {
	client       addressesClient
	publicIp     *string
	allocationId *string
	identifier   string
}

func NewAddress(client addressesClient, publicIp, allocationId *string) Address {
	return Address{
		client:       client,
		publicIp:     publicIp,
		allocationId: allocationId,
		identifier:   *publicIp,
	}
}

func (a Address) Delete() error {
	_, err := a.client.ReleaseAddress(&awsec2.ReleaseAddressInput{
		AllocationId: a.allocationId,
	})
	return err
}
