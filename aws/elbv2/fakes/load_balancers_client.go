package fakes

import (
	"sync"

	awselbv2 "github.com/aws/aws-sdk-go/service/elbv2"
)

type LoadBalancersClient struct {
	DeleteLoadBalancerCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DeleteLoadBalancerInput *awselbv2.DeleteLoadBalancerInput
		}
		Returns struct {
			DeleteLoadBalancerOutput *awselbv2.DeleteLoadBalancerOutput
			Error                    error
		}
		Stub func(*awselbv2.DeleteLoadBalancerInput) (*awselbv2.DeleteLoadBalancerOutput, error)
	}
	DescribeLoadBalancersCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DescribeLoadBalancersInput *awselbv2.DescribeLoadBalancersInput
		}
		Returns struct {
			DescribeLoadBalancersOutput *awselbv2.DescribeLoadBalancersOutput
			Error                       error
		}
		Stub func(*awselbv2.DescribeLoadBalancersInput) (*awselbv2.DescribeLoadBalancersOutput, error)
	}
}

func (f *LoadBalancersClient) DeleteLoadBalancer(param1 *awselbv2.DeleteLoadBalancerInput) (*awselbv2.DeleteLoadBalancerOutput, error) {
	f.DeleteLoadBalancerCall.Lock()
	defer f.DeleteLoadBalancerCall.Unlock()
	f.DeleteLoadBalancerCall.CallCount++
	f.DeleteLoadBalancerCall.Receives.DeleteLoadBalancerInput = param1
	if f.DeleteLoadBalancerCall.Stub != nil {
		return f.DeleteLoadBalancerCall.Stub(param1)
	}
	return f.DeleteLoadBalancerCall.Returns.DeleteLoadBalancerOutput, f.DeleteLoadBalancerCall.Returns.Error
}
func (f *LoadBalancersClient) DescribeLoadBalancers(param1 *awselbv2.DescribeLoadBalancersInput) (*awselbv2.DescribeLoadBalancersOutput, error) {
	f.DescribeLoadBalancersCall.Lock()
	defer f.DescribeLoadBalancersCall.Unlock()
	f.DescribeLoadBalancersCall.CallCount++
	f.DescribeLoadBalancersCall.Receives.DescribeLoadBalancersInput = param1
	if f.DescribeLoadBalancersCall.Stub != nil {
		return f.DescribeLoadBalancersCall.Stub(param1)
	}
	return f.DescribeLoadBalancersCall.Returns.DescribeLoadBalancersOutput, f.DescribeLoadBalancersCall.Returns.Error
}
