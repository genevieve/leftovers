package openstack

import (
	"errors"
	"time"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumes"
	"github.com/gophercloud/gophercloud/pagination"
)

type VolumesAPI struct {
	serviceClient *gophercloud.ServiceClient
	waitTime      time.Duration
	maxRetries    int
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
	var retries int

	for {
		volume, err := getVolume(api.serviceClient, id)
		if err != nil {
			return err
		}

		status := volume.Status
		if status == "available" || status == "error" || status == "error_restoring" || status == "error_extending" || status == "error_managing" {
			break
		}

		if retries == api.maxRetries {
			return errors.New("volume failed to reach desired state")
		}

		retries++
		time.Sleep(api.waitTime)
	}

	return volumes.Delete(api.serviceClient, id, volumes.DeleteOpts{}).ExtractErr()
}

func getVolume(serviceClient *gophercloud.ServiceClient, id string) (*volumes.Volume, error) {
	volume, err := volumes.Get(serviceClient, id).Extract()
	if err != nil {
		return &volumes.Volume{}, err
	}

	return volume, nil
}
