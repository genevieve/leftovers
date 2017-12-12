package fakes

import "github.com/aws/aws-sdk-go/service/ec2"

type VolumesClient struct {
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

func (e *VolumesClient) DescribeVolumes(input *ec2.DescribeVolumesInput) (*ec2.DescribeVolumesOutput, error) {
	e.DescribeVolumesCall.CallCount++
	e.DescribeVolumesCall.Receives.Input = input

	return e.DescribeVolumesCall.Returns.Output, e.DescribeVolumesCall.Returns.Error
}

func (e *VolumesClient) DeleteVolume(input *ec2.DeleteVolumeInput) (*ec2.DeleteVolumeOutput, error) {
	e.DeleteVolumeCall.CallCount++
	e.DeleteVolumeCall.Receives.Input = input

	return e.DeleteVolumeCall.Returns.Output, e.DeleteVolumeCall.Returns.Error
}
