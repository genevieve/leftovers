package elb

import (
	awselb "github.com/aws/aws-sdk-go/service/elb"
)

type elbClient interface {
	DescribeLoadBalancers(*awselb.DescribeLoadBalancersInput) (*awselb.DescribeLoadBalancersOutput, error)
	DeleteLoadBalancer(*awselb.DeleteLoadBalancerInput) (*awselb.DeleteLoadBalancerOutput, error)
}
