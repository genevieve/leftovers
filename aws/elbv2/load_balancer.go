package elbv2

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
