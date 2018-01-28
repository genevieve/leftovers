package elb

import (
	"fmt"
	"strings"

	awselb "github.com/aws/aws-sdk-go/service/elb"
)

type loadBalancersClient interface {
	DescribeLoadBalancers(*awselb.DescribeLoadBalancersInput) (*awselb.DescribeLoadBalancersOutput, error)
	DeleteLoadBalancer(*awselb.DeleteLoadBalancerInput) (*awselb.DeleteLoadBalancerOutput, error)
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
		delete[lb.identifier] = ""
	}

	return delete, nil
}

func (l LoadBalancers) list(filter string) ([]LoadBalancer, error) {
	loadBalancers, err := l.client.DescribeLoadBalancers(&awselb.DescribeLoadBalancersInput{})
	if err != nil {
		return nil, fmt.Errorf("Describing load balancers: %s", err)
	}

	var resources []LoadBalancer
	for _, lb := range loadBalancers.LoadBalancerDescriptions {
		resource := NewLoadBalancer(l.client, lb.LoadBalancerName)

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
	var resources []LoadBalancer
	for name, _ := range loadBalancers {
		resources = append(resources, NewLoadBalancer(l.client, &name))
	}

	return l.cleanup(resources)
}

func (l LoadBalancers) cleanup(resources []LoadBalancer) error {
	for _, resource := range resources {
		err := resource.Delete()
		if err == nil {
			l.logger.Printf("SUCCESS deleting load balancer %s\n", resource.identifier)
		} else {
			l.logger.Printf("ERROR deleting load balancer %s: %s\n", resource.identifier, err)
		}
	}

	return nil
}
