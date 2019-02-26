package fakes

import "github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumes"

type VolumesLister struct {
	ListCall struct {
		CallCount int
		Returns   struct {
			Volumes []volumes.Volume
			Error   error
		}
	}
}

func (v *VolumesLister) List() ([]volumes.Volume, error) {
	v.ListCall.CallCount++

	return v.ListCall.Returns.Volumes, v.ListCall.Returns.Error
}
