package elbv2

import (
	"fmt"
	awselbv2 "github.com/aws/aws-sdk-go/service/elbv2"

	"github.com/genevieve/leftovers/common"
)

//go:generate faux --interface loadBalancersClient --output fakes/load_balancers_client.go
type loadBalancersClient interface {
	DescribeLoadBalancers(*awselbv2.DescribeLoadBalancersInput) (*awselbv2.DescribeLoadBalancersOutput, error)
	DeleteLoadBalancer(*awselbv2.DeleteLoadBalancerInput) (*awselbv2.DeleteLoadBalancerOutput, error)

	DescribeTags(*awselbv2.DescribeTagsInput) (*awselbv2.DescribeTagsOutput, error)
}

type LoadBalancers struct {
	client loadBalancersClient
	logger logger
}

func NewLoadBalancers(client loadBalancersClient, logger logger) LoadBalancers {
	return LoadBalancers{
		client: client,
		logger: logger,
	}
}

func (l LoadBalancers) List(filter string, regex bool) ([]common.Deletable, error) {
	loadBalancers, err := l.client.DescribeLoadBalancers(&awselbv2.DescribeLoadBalancersInput{})
	if err != nil {
		return nil, fmt.Errorf("Describe ELBV2 Load Balancers: %s", err)

	}

	var resources []common.Deletable
	for _, lb := range loadBalancers.LoadBalancers {
		r := NewLoadBalancer(l.client, lb.LoadBalancerName, lb.LoadBalancerArn)

		if !common.ResourceMatches(r.Name(),  filter, regex) {
			continue
		}

		proceed := l.logger.PromptWithDetails(r.Type(), r.Name())
		if !proceed {
			continue
		}

		resources = append(resources, r)
	}

	return resources, nil
}

func (l LoadBalancers) Type() string {
	return "elbv2-load-balancer"
}
