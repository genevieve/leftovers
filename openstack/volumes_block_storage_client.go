package openstack

import (
	"fmt"

	"github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumes"
	"github.com/gophercloud/gophercloud/pagination"
)

//go:generate faux --interface volumesAPI --output fakes/volumes_api.go
type volumesAPI interface {
	PagerToPage(pager pagination.Pager) (pagination.Page, error)
	PageToVolumes(pagination.Page) ([]volumes.Volume, error)
	GetVolumesPager() pagination.Pager
	DeleteVolume(id string) error
}

type VolumesBlockStorageClient struct {
	api volumesAPI
}

func NewVolumesBlockStorageClient(api volumesAPI) VolumesBlockStorageClient {
	return VolumesBlockStorageClient{
		api: api,
	}
}

func (v VolumesBlockStorageClient) List() ([]volumes.Volume, error) {
	pager := v.api.GetVolumesPager()

	page, err := v.api.PagerToPage(pager)
	if err != nil {
		return nil, fmt.Errorf("pager to page: %s", err)
	}

	result, err := v.api.PageToVolumes(page)
	if err != nil {
		return nil, fmt.Errorf("page to volumes: %s", err)
	}

	return result, nil
}

func (v VolumesBlockStorageClient) Delete(volumeID string) error {
	return v.api.DeleteVolume(volumeID)
}
