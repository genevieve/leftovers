package fakes

import (
	"sync"

	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
)

type NetworkInterfacesClient struct {
	DeleteNetworkInterfaceCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DeleteNetworkInterfaceInput *awsec2.DeleteNetworkInterfaceInput
		}
		Returns struct {
			DeleteNetworkInterfaceOutput *awsec2.DeleteNetworkInterfaceOutput
			Error                        error
		}
		Stub func(*awsec2.DeleteNetworkInterfaceInput) (*awsec2.DeleteNetworkInterfaceOutput, error)
	}
	DescribeNetworkInterfacesCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DescribeNetworkInterfacesInput *awsec2.DescribeNetworkInterfacesInput
		}
		Returns struct {
			DescribeNetworkInterfacesOutput *awsec2.DescribeNetworkInterfacesOutput
			Error                           error
		}
		Stub func(*awsec2.DescribeNetworkInterfacesInput) (*awsec2.DescribeNetworkInterfacesOutput, error)
	}
}

func (f *NetworkInterfacesClient) DeleteNetworkInterface(param1 *awsec2.DeleteNetworkInterfaceInput) (*awsec2.DeleteNetworkInterfaceOutput, error) {
	f.DeleteNetworkInterfaceCall.Lock()
	defer f.DeleteNetworkInterfaceCall.Unlock()
	f.DeleteNetworkInterfaceCall.CallCount++
	f.DeleteNetworkInterfaceCall.Receives.DeleteNetworkInterfaceInput = param1
	if f.DeleteNetworkInterfaceCall.Stub != nil {
		return f.DeleteNetworkInterfaceCall.Stub(param1)
	}
	return f.DeleteNetworkInterfaceCall.Returns.DeleteNetworkInterfaceOutput, f.DeleteNetworkInterfaceCall.Returns.Error
}
func (f *NetworkInterfacesClient) DescribeNetworkInterfaces(param1 *awsec2.DescribeNetworkInterfacesInput) (*awsec2.DescribeNetworkInterfacesOutput, error) {
	f.DescribeNetworkInterfacesCall.Lock()
	defer f.DescribeNetworkInterfacesCall.Unlock()
	f.DescribeNetworkInterfacesCall.CallCount++
	f.DescribeNetworkInterfacesCall.Receives.DescribeNetworkInterfacesInput = param1
	if f.DescribeNetworkInterfacesCall.Stub != nil {
		return f.DescribeNetworkInterfacesCall.Stub(param1)
	}
	return f.DescribeNetworkInterfacesCall.Returns.DescribeNetworkInterfacesOutput, f.DescribeNetworkInterfacesCall.Returns.Error
}
