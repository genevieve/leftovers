package ec2

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
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

func (v Vpcs) List(filter string) (map[string]string, error) {
	delete := map[string]string{}

	vpcs, err := v.client.DescribeVpcs(&awsec2.DescribeVpcsInput{})
	if err != nil {
		return delete, fmt.Errorf("Describing vpcs: %s", err)
	}

	for _, vpc := range vpcs.Vpcs {
		if *vpc.IsDefault {
			continue
		}

		vpcId := *vpc.VpcId

		n := v.clearerName(vpcId, vpc.Tags)

		if n != vpcId && !strings.Contains(n, filter) {
			continue
		}

		proceed := v.logger.Prompt(fmt.Sprintf("Are you sure you want to delete vpc %s?", n))
		if !proceed {
			continue
		}

		delete[n] = vpcId
	}

	return delete, nil
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
