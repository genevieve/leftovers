package elbv2

import (
	"fmt"
	"strings"

	awselbv2 "github.com/aws/aws-sdk-go/service/elbv2"
)

type LoadBalancer struct {
	client     loadBalancersClient
	name       *string
	arn        *string
	identifier string
	rtype      string
}

func NewLoadBalancer(client loadBalancersClient, name, arn *string) LoadBalancer {
	identifier := *name

	tagsOutput, err := client.DescribeTags(&awselbv2.DescribeTagsInput{ResourceArns: []*string{arn}})
	if err == nil && tagsOutput != nil && len(tagsOutput.TagDescriptions) == 1 {
		tags := tagsOutput.TagDescriptions[0].Tags

		var extra []string
		for _, t := range tags {
			extra = append(extra, fmt.Sprintf("%s:%s", *t.Key, *t.Value))
		}

		if len(extra) > 0 {
			identifier = fmt.Sprintf("%s (%s)", *name, strings.Join(extra, ", "))
		}
	}

	return LoadBalancer{
		client:     client,
		name:       name,
		arn:        arn,
		identifier: identifier,
		rtype:      "ELBV2 Load Balancer",
	}
}

func (l LoadBalancer) Delete() error {
	_, err := l.client.DeleteLoadBalancer(&awselbv2.DeleteLoadBalancerInput{
		LoadBalancerArn: l.arn,
	})
	if err != nil {
		return fmt.Errorf("Delete: %s", err)
	}

	return nil
}

func (l LoadBalancer) Name() string {
	return l.identifier
}

func (l LoadBalancer) Type() string {
	return l.rtype
}
