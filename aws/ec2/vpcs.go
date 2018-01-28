package ec2

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
)

type vpcsClient interface {
	DescribeVpcs(*awsec2.DescribeVpcsInput) (*awsec2.DescribeVpcsOutput, error)
	DeleteVpc(*awsec2.DeleteVpcInput) (*awsec2.DeleteVpcOutput, error)
}

type Vpcs struct {
	client   vpcsClient
	logger   logger
	routes   routeTables
	subnets  subnets
	gateways internetGateways
}

func NewVpcs(client vpcsClient,
	logger logger,
	routes routeTables,
	subnets subnets,
	gateways internetGateways) Vpcs {
	return Vpcs{
		client:   client,
		logger:   logger,
		routes:   routes,
		subnets:  subnets,
		gateways: gateways,
	}
}

func (v Vpcs) List(filter string) (map[string]string, error) {
	vpcs, err := v.list(filter)
	if err != nil {
		return nil, err
	}

	delete := map[string]string{}
	for _, vpc := range vpcs {
		delete[vpc.name] = *vpc.id
	}

	return delete, nil
}

func (v Vpcs) list(filter string) ([]Vpc, error) {
	output, err := v.client.DescribeVpcs(&awsec2.DescribeVpcsInput{})
	if err != nil {
		return nil, fmt.Errorf("Describing vpcs: %s", err)
	}

	var resources []Vpc
	for _, vpc := range output.Vpcs {
		resource := NewVpc(v.client, vpc.VpcId, vpc.Tags)

		if *vpc.IsDefault {
			continue
		}

		if !strings.Contains(resource.name, filter) {
			continue
		}

		proceed := v.logger.Prompt(fmt.Sprintf("Are you sure you want to delete vpc %s?", resource.name))
		if !proceed {
			continue
		}

		resources = append(resources, resource)
	}

	return resources, nil
}

func (v Vpcs) Delete(vpcs map[string]string) error {
	for name, id := range vpcs {
		err := v.routes.Delete(id)
		if err != nil {
			return fmt.Errorf("Deleting routes for %s: %s", name, err)
		}

		err = v.subnets.Delete(id)
		if err != nil {
			return fmt.Errorf("Deleting subnets for %s: %s", name, err)
		}

		err = v.gateways.Delete(id)
		if err != nil {
			return fmt.Errorf("Deleting internet gateways for %s: %s", name, err)
		}

		_, err = v.client.DeleteVpc(&awsec2.DeleteVpcInput{VpcId: aws.String(id)})

		if err == nil {
			v.logger.Printf("SUCCESS deleting vpc %s\n", name)
		} else {
			v.logger.Printf("ERROR deleting vpc %s: %s\n", name, err)
		}
	}

	return nil
}
