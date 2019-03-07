package openstack

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"github.com/gophercloud/gophercloud/pagination"
)

type ComputeInstanceAPI interface {
	GetComputeInstancePager(*gophercloud.ServiceClient) pagination.Pager
	PagerToPage(pagination.Pager) (pagination.Page, error)
	PageToServers(pagination.Page) ([]servers.Server, error)
	Delete(serviceClient *gophercloud.ServiceClient, instanceID string) error
}

type ComputeInstanceClient struct {
	api     ComputeInstanceAPI
	service *gophercloud.ServiceClient
}

func NewComputeInstanceClient(service *gophercloud.ServiceClient, api ComputeInstanceAPI) ComputeInstanceClient {
	return ComputeInstanceClient{api, service}
}

func (client ComputeInstanceClient) List() ([]servers.Server, error) {
	pager := client.api.GetComputeInstancePager(client.service)
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
	return client.api.Delete(client.service, instanceID)
}

func (client ComputeInstanceClient) GetComputeInstanceDeleter() ComputeInstanceDeleter {
	return client
}

func (client ComputeInstanceClient) GetComputeInstanceLister() ComputeInstanceLister {
	return client
}
