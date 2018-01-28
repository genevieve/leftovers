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

func (d Addresses) List(filter string) (map[string]string, error) {
	addresses, err := d.list(filter)
	if err != nil {
		return nil, err
	}

	delete := map[string]string{}
	for _, a := range addresses {
		delete[*a.publicIp] = *a.allocationId
	}

	return delete, nil
}

func (d Addresses) list(filter string) ([]Address, error) {
	addresses, err := d.client.DescribeAddresses(&awsec2.DescribeAddressesInput{})
	if err != nil {
		return nil, fmt.Errorf("Describing addresses: %s", err)
	}

	var resources []Address
	for _, a := range addresses.Addresses {
		resource := NewAddress(d.client, a.PublicIp, a.AllocationId)

		if d.inUse(a) {
			continue
		}

		proceed := d.logger.Prompt(fmt.Sprintf("Are you sure you want to release address %s?", resource.identifier))
		if !proceed {
			continue
		}

		resources = append(resources, resource)
	}

	return resources, nil
}

func (a Addresses) Delete(addresses map[string]string) error {
	var resources []Address
	for ip, id := range addresses {
		resources = append(resources, NewAddress(a.client, &ip, &id))
	}

	return a.cleanup(resources)
}

func (a Addresses) cleanup(resources []Address) error {
	for _, resource := range resources {
		err := resource.Delete()

		if err == nil {
			a.logger.Printf("SUCCESS releasing address %s\n", resource.identifier)
		} else {
			a.logger.Printf("ERROR releasing address %s: %s\n", resource.identifier, err)
		}
	}

	return nil
}

func (d Addresses) inUse(a *awsec2.Address) bool {
	return a.InstanceId != nil && *a.InstanceId != ""
}
