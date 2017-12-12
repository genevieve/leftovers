package ec2

import (
	"fmt"

	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
)

type vpcClient interface {
	DescribeVpcs(*awsec2.DescribeVpcsInput) (*awsec2.DescribeVpcsOutput, error)
	DeleteVpc(*awsec2.DeleteVpcInput) (*awsec2.DeleteVpcOutput, error)
}

type Vpcs struct {
	client vpcClient
	logger logger
}

func NewVpcs(client vpcClient, logger logger) Vpcs {
	return Vpcs{
		client: client,
		logger: logger,
	}
}

func (p Vpcs) Delete() error {
	vpcs, err := p.client.DescribeVpcs(&awsec2.DescribeVpcsInput{})
	if err != nil {
		return fmt.Errorf("Describing vpcs: %s", err)
	}

	for _, v := range vpcs.Vpcs {
		vpcId := *v.VpcId
		n := vpcName(v)

		proceed := p.logger.Prompt(fmt.Sprintf("Are you sure you want to delete vpc %s%s?", vpcId, n))
		if !proceed {
			continue
		}

		_, err := p.client.DeleteVpc(&awsec2.DeleteVpcInput{VpcId: v.VpcId})
		if err == nil {
			p.logger.Printf("SUCCESS deleting vpc %s%s\n", vpcId, n)
		} else {
			p.logger.Printf("ERROR deleting vpc %s%s: %s\n", vpcId, n, err)
		}
	}

	return nil
}

func vpcName(v *awsec2.Vpc) string {
	for _, t := range v.Tags {
		if *t.Key == "Name" {
			return fmt.Sprintf("/%s", *t.Value)
		}
	}
	return ""
}
