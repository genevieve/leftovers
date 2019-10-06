package fakes

import (
	"sync"

	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
)

type ComputeClient struct {
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
	ListCall struct {
		sync.Mutex
		CallCount int
		Returns   struct {
			ServerSlice []servers.Server
			Error       error
		}
		Stub func() ([]servers.Server, error)
	}
}

func (f *ComputeClient) Delete(param1 string) error {
	f.DeleteCall.Lock()
	defer f.DeleteCall.Unlock()
	f.DeleteCall.CallCount++
	f.DeleteCall.Receives.InstanceID = param1
	if f.DeleteCall.Stub != nil {
		return f.DeleteCall.Stub(param1)
	}
	return f.DeleteCall.Returns.Error
}
func (f *ComputeClient) List() ([]servers.Server, error) {
	f.ListCall.Lock()
	defer f.ListCall.Unlock()
	f.ListCall.CallCount++
	if f.ListCall.Stub != nil {
		return f.ListCall.Stub()
	}
	return f.ListCall.Returns.ServerSlice, f.ListCall.Returns.Error
}
