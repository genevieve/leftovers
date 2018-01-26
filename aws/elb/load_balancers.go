package elb

import (
	"fmt"
	"strings"

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

func (l LoadBalancers) List(filter string) (map[string]string, error) {
	delete := map[string]string{}

	loadBalancers, err := l.client.DescribeLoadBalancers(&awselb.DescribeLoadBalancersInput{})
	if err != nil {
		return delete, fmt.Errorf("Describing load balancers: %s", err)
	}

	for _, lb := range loadBalancers.LoadBalancerDescriptions {
		n := *lb.LoadBalancerName

		if !strings.Contains(n, filter) {
			continue
		}

		proceed := l.logger.Prompt(fmt.Sprintf("Are you sure you want to delete load balancer %s?", n))
		if !proceed {
			continue
		}

		delete[n] = ""
	}

	return delete, nil
}

func (l LoadBalancers) Delete(loadBalancers map[string]string) error {
	for name, _ := range loadBalancers {
		_, err := l.client.DeleteLoadBalancer(&awselb.DeleteLoadBalancerInput{LoadBalancerName: aws.String(name)})

		if err == nil {
			l.logger.Printf("SUCCESS deleting load balancer %s\n", name)
		} else {
			l.logger.Printf("ERROR deleting load balancer %s: %s\n", name, err)
		}
	}

	return nil
}
