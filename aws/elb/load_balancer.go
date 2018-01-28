package elb

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
