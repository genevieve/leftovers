package fakes

import "github.com/aws/aws-sdk-go/service/ec2"

type NatGatewaysClient struct {
	DescribeNatGatewaysCall struct {
		CallCount int
		Receives  struct {
			Input *ec2.DescribeNatGatewaysInput
		}
		Returns struct {
			Output *ec2.DescribeNatGatewaysOutput
			Error  error
		}
	}

	DeleteNatGatewayCall struct {
		CallCount int
		Receives  struct {
			Input *ec2.DeleteNatGatewayInput
		}
		Returns struct {
			Output *ec2.DeleteNatGatewayOutput
			Error  error
		}
	}
}

func (e *NatGatewaysClient) DescribeNatGateways(input *ec2.DescribeNatGatewaysInput) (*ec2.DescribeNatGatewaysOutput, error) {
	e.DescribeNatGatewaysCall.CallCount++
	e.DescribeNatGatewaysCall.Receives.Input = input

	return e.DescribeNatGatewaysCall.Returns.Output, e.DescribeNatGatewaysCall.Returns.Error
}

func (e *NatGatewaysClient) DeleteNatGateway(input *ec2.DeleteNatGatewayInput) (*ec2.DeleteNatGatewayOutput, error) {
	e.DeleteNatGatewayCall.CallCount++
	e.DeleteNatGatewayCall.Receives.Input = input

	return e.DeleteNatGatewayCall.Returns.Output, e.DeleteNatGatewayCall.Returns.Error
}
