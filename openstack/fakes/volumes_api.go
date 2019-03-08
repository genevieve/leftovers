package fakes

import (
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumes"
	"github.com/gophercloud/gophercloud/pagination"
)

type VolumesAPI struct {
	PagerToPageCall struct {
		CallCount int
		Receives  struct {
			Pager pagination.Pager
		}
		Returns struct {
			Page  pagination.Page
			Error error
		}
	}

	PageToVolumesCall struct {
		CallCount int
		Receives  struct {
			Page pagination.Page
		}
		Returns struct {
			Volumes []volumes.Volume
			Error   error
		}
	}

	GetVolumesPagerCall struct {
		CallCount int
		Returns   struct {
			Pager pagination.Pager
		}
	}

	DeleteVolumeCall struct {
		CallCount int
		Receives  struct {
			VolumeID string
		}
		ReceivesForCall []struct {
			VolumeID string
		}
		Returns struct {
			Error error
		}
	}
}

func (v *VolumesAPI) PagerToPage(pager pagination.Pager) (pagination.Page, error) {
	v.PagerToPageCall.CallCount++
	v.PagerToPageCall.Receives.Pager = pager

	return v.PagerToPageCall.Returns.Page, v.PagerToPageCall.Returns.Error
}

func (v *VolumesAPI) PageToVolumes(page pagination.Page) ([]volumes.Volume, error) {
	v.PageToVolumesCall.CallCount++
	v.PageToVolumesCall.Receives.Page = page

	return v.PageToVolumesCall.Returns.Volumes, v.PageToVolumesCall.Returns.Error
}

func (v *VolumesAPI) GetVolumesPager() pagination.Pager {
	v.GetVolumesPagerCall.CallCount++

	return v.GetVolumesPagerCall.Returns.Pager
}

func (v *VolumesAPI) DeleteVolume(volumeID string) error {
	v.DeleteVolumeCall.CallCount++
	v.DeleteVolumeCall.Receives.VolumeID = volumeID

	v.DeleteVolumeCall.ReceivesForCall = append(v.DeleteVolumeCall.ReceivesForCall, v.DeleteVolumeCall.Receives)

	return v.DeleteVolumeCall.Returns.Error
}
