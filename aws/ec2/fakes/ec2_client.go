package fakes

import (
	"github.com/aws/aws-sdk-go/service/ec2"
)

type EC2Client struct {
	DescribeKeyPairsCall struct {
		CallCount int
		Receives  struct {
			Input *ec2.DescribeKeyPairsInput
		}
		Returns struct {
			Output *ec2.DescribeKeyPairsOutput
			Error  error
		}
	}

	DeleteKeyPairCall struct {
		CallCount int
		Receives  struct {
			Input *ec2.DeleteKeyPairInput
		}
		Returns struct {
			Output *ec2.DeleteKeyPairOutput
			Error  error
		}
	}

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

	DescribeTagsCall struct {
		CallCount int
		Receives  struct {
			Input *ec2.DescribeTagsInput
		}
		Returns struct {
			Output *ec2.DescribeTagsOutput
			Error  error
		}
	}

	DeleteTagsCall struct {
		CallCount int
		Receives  struct {
			Input *ec2.DeleteTagsInput
		}
		Returns struct {
			Output *ec2.DeleteTagsOutput
			Error  error
		}
	}

	DescribeVolumesCall struct {
		CallCount int
		Receives  struct {
			Input *ec2.DescribeVolumesInput
		}
		Returns struct {
			Output *ec2.DescribeVolumesOutput
			Error  error
		}
	}

	DeleteVolumeCall struct {
		CallCount int
		Receives  struct {
			Input *ec2.DeleteVolumeInput
		}
		Returns struct {
			Output *ec2.DeleteVolumeOutput
			Error  error
		}
	}
}

func (e *EC2Client) DescribeKeyPairs(input *ec2.DescribeKeyPairsInput) (*ec2.DescribeKeyPairsOutput, error) {
	e.DescribeKeyPairsCall.CallCount++
	e.DescribeKeyPairsCall.Receives.Input = input

	return e.DescribeKeyPairsCall.Returns.Output, e.DescribeKeyPairsCall.Returns.Error
}

func (e *EC2Client) DeleteKeyPair(input *ec2.DeleteKeyPairInput) (*ec2.DeleteKeyPairOutput, error) {
	e.DeleteKeyPairCall.CallCount++
	e.DeleteKeyPairCall.Receives.Input = input

	return e.DeleteKeyPairCall.Returns.Output, e.DeleteKeyPairCall.Returns.Error
}

func (e *EC2Client) DescribeInstances(input *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
	e.DescribeInstancesCall.CallCount++
	e.DescribeInstancesCall.Receives.Input = input

	return e.DescribeInstancesCall.Returns.Output, e.DescribeInstancesCall.Returns.Error
}

func (e *EC2Client) TerminateInstances(input *ec2.TerminateInstancesInput) (*ec2.TerminateInstancesOutput, error) {
	e.TerminateInstancesCall.CallCount++
	e.TerminateInstancesCall.Receives.Input = input

	return e.TerminateInstancesCall.Returns.Output, e.TerminateInstancesCall.Returns.Error
}

func (e *EC2Client) DescribeTags(input *ec2.DescribeTagsInput) (*ec2.DescribeTagsOutput, error) {
	e.DescribeTagsCall.CallCount++
	e.DescribeTagsCall.Receives.Input = input

	return e.DescribeTagsCall.Returns.Output, e.DescribeTagsCall.Returns.Error
}

func (e *EC2Client) DeleteTags(input *ec2.DeleteTagsInput) (*ec2.DeleteTagsOutput, error) {
	e.DeleteTagsCall.CallCount++
	e.DeleteTagsCall.Receives.Input = input

	return e.DeleteTagsCall.Returns.Output, e.DeleteTagsCall.Returns.Error
}

func (e *EC2Client) DescribeVolumes(input *ec2.DescribeVolumesInput) (*ec2.DescribeVolumesOutput, error) {
	e.DescribeVolumesCall.CallCount++
	e.DescribeVolumesCall.Receives.Input = input

	return e.DescribeVolumesCall.Returns.Output, e.DescribeVolumesCall.Returns.Error
}

func (e *EC2Client) DeleteVolume(input *ec2.DeleteVolumeInput) (*ec2.DeleteVolumeOutput, error) {
	e.DeleteVolumeCall.CallCount++
	e.DeleteVolumeCall.Receives.Input = input

	return e.DeleteVolumeCall.Returns.Output, e.DeleteVolumeCall.Returns.Error
}
