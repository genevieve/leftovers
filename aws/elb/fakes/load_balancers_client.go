package fakes

import (
	"sync"

	awselb "github.com/aws/aws-sdk-go/service/elb"
)

type LoadBalancersClient struct {
	DeleteLoadBalancerCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DeleteLoadBalancerInput *awselb.DeleteLoadBalancerInput
		}
		Returns struct {
			DeleteLoadBalancerOutput *awselb.DeleteLoadBalancerOutput
			Error                    error
		}
		Stub func(*awselb.DeleteLoadBalancerInput) (*awselb.DeleteLoadBalancerOutput, error)
	}
	DescribeLoadBalancersCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DescribeLoadBalancersInput *awselb.DescribeLoadBalancersInput
		}
		Returns struct {
			DescribeLoadBalancersOutput *awselb.DescribeLoadBalancersOutput
			Error                       error
		}
		Stub func(*awselb.DescribeLoadBalancersInput) (*awselb.DescribeLoadBalancersOutput, error)
	}
	DescribeTagsCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DescribeTagsInput *awselb.DescribeTagsInput
		}
		Returns struct {
			DescribeTagsOutput *awselb.DescribeTagsOutput
			Error              error
		}
		Stub func(*awselb.DescribeTagsInput) (*awselb.DescribeTagsOutput, error)
	}
}

func (f *LoadBalancersClient) DeleteLoadBalancer(param1 *awselb.DeleteLoadBalancerInput) (*awselb.DeleteLoadBalancerOutput, error) {
	f.DeleteLoadBalancerCall.Lock()
	defer f.DeleteLoadBalancerCall.Unlock()
	f.DeleteLoadBalancerCall.CallCount++
	f.DeleteLoadBalancerCall.Receives.DeleteLoadBalancerInput = param1
	if f.DeleteLoadBalancerCall.Stub != nil {
		return f.DeleteLoadBalancerCall.Stub(param1)
	}
	return f.DeleteLoadBalancerCall.Returns.DeleteLoadBalancerOutput, f.DeleteLoadBalancerCall.Returns.Error
}
func (f *LoadBalancersClient) DescribeLoadBalancers(param1 *awselb.DescribeLoadBalancersInput) (*awselb.DescribeLoadBalancersOutput, error) {
	f.DescribeLoadBalancersCall.Lock()
	defer f.DescribeLoadBalancersCall.Unlock()
	f.DescribeLoadBalancersCall.CallCount++
	f.DescribeLoadBalancersCall.Receives.DescribeLoadBalancersInput = param1
	if f.DescribeLoadBalancersCall.Stub != nil {
		return f.DescribeLoadBalancersCall.Stub(param1)
	}
	return f.DescribeLoadBalancersCall.Returns.DescribeLoadBalancersOutput, f.DescribeLoadBalancersCall.Returns.Error
}
func (f *LoadBalancersClient) DescribeTags(param1 *awselb.DescribeTagsInput) (*awselb.DescribeTagsOutput, error) {
	f.DescribeTagsCall.Lock()
	defer f.DescribeTagsCall.Unlock()
	f.DescribeTagsCall.CallCount++
	f.DescribeTagsCall.Receives.DescribeTagsInput = param1
	if f.DescribeTagsCall.Stub != nil {
		return f.DescribeTagsCall.Stub(param1)
	}
	return f.DescribeTagsCall.Returns.DescribeTagsOutput, f.DescribeTagsCall.Returns.Error
}
