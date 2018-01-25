package elbv2

import (
	"fmt"
	"strings"

	awselbv2 "github.com/aws/aws-sdk-go/service/elbv2"
)

type loadBalancersClient interface {
	DescribeLoadBalancers(*awselbv2.DescribeLoadBalancersInput) (*awselbv2.DescribeLoadBalancersOutput, error)
	DeleteLoadBalancer(*awselbv2.DeleteLoadBalancerInput) (*awselbv2.DeleteLoadBalancerOutput, error)
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

func (o LoadBalancers) Delete(filter string) error {
	loadBalancers, err := o.client.DescribeLoadBalancers(&awselbv2.DescribeLoadBalancersInput{})
	if err != nil {
		return fmt.Errorf("Describing load balancers: %s", err)
	}

	for _, l := range loadBalancers.LoadBalancers {
		n := *l.LoadBalancerName

		if !strings.Contains(n, filter) {
			continue
		}

		proceed := o.logger.Prompt(fmt.Sprintf("Are you sure you want to delete load balancer %s?", n))
		if !proceed {
			continue
		}

		_, err := o.client.DeleteLoadBalancer(&awselbv2.DeleteLoadBalancerInput{LoadBalancerArn: l.LoadBalancerArn})
		if err == nil {
			o.logger.Printf("SUCCESS deleting load balancer %s\n", n)
		} else {
			o.logger.Printf("ERROR deleting load balancer %s: %s\n", n, err)
		}
	}

	return nil
}
