package fakes

import "github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumes"

type VolumesClient struct {
	ListCall struct {
		CallCount int
		Returns   struct {
			Volumes []volumes.Volume
			Error   error
		}
	}
	DeleteCall struct {
		CallCount int
		Receives  struct {
			VolumeID string
		}
		Returns struct {
			Error error
		}
	}
}

func (v *VolumesClient) Delete(volumeID string) error {
	v.DeleteCall.CallCount++
	v.DeleteCall.Receives.VolumeID = volumeID

	return v.DeleteCall.Returns.Error
}

func (v *VolumesClient) List() ([]volumes.Volume, error) {
	v.ListCall.CallCount++

	return v.ListCall.Returns.Volumes, v.ListCall.Returns.Error
}
