package ec2

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
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
	delete := map[string]string{}

	networkInterfaces, err := e.client.DescribeNetworkInterfaces(&awsec2.DescribeNetworkInterfacesInput{})
	if err != nil {
		return delete, fmt.Errorf("Describing network interfaces: %s", err)
	}

	for _, i := range networkInterfaces.NetworkInterfaces {
		n := e.clearerName(*i.NetworkInterfaceId, i.TagSet)

		if !strings.Contains(n, filter) {
			continue
		}

		proceed := e.logger.Prompt(fmt.Sprintf("Are you sure you want to delete network interface %s?", n))
		if !proceed {
			continue
		}

		delete[n] = *i.NetworkInterfaceId
	}

	return delete, nil
}

func (n NetworkInterfaces) Delete(networkInterfaces map[string]string) error {
	for name, id := range networkInterfaces {
		_, err := n.client.DeleteNetworkInterface(&awsec2.DeleteNetworkInterfaceInput{
			NetworkInterfaceId: aws.String(id),
		})

		if err == nil {
			n.logger.Printf("SUCCESS deleting network interface %s\n", name)
		} else {
			n.logger.Printf("ERROR deleting network interface %s: %s\n", name, err)
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
