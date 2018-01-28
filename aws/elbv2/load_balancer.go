package elbv2

import awselbv2 "github.com/aws/aws-sdk-go/service/elbv2"

type LoadBalancer struct {
	client     loadBalancersClient
	name       *string
	arn        *string
	identifier string
}

func NewLoadBalancer(client loadBalancersClient, name, arn *string) LoadBalancer {
	return LoadBalancer{
		client:     client,
		name:       name,
		arn:        arn,
		identifier: *name,
	}
}

func (l LoadBalancer) Delete() error {
	_, err := l.client.DeleteLoadBalancer(&awselbv2.DeleteLoadBalancerInput{
		LoadBalancerArn: l.arn,
	})
	return err
}
