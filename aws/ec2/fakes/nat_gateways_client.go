package fakes

import (
	"sync"

	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
)

type NatGatewaysClient struct {
	DeleteNatGatewayCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DeleteNatGatewayInput *awsec2.DeleteNatGatewayInput
		}
		Returns struct {
			DeleteNatGatewayOutput *awsec2.DeleteNatGatewayOutput
			Error                  error
		}
		Stub func(*awsec2.DeleteNatGatewayInput) (*awsec2.DeleteNatGatewayOutput, error)
	}
	DescribeNatGatewaysCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DescribeNatGatewaysInput *awsec2.DescribeNatGatewaysInput
		}
		Returns struct {
			DescribeNatGatewaysOutput *awsec2.DescribeNatGatewaysOutput
			Error                     error
		}
		Stub func(*awsec2.DescribeNatGatewaysInput) (*awsec2.DescribeNatGatewaysOutput, error)
	}
}

func (f *NatGatewaysClient) DeleteNatGateway(param1 *awsec2.DeleteNatGatewayInput) (*awsec2.DeleteNatGatewayOutput, error) {
	f.DeleteNatGatewayCall.Lock()
	defer f.DeleteNatGatewayCall.Unlock()
	f.DeleteNatGatewayCall.CallCount++
	f.DeleteNatGatewayCall.Receives.DeleteNatGatewayInput = param1
	if f.DeleteNatGatewayCall.Stub != nil {
		return f.DeleteNatGatewayCall.Stub(param1)
	}
	return f.DeleteNatGatewayCall.Returns.DeleteNatGatewayOutput, f.DeleteNatGatewayCall.Returns.Error
}
func (f *NatGatewaysClient) DescribeNatGateways(param1 *awsec2.DescribeNatGatewaysInput) (*awsec2.DescribeNatGatewaysOutput, error) {
	f.DescribeNatGatewaysCall.Lock()
	defer f.DescribeNatGatewaysCall.Unlock()
	f.DescribeNatGatewaysCall.CallCount++
	f.DescribeNatGatewaysCall.Receives.DescribeNatGatewaysInput = param1
	if f.DescribeNatGatewaysCall.Stub != nil {
		return f.DescribeNatGatewaysCall.Stub(param1)
	}
	return f.DescribeNatGatewaysCall.Returns.DescribeNatGatewaysOutput, f.DescribeNatGatewaysCall.Returns.Error
}
