package fakes

import (
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
		Returns   struct {
			ComputeInstancePager pagination.Pager
		}
	}
	DeleteCall struct {
		CallCount int
		Receives  struct {
			InstanceID string
		}
		Returns struct {
			Error error
		}
	}
}

func (api *ComputeInstanceAPI) GetComputeInstancePager() pagination.Pager {
	api.GetComputeInstancePagerCall.CallCount++

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

func (api *ComputeInstanceAPI) Delete(instanceID string) error {
	api.DeleteCall.CallCount++
	api.DeleteCall.Receives.InstanceID = instanceID

	return api.DeleteCall.Returns.Error
}
