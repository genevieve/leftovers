package fakes

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"github.com/gophercloud/gophercloud/pagination"
)

type ComputeInstanceAPI struct {
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
	PageToServersCall struct {
		CallCount int
		Receives  struct {
			Page pagination.Page
		}
		Returns struct {
			Servers []servers.Server
			Error   error
		}
	}
	GetComputeInstancePagerCall struct {
		CallCount int
		Receives  struct {
			Service *gophercloud.ServiceClient
		}
		Returns struct {
			ComputeInstancePager pagination.Pager
		}
	}
	DeleteCall struct {
		CallCount int
		Receives  struct {
			ServiceClient *gophercloud.ServiceClient
			InstanceID    string
		}
		Returns struct {
			Error error
		}
	}
}

func (api *ComputeInstanceAPI) GetComputeInstancePager(service *gophercloud.ServiceClient) pagination.Pager {
	api.GetComputeInstancePagerCall.CallCount++
	api.GetComputeInstancePagerCall.Receives.Service = service

	return api.GetComputeInstancePagerCall.Returns.ComputeInstancePager
}

func (api *ComputeInstanceAPI) PagerToPage(pager pagination.Pager) (pagination.Page, error) {
	api.PagerToPageCall.CallCount++
	api.PagerToPageCall.Receives.Pager = pager

	return api.PagerToPageCall.Returns.Page, api.PagerToPageCall.Returns.Error
}

func (api *ComputeInstanceAPI) PageToServers(page pagination.Page) ([]servers.Server, error) {
	api.PageToServersCall.CallCount++
	api.PageToServersCall.Receives.Page = page

	return api.PageToServersCall.Returns.Servers, api.PageToServersCall.Returns.Error
}

func (api *ComputeInstanceAPI) Delete(serviceClient *gophercloud.ServiceClient, instanceID string) error {
	api.DeleteCall.CallCount++
	api.DeleteCall.Receives.InstanceID = instanceID
	api.DeleteCall.Receives.ServiceClient = serviceClient

	return api.DeleteCall.Returns.Error
}
