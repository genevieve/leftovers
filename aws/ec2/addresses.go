package ec2

import (
	"fmt"

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

func (d Addresses) Delete(filter string) error {
	addresses, err := d.client.DescribeAddresses(&awsec2.DescribeAddressesInput{})
	if err != nil {
		return fmt.Errorf("Describing addresses: %s", err)
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

		_, err := d.client.ReleaseAddress(&awsec2.ReleaseAddressInput{
			AllocationId: a.AllocationId,
		})
		if err == nil {
			d.logger.Printf("SUCCESS releasing address %s\n", n)
		} else {
			d.logger.Printf("ERROR releasing address %s: %s\n", n, err)
		}
	}

	return nil
}

func (d Addresses) inUse(a *awsec2.Address) bool {
	return a.InstanceId != nil && *a.InstanceId != ""
}
