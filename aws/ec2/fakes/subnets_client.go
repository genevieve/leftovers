package fakes

import "github.com/aws/aws-sdk-go/service/ec2"

type SubnetsClient struct {
	DescribeSubnetsCall struct {
		CallCount int
		Receives  struct {
			Input *ec2.DescribeSubnetsInput
		}
		Returns struct {
			Output *ec2.DescribeSubnetsOutput
			Error  error
		}
	}

	DeleteSubnetCall struct {
		CallCount int
		Receives  struct {
			Input *ec2.DeleteSubnetInput
		}
		Returns struct {
			Output *ec2.DeleteSubnetOutput
			Error  error
		}
	}
}

func (i *SubnetsClient) DescribeSubnets(input *ec2.DescribeSubnetsInput) (*ec2.DescribeSubnetsOutput, error) {
	i.DescribeSubnetsCall.CallCount++
	i.DescribeSubnetsCall.Receives.Input = input

	return i.DescribeSubnetsCall.Returns.Output, i.DescribeSubnetsCall.Returns.Error
}

func (i *SubnetsClient) DeleteSubnet(input *ec2.DeleteSubnetInput) (*ec2.DeleteSubnetOutput, error) {
	i.DeleteSubnetCall.CallCount++
	i.DeleteSubnetCall.Receives.Input = input

	return i.DeleteSubnetCall.Returns.Output, i.DeleteSubnetCall.Returns.Error
}
