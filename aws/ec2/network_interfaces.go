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

func (e NetworkInterfaces) Delete() error {
	networkInterfaces, err := e.client.DescribeNetworkInterfaces(&awsec2.DescribeNetworkInterfacesInput{})
	if err != nil {
		return fmt.Errorf("Describing network interfaces: %s", err)
	}

	for _, i := range networkInterfaces.NetworkInterfaces {
		n := e.clearerName(*i.NetworkInterfaceId, i.TagSet)

		proceed := e.logger.Prompt(fmt.Sprintf("Are you sure you want to delete network interface %s?", n))
		if !proceed {
			continue
		}

		_, err := e.client.DeleteNetworkInterface(&awsec2.DeleteNetworkInterfaceInput{NetworkInterfaceId: i.NetworkInterfaceId})
		if err == nil {
			e.logger.Printf("SUCCESS deleting network interface %s\n", n)
		} else {
			e.logger.Printf("ERROR deleting network interface %s: %s\n", n, err)
		}
	}

	return nil
}

func (e NetworkInterfaces) clearerName(id string, tags []*awsec2.Tag) string {
	extra := []string{}
	for _, t := range tags {
		extra = append(extra, fmt.Sprintf("%s:%s", *t.Key, *t.Value))
	}

	if len(extra) > 0 {
		return fmt.Sprintf("%s (%s)", id, strings.Join(extra, ", "))
	}

	return id
}
