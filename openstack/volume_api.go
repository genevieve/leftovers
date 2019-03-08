package openstack

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumes"
	"github.com/gophercloud/gophercloud/pagination"
)

type VolumesAPI struct {
	serviceClient *gophercloud.ServiceClient
}

func (api VolumesAPI) GetVolumesPager() pagination.Pager {
	return volumes.List(api.serviceClient, volumes.ListOpts{})
}

func (api VolumesAPI) PagerToPage(pager pagination.Pager) (pagination.Page, error) {
	return pager.AllPages()
}

func (api VolumesAPI) PageToVolumes(page pagination.Page) ([]volumes.Volume, error) {
	return volumes.ExtractVolumes(page)
}

func (api VolumesAPI) DeleteVolume(id string) error {
	return volumes.Delete(api.serviceClient, id, volumes.DeleteOpts{}).ExtractErr()
}
