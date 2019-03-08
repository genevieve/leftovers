package openstack

import (
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumes"
	"github.com/gophercloud/gophercloud/pagination"
)

type volumesAPI interface {
	PagerToPage(pager pagination.Pager) (pagination.Page, error)
	PageToVolumes(pagination.Page) ([]volumes.Volume, error)
	GetVolumesPager() pagination.Pager
	DeleteVolume(volumeID string) error
}

type VolumesBlockStorageClient struct {
	volumesAPI volumesAPI
}

func NewVolumesBlockStorageClient(volumesAPI volumesAPI) VolumesBlockStorageClient {
	return VolumesBlockStorageClient{
		volumesAPI: volumesAPI,
	}
}

func (vs VolumesBlockStorageClient) List() ([]volumes.Volume, error) {
	pager := vs.volumesAPI.GetVolumesPager()

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
	return vs.volumesAPI.DeleteVolume(volumeID)
}
