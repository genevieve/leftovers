package elb

import (
	"fmt"
	awselb "github.com/aws/aws-sdk-go/service/elb"

	"github.com/genevieve/leftovers/common"
)

//go:generate faux --interface loadBalancersClient --output fakes/load_balancers_client.go
type loadBalancersClient interface {
	DescribeLoadBalancers(*awselb.DescribeLoadBalancersInput) (*awselb.DescribeLoadBalancersOutput, error)
	DeleteLoadBalancer(*awselb.DeleteLoadBalancerInput) (*awselb.DeleteLoadBalancerOutput, error)

	DescribeTags(*awselb.DescribeTagsInput) (*awselb.DescribeTagsOutput, error)
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
	loadBalancers, err := l.client.DescribeLoadBalancers(&awselb.DescribeLoadBalancersInput{})
	if err != nil {
		return nil, fmt.Errorf("Describe ELB Load Balancers: %s", err)
	}

	var resources []common.Deletable
	for _, lb := range loadBalancers.LoadBalancerDescriptions {
		r := NewLoadBalancer(l.client, lb.LoadBalancerName)

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
	return "elb-load-balancer"
}
