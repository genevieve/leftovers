package openstack

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumes"
	"github.com/gophercloud/gophercloud/pagination"
)

type VolumesBlockStorageClient struct {
	serviceClient *gophercloud.ServiceClient
	volumesAPI    volumesAPI
}

type volumesAPI interface {
	PagerToPage(pager pagination.Pager) (pagination.Page, error)
	PageToVolumes(pagination.Page) ([]volumes.Volume, error)
	GetVolumesPager(*gophercloud.ServiceClient, volumes.ListOpts) pagination.Pager
	DeleteVolume(*gophercloud.ServiceClient, string, volumes.DeleteOpts) error
}

func NewVolumesBlockStorageClient(serviceClient *gophercloud.ServiceClient, volumesAPI volumesAPI) VolumesBlockStorageClient {
	return VolumesBlockStorageClient{serviceClient, volumesAPI}
}

func (vs VolumesBlockStorageClient) List() ([]volumes.Volume, error) {
	pager := vs.volumesAPI.GetVolumesPager(vs.serviceClient, volumes.ListOpts{})

	page, err := vs.volumesAPI.PagerToPage(pager)
	if err != nil {
		return nil, err
	}
	result, err := vs.volumesAPI.PageToVolumes(page)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (vs VolumesBlockStorageClient) Delete(volumeID string) error {
	return vs.volumesAPI.DeleteVolume(vs.serviceClient, volumeID, volumes.DeleteOpts{})
}

func (vs VolumesBlockStorageClient) GetVolumesDeleter() VolumesDeleter {
	return vs
}

func (vs VolumesBlockStorageClient) GetVolumesLister() VolumesLister {
	return vs
}
