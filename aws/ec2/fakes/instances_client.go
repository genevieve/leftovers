package fakes

import "github.com/aws/aws-sdk-go/service/ec2"

type InstancesClient struct {
	DescribeInstancesCall struct {
		CallCount int
		Receives  struct {
			Input *ec2.DescribeInstancesInput
		}
		Returns struct {
			Output *ec2.DescribeInstancesOutput
			Error  error
		}
	}

	TerminateInstancesCall struct {
		CallCount int
		Receives  struct {
			Input *ec2.TerminateInstancesInput
		}
		Returns struct {
			Output *ec2.TerminateInstancesOutput
			Error  error
		}
	}

	DescribeAddressesCall struct {
		CallCount int
		Receives  struct {
			Input *ec2.DescribeAddressesInput
		}
		Returns struct {
			Output *ec2.DescribeAddressesOutput
			Error  error
		}
	}

	ReleaseAddressCall struct {
		CallCount int
		Receives  struct {
			Input *ec2.ReleaseAddressInput
		}
		Returns struct {
			Output *ec2.ReleaseAddressOutput
			Error  error
		}
	}
}

func (i *InstancesClient) DescribeInstances(input *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
	i.DescribeInstancesCall.CallCount++
	i.DescribeInstancesCall.Receives.Input = input

	return i.DescribeInstancesCall.Returns.Output, i.DescribeInstancesCall.Returns.Error
}

func (i *InstancesClient) TerminateInstances(input *ec2.TerminateInstancesInput) (*ec2.TerminateInstancesOutput, error) {
	i.TerminateInstancesCall.CallCount++
	i.TerminateInstancesCall.Receives.Input = input

	return i.TerminateInstancesCall.Returns.Output, i.TerminateInstancesCall.Returns.Error
}

func (i *InstancesClient) DescribeAddresses(input *ec2.DescribeAddressesInput) (*ec2.DescribeAddressesOutput, error) {
	i.DescribeAddressesCall.CallCount++
	i.DescribeAddressesCall.Receives.Input = input

	return i.DescribeAddressesCall.Returns.Output, i.DescribeAddressesCall.Returns.Error
}

func (i *InstancesClient) ReleaseAddress(input *ec2.ReleaseAddressInput) (*ec2.ReleaseAddressOutput, error) {
	i.ReleaseAddressCall.CallCount++
	i.ReleaseAddressCall.Receives.Input = input

	return i.ReleaseAddressCall.Returns.Output, i.ReleaseAddressCall.Returns.Error
}
