package openstack

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumes"
	"github.com/gophercloud/gophercloud/pagination"
)

type VolumesAPI struct{}

func (api VolumesAPI) PagerToPage(pager pagination.Pager) (pagination.Page, error) {
	return pager.AllPages()
}

func (api VolumesAPI) PageToVolumes(page pagination.Page) ([]volumes.Volume, error) {
	return volumes.ExtractVolumes(page)
}

func (api VolumesAPI) GetVolumesPager(serviceClient *gophercloud.ServiceClient, opts volumes.ListOpts) pagination.Pager {
	return volumes.List(serviceClient, opts)
}

func (api VolumesAPI) DeleteVolume(serviceClient *gophercloud.ServiceClient, id string, opts volumes.DeleteOpts) error {
	return volumes.Delete(serviceClient, id, opts).ExtractErr()
}
