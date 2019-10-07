package fakes

import (
	"sync"

	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"github.com/gophercloud/gophercloud/pagination"
)

type ComputeInstanceAPI struct {
	DeleteCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			InstanceID string
		}
		Returns struct {
			Error error
		}
		Stub func(string) error
	}
	GetComputeInstancePagerCall struct {
		sync.Mutex
		CallCount int
		Returns   struct {
			Pager pagination.Pager
		}
		Stub func() pagination.Pager
	}
	PageToServersCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Page pagination.Page
		}
		Returns struct {
			ServerSlice []servers.Server
			Error       error
		}
		Stub func(pagination.Page) ([]servers.Server, error)
	}
	PagerToPageCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Pager pagination.Pager
		}
		Returns struct {
			Page  pagination.Page
			Error error
		}
		Stub func(pagination.Pager) (pagination.Page, error)
	}
}

func (f *ComputeInstanceAPI) Delete(param1 string) error {
	f.DeleteCall.Lock()
	defer f.DeleteCall.Unlock()
	f.DeleteCall.CallCount++
	f.DeleteCall.Receives.InstanceID = param1
	if f.DeleteCall.Stub != nil {
		return f.DeleteCall.Stub(param1)
	}
	return f.DeleteCall.Returns.Error
}
func (f *ComputeInstanceAPI) GetComputeInstancePager() pagination.Pager {
	f.GetComputeInstancePagerCall.Lock()
	defer f.GetComputeInstancePagerCall.Unlock()
	f.GetComputeInstancePagerCall.CallCount++
	if f.GetComputeInstancePagerCall.Stub != nil {
		return f.GetComputeInstancePagerCall.Stub()
	}
	return f.GetComputeInstancePagerCall.Returns.Pager
}
func (f *ComputeInstanceAPI) PageToServers(param1 pagination.Page) ([]servers.Server, error) {
	f.PageToServersCall.Lock()
	defer f.PageToServersCall.Unlock()
	f.PageToServersCall.CallCount++
	f.PageToServersCall.Receives.Page = param1
	if f.PageToServersCall.Stub != nil {
		return f.PageToServersCall.Stub(param1)
	}
	return f.PageToServersCall.Returns.ServerSlice, f.PageToServersCall.Returns.Error
}
func (f *ComputeInstanceAPI) PagerToPage(param1 pagination.Pager) (pagination.Page, error) {
	f.PagerToPageCall.Lock()
	defer f.PagerToPageCall.Unlock()
	f.PagerToPageCall.CallCount++
	f.PagerToPageCall.Receives.Pager = param1
	if f.PagerToPageCall.Stub != nil {
		return f.PagerToPageCall.Stub(param1)
	}
	return f.PagerToPageCall.Returns.Page, f.PagerToPageCall.Returns.Error
}
