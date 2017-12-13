package fakes

import "github.com/aws/aws-sdk-go/service/ec2"

type InternetGatewaysClient struct {
	DescribeInternetGatewaysCall struct {
		CallCount int
		Receives  struct {
			Input *ec2.DescribeInternetGatewaysInput
		}
		Returns struct {
			Output *ec2.DescribeInternetGatewaysOutput
			Error  error
		}
	}

	DetachInternetGatewayCall struct {
		CallCount int
		Receives  struct {
			Input *ec2.DetachInternetGatewayInput
		}
		Returns struct {
			Output *ec2.DetachInternetGatewayOutput
			Error  error
		}
	}

	DeleteInternetGatewayCall struct {
		CallCount int
		Receives  struct {
			Input *ec2.DeleteInternetGatewayInput
		}
		Returns struct {
			Output *ec2.DeleteInternetGatewayOutput
			Error  error
		}
	}
}

func (i *InternetGatewaysClient) DescribeInternetGateways(input *ec2.DescribeInternetGatewaysInput) (*ec2.DescribeInternetGatewaysOutput, error) {
	i.DescribeInternetGatewaysCall.CallCount++
	i.DescribeInternetGatewaysCall.Receives.Input = input

	return i.DescribeInternetGatewaysCall.Returns.Output, i.DescribeInternetGatewaysCall.Returns.Error
}

func (i *InternetGatewaysClient) DetachInternetGateway(input *ec2.DetachInternetGatewayInput) (*ec2.DetachInternetGatewayOutput, error) {
	i.DetachInternetGatewayCall.CallCount++
	i.DetachInternetGatewayCall.Receives.Input = input

	return i.DetachInternetGatewayCall.Returns.Output, i.DetachInternetGatewayCall.Returns.Error
}

func (i *InternetGatewaysClient) DeleteInternetGateway(input *ec2.DeleteInternetGatewayInput) (*ec2.DeleteInternetGatewayOutput, error) {
	i.DeleteInternetGatewayCall.CallCount++
	i.DeleteInternetGatewayCall.Receives.Input = input

	return i.DeleteInternetGatewayCall.Returns.Output, i.DeleteInternetGatewayCall.Returns.Error
}
