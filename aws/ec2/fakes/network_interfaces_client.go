package fakes

import "github.com/aws/aws-sdk-go/service/ec2"

type NetworkInterfaceClient struct {
	DescribeNetworkInterfacesCall struct {
		CallCount int
		Receives  struct {
			Input *ec2.DescribeNetworkInterfacesInput
		}
		Returns struct {
			Output *ec2.DescribeNetworkInterfacesOutput
			Error  error
		}
	}

	DeleteNetworkInterfaceCall struct {
		CallCount int
		Receives  struct {
			Input *ec2.DeleteNetworkInterfaceInput
		}
		Returns struct {
			Output *ec2.DeleteNetworkInterfaceOutput
			Error  error
		}
	}
}

func (e *NetworkInterfaceClient) DescribeNetworkInterfaces(input *ec2.DescribeNetworkInterfacesInput) (*ec2.DescribeNetworkInterfacesOutput, error) {
	e.DescribeNetworkInterfacesCall.CallCount++
	e.DescribeNetworkInterfacesCall.Receives.Input = input

	return e.DescribeNetworkInterfacesCall.Returns.Output, e.DescribeNetworkInterfacesCall.Returns.Error
}

func (e *NetworkInterfaceClient) DeleteNetworkInterface(input *ec2.DeleteNetworkInterfaceInput) (*ec2.DeleteNetworkInterfaceOutput, error) {
	e.DeleteNetworkInterfaceCall.CallCount++
	e.DeleteNetworkInterfaceCall.Receives.Input = input

	return e.DeleteNetworkInterfaceCall.Returns.Output, e.DeleteNetworkInterfaceCall.Returns.Error
}
