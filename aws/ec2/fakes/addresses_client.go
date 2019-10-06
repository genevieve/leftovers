package fakes

import (
	"sync"

	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
)

type AddressesClient struct {
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
}

func (f *AddressesClient) DescribeAddresses(param1 *awsec2.DescribeAddressesInput) (*awsec2.DescribeAddressesOutput, error) {
	f.DescribeAddressesCall.Lock()
	defer f.DescribeAddressesCall.Unlock()
	f.DescribeAddressesCall.CallCount++
	f.DescribeAddressesCall.Receives.DescribeAddressesInput = param1
	if f.DescribeAddressesCall.Stub != nil {
		return f.DescribeAddressesCall.Stub(param1)
	}
	return f.DescribeAddressesCall.Returns.DescribeAddressesOutput, f.DescribeAddressesCall.Returns.Error
}
func (f *AddressesClient) ReleaseAddress(param1 *awsec2.ReleaseAddressInput) (*awsec2.ReleaseAddressOutput, error) {
	f.ReleaseAddressCall.Lock()
	defer f.ReleaseAddressCall.Unlock()
	f.ReleaseAddressCall.CallCount++
	f.ReleaseAddressCall.Receives.ReleaseAddressInput = param1
	if f.ReleaseAddressCall.Stub != nil {
		return f.ReleaseAddressCall.Stub(param1)
	}
	return f.ReleaseAddressCall.Returns.ReleaseAddressOutput, f.ReleaseAddressCall.Returns.Error
}
