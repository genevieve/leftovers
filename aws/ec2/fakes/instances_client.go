package fakes

import (
	"sync"

	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
)

type InstancesClient struct {
	DescribeAddressesCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DescribeAddressesInput *awsec2.DescribeAddressesInput
		}
		Returns struct {
			DescribeAddressesOutput *awsec2.DescribeAddressesOutput
			Error                   error
		}
		Stub func(*awsec2.DescribeAddressesInput) (*awsec2.DescribeAddressesOutput, error)
	}
	DescribeInstancesCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DescribeInstancesInput *awsec2.DescribeInstancesInput
		}
		Returns struct {
			DescribeInstancesOutput *awsec2.DescribeInstancesOutput
			Error                   error
		}
		Stub func(*awsec2.DescribeInstancesInput) (*awsec2.DescribeInstancesOutput, error)
	}
	ReleaseAddressCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			ReleaseAddressInput *awsec2.ReleaseAddressInput
		}
		Returns struct {
			ReleaseAddressOutput *awsec2.ReleaseAddressOutput
			Error                error
		}
		Stub func(*awsec2.ReleaseAddressInput) (*awsec2.ReleaseAddressOutput, error)
	}
	TerminateInstancesCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			TerminateInstancesInput *awsec2.TerminateInstancesInput
		}
		Returns struct {
			TerminateInstancesOutput *awsec2.TerminateInstancesOutput
			Error                    error
		}
		Stub func(*awsec2.TerminateInstancesInput) (*awsec2.TerminateInstancesOutput, error)
	}
}

func (f *InstancesClient) DescribeAddresses(param1 *awsec2.DescribeAddressesInput) (*awsec2.DescribeAddressesOutput, error) {
	f.DescribeAddressesCall.Lock()
	defer f.DescribeAddressesCall.Unlock()
	f.DescribeAddressesCall.CallCount++
	f.DescribeAddressesCall.Receives.DescribeAddressesInput = param1
	if f.DescribeAddressesCall.Stub != nil {
		return f.DescribeAddressesCall.Stub(param1)
	}
	return f.DescribeAddressesCall.Returns.DescribeAddressesOutput, f.DescribeAddressesCall.Returns.Error
}
func (f *InstancesClient) DescribeInstances(param1 *awsec2.DescribeInstancesInput) (*awsec2.DescribeInstancesOutput, error) {
	f.DescribeInstancesCall.Lock()
	defer f.DescribeInstancesCall.Unlock()
	f.DescribeInstancesCall.CallCount++
	f.DescribeInstancesCall.Receives.DescribeInstancesInput = param1
	if f.DescribeInstancesCall.Stub != nil {
		return f.DescribeInstancesCall.Stub(param1)
	}
	return f.DescribeInstancesCall.Returns.DescribeInstancesOutput, f.DescribeInstancesCall.Returns.Error
}
func (f *InstancesClient) ReleaseAddress(param1 *awsec2.ReleaseAddressInput) (*awsec2.ReleaseAddressOutput, error) {
	f.ReleaseAddressCall.Lock()
	defer f.ReleaseAddressCall.Unlock()
	f.ReleaseAddressCall.CallCount++
	f.ReleaseAddressCall.Receives.ReleaseAddressInput = param1
	if f.ReleaseAddressCall.Stub != nil {
		return f.ReleaseAddressCall.Stub(param1)
	}
	return f.ReleaseAddressCall.Returns.ReleaseAddressOutput, f.ReleaseAddressCall.Returns.Error
}
func (f *InstancesClient) TerminateInstances(param1 *awsec2.TerminateInstancesInput) (*awsec2.TerminateInstancesOutput, error) {
	f.TerminateInstancesCall.Lock()
	defer f.TerminateInstancesCall.Unlock()
	f.TerminateInstancesCall.CallCount++
	f.TerminateInstancesCall.Receives.TerminateInstancesInput = param1
	if f.TerminateInstancesCall.Stub != nil {
		return f.TerminateInstancesCall.Stub(param1)
	}
	return f.TerminateInstancesCall.Returns.TerminateInstancesOutput, f.TerminateInstancesCall.Returns.Error
}
