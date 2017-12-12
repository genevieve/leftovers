package elb

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
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

func (o LoadBalancers) Delete() error {
	loadBalancers, err := o.client.DescribeLoadBalancers(&awselb.DescribeLoadBalancersInput{})
	if err != nil {
		return fmt.Errorf("Describing load balancers: %s", err)
	}

	for _, l := range loadBalancers.LoadBalancerDescriptions {
		n := *l.LoadBalancerName

		proceed := o.logger.Prompt(fmt.Sprintf("Are you sure you want to delete load balancer %s?", n))
		if !proceed {
			continue
		}

		_, err := o.client.DeleteLoadBalancer(&awselb.DeleteLoadBalancerInput{LoadBalancerName: aws.String(n)})
		if err == nil {
			o.logger.Printf("SUCCESS deleting load balancer %s\n", n)
		} else {
			o.logger.Printf("ERROR deleting load balancer %s: %s\n", n, err)
		}
	}

	return nil
}
