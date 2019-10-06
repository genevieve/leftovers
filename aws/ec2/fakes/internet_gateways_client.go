package fakes

import (
	"sync"

	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
)

type InternetGatewaysClient struct {
	DeleteInternetGatewayCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DeleteInternetGatewayInput *awsec2.DeleteInternetGatewayInput
		}
		Returns struct {
			DeleteInternetGatewayOutput *awsec2.DeleteInternetGatewayOutput
			Error                       error
		}
		Stub func(*awsec2.DeleteInternetGatewayInput) (*awsec2.DeleteInternetGatewayOutput, error)
	}
	DescribeInternetGatewaysCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DescribeInternetGatewaysInput *awsec2.DescribeInternetGatewaysInput
		}
		Returns struct {
			DescribeInternetGatewaysOutput *awsec2.DescribeInternetGatewaysOutput
			Error                          error
		}
		Stub func(*awsec2.DescribeInternetGatewaysInput) (*awsec2.DescribeInternetGatewaysOutput, error)
	}
	DetachInternetGatewayCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DetachInternetGatewayInput *awsec2.DetachInternetGatewayInput
		}
		Returns struct {
			DetachInternetGatewayOutput *awsec2.DetachInternetGatewayOutput
			Error                       error
		}
		Stub func(*awsec2.DetachInternetGatewayInput) (*awsec2.DetachInternetGatewayOutput, error)
	}
}

func (f *InternetGatewaysClient) DeleteInternetGateway(param1 *awsec2.DeleteInternetGatewayInput) (*awsec2.DeleteInternetGatewayOutput, error) {
	f.DeleteInternetGatewayCall.Lock()
	defer f.DeleteInternetGatewayCall.Unlock()
	f.DeleteInternetGatewayCall.CallCount++
	f.DeleteInternetGatewayCall.Receives.DeleteInternetGatewayInput = param1
	if f.DeleteInternetGatewayCall.Stub != nil {
		return f.DeleteInternetGatewayCall.Stub(param1)
	}
	return f.DeleteInternetGatewayCall.Returns.DeleteInternetGatewayOutput, f.DeleteInternetGatewayCall.Returns.Error
}
func (f *InternetGatewaysClient) DescribeInternetGateways(param1 *awsec2.DescribeInternetGatewaysInput) (*awsec2.DescribeInternetGatewaysOutput, error) {
	f.DescribeInternetGatewaysCall.Lock()
	defer f.DescribeInternetGatewaysCall.Unlock()
	f.DescribeInternetGatewaysCall.CallCount++
	f.DescribeInternetGatewaysCall.Receives.DescribeInternetGatewaysInput = param1
	if f.DescribeInternetGatewaysCall.Stub != nil {
		return f.DescribeInternetGatewaysCall.Stub(param1)
	}
	return f.DescribeInternetGatewaysCall.Returns.DescribeInternetGatewaysOutput, f.DescribeInternetGatewaysCall.Returns.Error
}
func (f *InternetGatewaysClient) DetachInternetGateway(param1 *awsec2.DetachInternetGatewayInput) (*awsec2.DetachInternetGatewayOutput, error) {
	f.DetachInternetGatewayCall.Lock()
	defer f.DetachInternetGatewayCall.Unlock()
	f.DetachInternetGatewayCall.CallCount++
	f.DetachInternetGatewayCall.Receives.DetachInternetGatewayInput = param1
	if f.DetachInternetGatewayCall.Stub != nil {
		return f.DetachInternetGatewayCall.Stub(param1)
	}
	return f.DetachInternetGatewayCall.Returns.DetachInternetGatewayOutput, f.DetachInternetGatewayCall.Returns.Error
}
