package fakes

import (
	"github.com/gophercloud/gophercloud"
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
		Receives  struct {
			Client  *gophercloud.ServiceClient
			Options volumes.ListOpts
		}
		Returns struct {
			Pager pagination.Pager
		}
	}

	DeleteVolumeCall struct {
		CallCount int
		Receives  struct {
			Client   *gophercloud.ServiceClient
			VolumeID string
			Options  volumes.DeleteOpts
		}
		ReceivesForCall []struct {
			Client   *gophercloud.ServiceClient
			VolumeID string
			Options  volumes.DeleteOpts
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

func (v *VolumesAPI) GetVolumesPager(client *gophercloud.ServiceClient, options volumes.ListOpts) pagination.Pager {
	v.GetVolumesPagerCall.CallCount++
	v.GetVolumesPagerCall.Receives.Client = client
	v.GetVolumesPagerCall.Receives.Options = options

	return v.GetVolumesPagerCall.Returns.Pager
}

func (v *VolumesAPI) DeleteVolume(client *gophercloud.ServiceClient, volumeID string, options volumes.DeleteOpts) error {
	v.DeleteVolumeCall.CallCount++
	v.DeleteVolumeCall.Receives.Client = client
	v.DeleteVolumeCall.Receives.VolumeID = volumeID
	v.DeleteVolumeCall.Receives.Options = options

	v.DeleteVolumeCall.ReceivesForCall = append(v.DeleteVolumeCall.ReceivesForCall, v.DeleteVolumeCall.Receives)

	return v.DeleteVolumeCall.Returns.Error
}
