package fakes

import (
	"sync"

	"github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumes"
)

type VolumesClient struct {
	DeleteCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			VolumeID string
		}
		Returns struct {
			Error error
		}
		Stub func(string) error
	}
	ListCall struct {
		sync.Mutex
		CallCount int
		Returns   struct {
			VolumeSlice []volumes.Volume
			Error       error
		}
		Stub func() ([]volumes.Volume, error)
	}
}

func (f *VolumesClient) Delete(param1 string) error {
	f.DeleteCall.Lock()
	defer f.DeleteCall.Unlock()
	f.DeleteCall.CallCount++
	f.DeleteCall.Receives.VolumeID = param1
	if f.DeleteCall.Stub != nil {
		return f.DeleteCall.Stub(param1)
	}
	return f.DeleteCall.Returns.Error
}
func (f *VolumesClient) List() ([]volumes.Volume, error) {
	f.ListCall.Lock()
	defer f.ListCall.Unlock()
	f.ListCall.CallCount++
	if f.ListCall.Stub != nil {
		return f.ListCall.Stub()
	}
	return f.ListCall.Returns.VolumeSlice, f.ListCall.Returns.Error
}
