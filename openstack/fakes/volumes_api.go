package fakes

import (
	"sync"

	"github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumes"
	"github.com/gophercloud/gophercloud/pagination"
)

type VolumesAPI struct {
	DeleteVolumeCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Id string
		}
		Returns struct {
			Error error
		}
		Stub func(string) error
	}
	GetVolumesPagerCall struct {
		sync.Mutex
		CallCount int
		Returns   struct {
			Pager pagination.Pager
		}
		Stub func() pagination.Pager
	}
	PageToVolumesCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Page pagination.Page
		}
		Returns struct {
			VolumeSlice []volumes.Volume
			Error       error
		}
		Stub func(pagination.Page) ([]volumes.Volume, error)
	}
	PagerToPageCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Pager pagination.Pager
		}
		Returns struct {
			Page  pagination.Page
			Error error
		}
		Stub func(pagination.Pager) (pagination.Page, error)
	}
}

func (f *VolumesAPI) DeleteVolume(param1 string) error {
	f.DeleteVolumeCall.Lock()
	defer f.DeleteVolumeCall.Unlock()
	f.DeleteVolumeCall.CallCount++
	f.DeleteVolumeCall.Receives.Id = param1
	if f.DeleteVolumeCall.Stub != nil {
		return f.DeleteVolumeCall.Stub(param1)
	}
	return f.DeleteVolumeCall.Returns.Error
}
func (f *VolumesAPI) GetVolumesPager() pagination.Pager {
	f.GetVolumesPagerCall.Lock()
	defer f.GetVolumesPagerCall.Unlock()
	f.GetVolumesPagerCall.CallCount++
	if f.GetVolumesPagerCall.Stub != nil {
		return f.GetVolumesPagerCall.Stub()
	}
	return f.GetVolumesPagerCall.Returns.Pager
}
func (f *VolumesAPI) PageToVolumes(param1 pagination.Page) ([]volumes.Volume, error) {
	f.PageToVolumesCall.Lock()
	defer f.PageToVolumesCall.Unlock()
	f.PageToVolumesCall.CallCount++
	f.PageToVolumesCall.Receives.Page = param1
	if f.PageToVolumesCall.Stub != nil {
		return f.PageToVolumesCall.Stub(param1)
	}
	return f.PageToVolumesCall.Returns.VolumeSlice, f.PageToVolumesCall.Returns.Error
}
func (f *VolumesAPI) PagerToPage(param1 pagination.Pager) (pagination.Page, error) {
	f.PagerToPageCall.Lock()
	defer f.PagerToPageCall.Unlock()
	f.PagerToPageCall.CallCount++
	f.PagerToPageCall.Receives.Pager = param1
	if f.PagerToPageCall.Stub != nil {
		return f.PagerToPageCall.Stub(param1)
	}
	return f.PagerToPageCall.Returns.Page, f.PagerToPageCall.Returns.Error
}
