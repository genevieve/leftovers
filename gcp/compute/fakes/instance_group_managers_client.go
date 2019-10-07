package fakes

import (
	"sync"

	gcpcompute "google.golang.org/api/compute/v1"
)

type InstanceGroupManagersClient struct {
	DeleteInstanceGroupManagerCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Zone                 string
			InstanceGroupManager string
		}
		Returns struct {
			Error error
		}
		Stub func(string, string) error
	}
	ListInstanceGroupManagersCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Zone string
		}
		Returns struct {
			InstanceGroupManagerSlice []*gcpcompute.InstanceGroupManager
			Error                     error
		}
		Stub func(string) ([]*gcpcompute.InstanceGroupManager, error)
	}
}

func (f *InstanceGroupManagersClient) DeleteInstanceGroupManager(param1 string, param2 string) error {
	f.DeleteInstanceGroupManagerCall.Lock()
	defer f.DeleteInstanceGroupManagerCall.Unlock()
	f.DeleteInstanceGroupManagerCall.CallCount++
	f.DeleteInstanceGroupManagerCall.Receives.Zone = param1
	f.DeleteInstanceGroupManagerCall.Receives.InstanceGroupManager = param2
	if f.DeleteInstanceGroupManagerCall.Stub != nil {
		return f.DeleteInstanceGroupManagerCall.Stub(param1, param2)
	}
	return f.DeleteInstanceGroupManagerCall.Returns.Error
}
func (f *InstanceGroupManagersClient) ListInstanceGroupManagers(param1 string) ([]*gcpcompute.InstanceGroupManager, error) {
	f.ListInstanceGroupManagersCall.Lock()
	defer f.ListInstanceGroupManagersCall.Unlock()
	f.ListInstanceGroupManagersCall.CallCount++
	f.ListInstanceGroupManagersCall.Receives.Zone = param1
	if f.ListInstanceGroupManagersCall.Stub != nil {
		return f.ListInstanceGroupManagersCall.Stub(param1)
	}
	return f.ListInstanceGroupManagersCall.Returns.InstanceGroupManagerSlice, f.ListInstanceGroupManagersCall.Returns.Error
}
