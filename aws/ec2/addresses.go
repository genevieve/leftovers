package ec2

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
)

type addressesClient interface {
	DescribeAddresses(*awsec2.DescribeAddressesInput) (*awsec2.DescribeAddressesOutput, error)
	ReleaseAddress(*awsec2.ReleaseAddressInput) (*awsec2.ReleaseAddressOutput, error)
}

type Addresses struct {
	client addressesClient
	logger logger
}

func NewAddresses(client addressesClient, logger logger) Addresses {
	return Addresses{
		client: client,
		logger: logger,
	}
}

func (d Addresses) List(filter string) (map[string]string, error) {
	delete := map[string]string{}

	addresses, err := d.client.DescribeAddresses(&awsec2.DescribeAddressesInput{})
	if err != nil {
		return delete, fmt.Errorf("Describing addresses: %s", err)
	}

	for _, a := range addresses.Addresses {
		if d.inUse(a) {
			continue
		}

		n := *a.PublicIp

		proceed := d.logger.Prompt(fmt.Sprintf("Are you sure you want to release address %s?", n))
		if !proceed {
			continue
		}

		delete[n] = *a.AllocationId
	}

	return delete, nil
}

func (a Addresses) Delete(addresses map[string]string) error {
	for ip, id := range addresses {
		_, err := a.client.ReleaseAddress(&awsec2.ReleaseAddressInput{AllocationId: aws.String(id)})

		if err == nil {
			a.logger.Printf("SUCCESS releasing address %s\n", ip)
		} else {
			a.logger.Printf("ERROR releasing address %s: %s\n", ip, err)
		}
	}

	return nil
}

func (d Addresses) inUse(a *awsec2.Address) bool {
	return a.InstanceId != nil && *a.InstanceId != ""
}
