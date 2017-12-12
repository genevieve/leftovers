package fakes

import "github.com/aws/aws-sdk-go/service/elb"

type LoadBalancersClient struct {
	DescribeLoadBalancersCall struct {
		CallCount int
		Receives  struct {
			Input *elb.DescribeLoadBalancersInput
		}
		Returns struct {
			Output *elb.DescribeLoadBalancersOutput
			Error  error
		}
	}

	DeleteLoadBalancerCall struct {
		CallCount int
		Receives  struct {
			Input *elb.DeleteLoadBalancerInput
		}
		Returns struct {
			Output *elb.DeleteLoadBalancerOutput
			Error  error
		}
	}
}

func (e *LoadBalancersClient) DescribeLoadBalancers(input *elb.DescribeLoadBalancersInput) (*elb.DescribeLoadBalancersOutput, error) {
	e.DescribeLoadBalancersCall.CallCount++
	e.DescribeLoadBalancersCall.Receives.Input = input

	return e.DescribeLoadBalancersCall.Returns.Output, e.DescribeLoadBalancersCall.Returns.Error
}

func (e *LoadBalancersClient) DeleteLoadBalancer(input *elb.DeleteLoadBalancerInput) (*elb.DeleteLoadBalancerOutput, error) {
	e.DeleteLoadBalancerCall.CallCount++
	e.DeleteLoadBalancerCall.Receives.Input = input

	return e.DeleteLoadBalancerCall.Returns.Output, e.DeleteLoadBalancerCall.Returns.Error
}
