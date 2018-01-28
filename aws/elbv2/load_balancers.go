package elbv2

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
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

func (l LoadBalancers) List(filter string) (map[string]string, error) {
	loadBalancers, err := l.list(filter)
	if err != nil {
		return nil, err
	}

	delete := map[string]string{}
	for _, lb := range loadBalancers {
		delete[lb.identifier] = *lb.arn
	}

	return delete, nil
}

func (l LoadBalancers) list(filter string) ([]LoadBalancer, error) {
	loadBalancers, err := l.client.DescribeLoadBalancers(&awselbv2.DescribeLoadBalancersInput{})
	if err != nil {
		return nil, fmt.Errorf("Describing load balancers: %s", err)
	}

	var resources []LoadBalancer
	for _, lb := range loadBalancers.LoadBalancers {
		resource := NewLoadBalancer(l.client, lb.LoadBalancerName, lb.LoadBalancerArn)

		if !strings.Contains(resource.identifier, filter) {
			continue
		}

		proceed := l.logger.Prompt(fmt.Sprintf("Are you sure you want to delete load balancer %s?", resource.identifier))
		if !proceed {
			continue
		}

		resources = append(resources, resource)
	}

	return resources, nil
}

func (l LoadBalancers) Delete(loadBalancers map[string]string) error {
	for name, arn := range loadBalancers {
		_, err := l.client.DeleteLoadBalancer(&awselbv2.DeleteLoadBalancerInput{
			LoadBalancerArn: aws.String(arn),
		})

		if err == nil {
			l.logger.Printf("SUCCESS deleting load balancer %s\n", name)
		} else {
			l.logger.Printf("ERROR deleting load balancer %s: %s\n", name, err)
		}
	}

	return nil
}
