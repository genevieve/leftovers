package ec2

import (
	"fmt"
	"strings"

	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
)

type networkInterfacesClient interface {
	DescribeNetworkInterfaces(*awsec2.DescribeNetworkInterfacesInput) (*awsec2.DescribeNetworkInterfacesOutput, error)
	DeleteNetworkInterface(*awsec2.DeleteNetworkInterfaceInput) (*awsec2.DeleteNetworkInterfaceOutput, error)
}

type NetworkInterfaces struct {
	client networkInterfacesClient
	logger logger
}

func NewNetworkInterfaces(client networkInterfacesClient, logger logger) NetworkInterfaces {
	return NetworkInterfaces{
		client: client,
		logger: logger,
	}
}

func (e NetworkInterfaces) List(filter string) (map[string]string, error) {
	networkInterfaces, err := e.list(filter)
	if err != nil {
		return nil, err
	}

	delete := map[string]string{}
	for _, n := range networkInterfaces {
		delete[n.identifier] = *n.id
	}

	return delete, nil
}

func (e NetworkInterfaces) list(filter string) ([]NetworkInterface, error) {
	networkInterfaces, err := e.client.DescribeNetworkInterfaces(&awsec2.DescribeNetworkInterfacesInput{})
	if err != nil {
		return nil, fmt.Errorf("Describing network interfaces: %s", err)
	}

	var resources []NetworkInterface
	for _, i := range networkInterfaces.NetworkInterfaces {
		resource := NewNetworkInterface(e.client, i.NetworkInterfaceId, i.TagSet)

		if !strings.Contains(resource.identifier, filter) {
			continue
		}

		proceed := e.logger.Prompt(fmt.Sprintf("Are you sure you want to delete network interface %s?", resource.identifier))
		if !proceed {
			continue
		}

		resources = append(resources, resource)
	}

	return resources, nil
}

func (n NetworkInterfaces) Delete(networkInterfaces map[string]string) error {
	var resources []NetworkInterface
	for _, id := range networkInterfaces {
		resources = append(resources, NewNetworkInterface(n.client, &id, []*awsec2.Tag{}))
	}

	return n.cleanup(resources)
}

func (n NetworkInterfaces) cleanup(resources []NetworkInterface) error {
	for _, resource := range resources {
		err := resource.Delete()

		if err == nil {
			n.logger.Printf("SUCCESS deleting network interface %s\n", resource.identifier)
		} else {
			n.logger.Printf("ERROR deleting network interface %s: %s\n", resource.identifier, err)
		}
	}

	return nil
}
