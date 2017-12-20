package ec2

import (
	"fmt"
	"strings"

	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
)

type vpcClient interface {
	DescribeVpcs(*awsec2.DescribeVpcsInput) (*awsec2.DescribeVpcsOutput, error)
	DeleteVpc(*awsec2.DeleteVpcInput) (*awsec2.DeleteVpcOutput, error)
}

type Vpcs struct {
	client   vpcClient
	logger   logger
	routes   routeTables
	subnets  subnets
	gateways internetGateways
}

func NewVpcs(client vpcClient,
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

func (p Vpcs) Delete() error {
	vpcs, err := p.client.DescribeVpcs(&awsec2.DescribeVpcsInput{})
	if err != nil {
		return fmt.Errorf("Describing vpcs: %s", err)
	}

	for _, v := range vpcs.Vpcs {
		if *v.IsDefault {
			continue
		}

		vpcId := *v.VpcId

		n := p.clearerName(vpcId, v.Tags)

		proceed := p.logger.Prompt(fmt.Sprintf("Are you sure you want to delete vpc %s?", n))
		if !proceed {
			continue
		}

		if err := p.routes.Delete(vpcId); err != nil {
			return fmt.Errorf("Deleting routes for %s: %s", n, err)
		}

		if err := p.subnets.Delete(vpcId); err != nil {
			return fmt.Errorf("Deleting subnets for %s: %s", n, err)
		}

		if err := p.gateways.Delete(vpcId); err != nil {
			return fmt.Errorf("Deleting internet gateways for %s: %s", n, err)
		}

		_, err := p.client.DeleteVpc(&awsec2.DeleteVpcInput{VpcId: v.VpcId})
		if err == nil {
			p.logger.Printf("SUCCESS deleting vpc %s\n", n)
		} else {
			p.logger.Printf("ERROR deleting vpc %s: %s\n", n, err)
		}
	}

	return nil
}

func (p Vpcs) clearerName(id string, tags []*awsec2.Tag) string {
	extra := []string{}
	for _, t := range tags {
		extra = append(extra, fmt.Sprintf("%s:%s", *t.Key, *t.Value))
	}

	if len(extra) > 0 {
		return fmt.Sprintf("%s (%s)", id, strings.Join(extra, ","))
	}

	return id
}
