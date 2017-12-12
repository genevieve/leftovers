package fakes

import "github.com/aws/aws-sdk-go/service/ec2"

type VpcClient struct {
	DescribeVpcsCall struct {
		CallCount int
		Receives  struct {
			Input *ec2.DescribeVpcsInput
		}
		Returns struct {
			Output *ec2.DescribeVpcsOutput
			Error  error
		}
	}

	DeleteVpcCall struct {
		CallCount int
		Receives  struct {
			Input *ec2.DeleteVpcInput
		}
		Returns struct {
			Output *ec2.DeleteVpcOutput
			Error  error
		}
	}
}

func (e *VpcClient) DescribeVpcs(input *ec2.DescribeVpcsInput) (*ec2.DescribeVpcsOutput, error) {
	e.DescribeVpcsCall.CallCount++
	e.DescribeVpcsCall.Receives.Input = input

	return e.DescribeVpcsCall.Returns.Output, e.DescribeVpcsCall.Returns.Error
}

func (e *VpcClient) DeleteVpc(input *ec2.DeleteVpcInput) (*ec2.DeleteVpcOutput, error) {
	e.DeleteVpcCall.CallCount++
	e.DeleteVpcCall.Receives.Input = input

	return e.DeleteVpcCall.Returns.Output, e.DeleteVpcCall.Returns.Error
}
