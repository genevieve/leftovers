package openstack

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"github.com/gophercloud/gophercloud/pagination"
)

type ComputeAPI struct {
	serviceClient *gophercloud.ServiceClient
}

func (api ComputeAPI) GetComputeInstancePager() pagination.Pager {
	return servers.List(api.serviceClient, servers.ListOpts{})
}

func (api ComputeAPI) PagerToPage(pager pagination.Pager) (pagination.Page, error) {
	return pager.AllPages()
}

func (api ComputeAPI) PageToServers(page pagination.Page) ([]servers.Server, error) {
	return servers.ExtractServers(page)
}

func (api ComputeAPI) Delete(instanceID string) error {
	return servers.Delete(api.serviceClient, instanceID).ExtractErr()
}
