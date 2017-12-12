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
}

func (e *InstancesClient) DescribeInstances(input *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
	e.DescribeInstancesCall.CallCount++
	e.DescribeInstancesCall.Receives.Input = input

	return e.DescribeInstancesCall.Returns.Output, e.DescribeInstancesCall.Returns.Error
}

func (e *InstancesClient) TerminateInstances(input *ec2.TerminateInstancesInput) (*ec2.TerminateInstancesOutput, error) {
	e.TerminateInstancesCall.CallCount++
	e.TerminateInstancesCall.Receives.Input = input

	return e.TerminateInstancesCall.Returns.Output, e.TerminateInstancesCall.Returns.Error
}
