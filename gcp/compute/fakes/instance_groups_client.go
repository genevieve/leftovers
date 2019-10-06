package fakes

import (
	"sync"

	gcpcompute "google.golang.org/api/compute/v1"
)

type InstanceGroupsClient struct {
	DeleteInstanceGroupCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Zone          string
			InstanceGroup string
		}
		Returns struct {
			Error error
		}
		Stub func(string, string) error
	}
	ListInstanceGroupsCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Zone string
		}
		Returns struct {
			InstanceGroupSlice []*gcpcompute.InstanceGroup
			Error              error
		}
		Stub func(string) ([]*gcpcompute.InstanceGroup, error)
	}
}

func (f *InstanceGroupsClient) DeleteInstanceGroup(param1 string, param2 string) error {
	f.DeleteInstanceGroupCall.Lock()
	defer f.DeleteInstanceGroupCall.Unlock()
	f.DeleteInstanceGroupCall.CallCount++
	f.DeleteInstanceGroupCall.Receives.Zone = param1
	f.DeleteInstanceGroupCall.Receives.InstanceGroup = param2
	if f.DeleteInstanceGroupCall.Stub != nil {
		return f.DeleteInstanceGroupCall.Stub(param1, param2)
	}
	return f.DeleteInstanceGroupCall.Returns.Error
}
func (f *InstanceGroupsClient) ListInstanceGroups(param1 string) ([]*gcpcompute.InstanceGroup, error) {
	f.ListInstanceGroupsCall.Lock()
	defer f.ListInstanceGroupsCall.Unlock()
	f.ListInstanceGroupsCall.CallCount++
	f.ListInstanceGroupsCall.Receives.Zone = param1
	if f.ListInstanceGroupsCall.Stub != nil {
		return f.ListInstanceGroupsCall.Stub(param1)
	}
	return f.ListInstanceGroupsCall.Returns.InstanceGroupSlice, f.ListInstanceGroupsCall.Returns.Error
}
