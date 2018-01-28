package ec2

import (
	"fmt"
	"strings"

	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
)

type Vpc struct {
	client     vpcsClient
	id         *string
	identifier string
}

func NewVpc(client vpcsClient, id *string, tags []*awsec2.Tag) Vpc {
	identifier := *id

	var extra []string
	for _, t := range tags {
		extra = append(extra, fmt.Sprintf("%s:%s", *t.Key, *t.Value))
	}

	if len(extra) > 0 {
		identifier = fmt.Sprintf("%s (%s)", *id, strings.Join(extra, ","))
	}

	return Vpc{
		client:     client,
		id:         id,
		identifier: identifier,
	}
}

func (v Vpc) Delete() error {
	_, err := v.client.DeleteVpc(&awsec2.DeleteVpcInput{VpcId: v.id})
	return err
}
