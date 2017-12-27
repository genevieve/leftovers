package fakes

import "github.com/aws/aws-sdk-go/service/elbv2"

type LoadBalancersClient struct {
	DescribeLoadBalancersCall struct {
		CallCount int
		Receives  struct {
			Input *elbv2.DescribeLoadBalancersInput
		}
		Returns struct {
			Output *elbv2.DescribeLoadBalancersOutput
			Error  error
		}
	}

	DeleteLoadBalancerCall struct {
		CallCount int
		Receives  struct {
			Input *elbv2.DeleteLoadBalancerInput
		}
		Returns struct {
			Output *elbv2.DeleteLoadBalancerOutput
			Error  error
		}
	}
}

func (e *LoadBalancersClient) DescribeLoadBalancers(input *elbv2.DescribeLoadBalancersInput) (*elbv2.DescribeLoadBalancersOutput, error) {
	e.DescribeLoadBalancersCall.CallCount++
	e.DescribeLoadBalancersCall.Receives.Input = input

	return e.DescribeLoadBalancersCall.Returns.Output, e.DescribeLoadBalancersCall.Returns.Error
}

func (e *LoadBalancersClient) DeleteLoadBalancer(input *elbv2.DeleteLoadBalancerInput) (*elbv2.DeleteLoadBalancerOutput, error) {
	e.DeleteLoadBalancerCall.CallCount++
	e.DeleteLoadBalancerCall.Receives.Input = input

	return e.DeleteLoadBalancerCall.Returns.Output, e.DeleteLoadBalancerCall.Returns.Error
}
