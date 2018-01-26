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
	delete := map[string]string{}

	loadBalancers, err := l.client.DescribeLoadBalancers(&awselbv2.DescribeLoadBalancersInput{})
	if err != nil {
		return delete, fmt.Errorf("Describing load balancers: %s", err)
	}

	for _, lb := range loadBalancers.LoadBalancers {
		n := *lb.LoadBalancerName

		if !strings.Contains(n, filter) {
			continue
		}

		proceed := l.logger.Prompt(fmt.Sprintf("Are you sure you want to delete load balancer %s?", n))
		if !proceed {
			continue
		}

		delete[n] = *lb.LoadBalancerArn
	}

	return delete, nil
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
