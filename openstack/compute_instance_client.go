package openstack

import (
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"github.com/gophercloud/gophercloud/pagination"
)

type ComputeInstanceAPI interface {
	GetComputeInstancePager() pagination.Pager
	PagerToPage(pagination.Pager) (pagination.Page, error)
	PageToServers(pagination.Page) ([]servers.Server, error)
	Delete(instanceID string) error
}

type ComputeInstanceClient struct {
	api ComputeInstanceAPI
}

func NewComputeInstanceClient(api ComputeInstanceAPI) ComputeInstanceClient {
	return ComputeInstanceClient{
		api: api,
	}
}

func (client ComputeInstanceClient) List() ([]servers.Server, error) {
	pager := client.api.GetComputeInstancePager()
	page, err := client.api.PagerToPage(pager)
	if err != nil {
		return nil, err
	}
	servers, err := client.api.PageToServers(page)
	if err != nil {
		return nil, err
	}
	return servers, nil
}

func (client ComputeInstanceClient) Delete(instanceID string) error {
	return client.api.Delete(instanceID)
}
