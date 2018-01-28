package elb

import (
	awselb "github.com/aws/aws-sdk-go/service/elb"
)

type LoadBalancer struct {
	client     loadBalancersClient
	name       *string
	identifier string
}

func NewLoadBalancer(client loadBalancersClient, name *string) LoadBalancer {
	return LoadBalancer{
		client:     client,
		name:       name,
		identifier: *name,
	}
}

func (l LoadBalancer) Delete() error {
	_, err := l.client.DeleteLoadBalancer(&awselb.DeleteLoadBalancerInput{
		LoadBalancerName: l.name,
	})
	return err
}
