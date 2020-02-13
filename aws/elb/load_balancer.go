package elb

import (
	"fmt"
	"strings"

	awselb "github.com/aws/aws-sdk-go/service/elb"
)

type LoadBalancer struct {
	client     loadBalancersClient
	name       *string
	identifier string
	rtype      string
}

func NewLoadBalancer(client loadBalancersClient, name *string) LoadBalancer {
	identifier := *name

	tagsOutput, err := client.DescribeTags(&awselb.DescribeTagsInput{LoadBalancerNames: []*string{name}})
	if err == nil && tagsOutput != nil && len(tagsOutput.TagDescriptions) == 1 {
		tags := tagsOutput.TagDescriptions[0].Tags

		var extra []string
		for _, t := range tags {
			extra = append(extra, fmt.Sprintf("%s:%s", *t.Key, *t.Value))
		}

		if len(extra) > 0 {
			identifier = fmt.Sprintf("%s (%s)", *name, strings.Join(extra, ", "))
		}
	}

	return LoadBalancer{
		client:     client,
		name:       name,
		identifier: identifier,
		rtype:      "ELB Load Balancer",
	}
}

func (l LoadBalancer) Delete() error {
	_, err := l.client.DeleteLoadBalancer(&awselb.DeleteLoadBalancerInput{
		LoadBalancerName: l.name,
	})

	if err != nil {
		return fmt.Errorf("Delete: %s", err)
	}

	return nil
}

func (l LoadBalancer) Name() string {
	return l.identifier
}

func (l LoadBalancer) Type() string {
	return l.rtype
}
