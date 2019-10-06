package fakes

import (
	"sync"

	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
)

type VolumesClient struct {
	DeleteVolumeCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DeleteVolumeInput *awsec2.DeleteVolumeInput
		}
		Returns struct {
			DeleteVolumeOutput *awsec2.DeleteVolumeOutput
			Error              error
		}
		Stub func(*awsec2.DeleteVolumeInput) (*awsec2.DeleteVolumeOutput, error)
	}
	DescribeVolumesCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DescribeVolumesInput *awsec2.DescribeVolumesInput
		}
		Returns struct {
			DescribeVolumesOutput *awsec2.DescribeVolumesOutput
			Error                 error
		}
		Stub func(*awsec2.DescribeVolumesInput) (*awsec2.DescribeVolumesOutput, error)
	}
}

func (f *VolumesClient) DeleteVolume(param1 *awsec2.DeleteVolumeInput) (*awsec2.DeleteVolumeOutput, error) {
	f.DeleteVolumeCall.Lock()
	defer f.DeleteVolumeCall.Unlock()
	f.DeleteVolumeCall.CallCount++
	f.DeleteVolumeCall.Receives.DeleteVolumeInput = param1
	if f.DeleteVolumeCall.Stub != nil {
		return f.DeleteVolumeCall.Stub(param1)
	}
	return f.DeleteVolumeCall.Returns.DeleteVolumeOutput, f.DeleteVolumeCall.Returns.Error
}
func (f *VolumesClient) DescribeVolumes(param1 *awsec2.DescribeVolumesInput) (*awsec2.DescribeVolumesOutput, error) {
	f.DescribeVolumesCall.Lock()
	defer f.DescribeVolumesCall.Unlock()
	f.DescribeVolumesCall.CallCount++
	f.DescribeVolumesCall.Receives.DescribeVolumesInput = param1
	if f.DescribeVolumesCall.Stub != nil {
		return f.DescribeVolumesCall.Stub(param1)
	}
	return f.DescribeVolumesCall.Returns.DescribeVolumesOutput, f.DescribeVolumesCall.Returns.Error
}
