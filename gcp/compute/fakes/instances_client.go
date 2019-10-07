package fakes

import (
	"sync"

	compute "google.golang.org/api/compute/v1"
)

type InstancesClient struct {
	DeleteInstanceCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Zone     string
			Instance string
		}
		Returns struct {
			Error error
		}
		Stub func(string, string) error
	}
	ListInstancesCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Zone string
		}
		Returns struct {
			InstanceSlice []*compute.Instance
			Error         error
		}
		Stub func(string) ([]*compute.Instance, error)
	}
}

func (f *InstancesClient) DeleteInstance(param1 string, param2 string) error {
	f.DeleteInstanceCall.Lock()
	defer f.DeleteInstanceCall.Unlock()
	f.DeleteInstanceCall.CallCount++
	f.DeleteInstanceCall.Receives.Zone = param1
	f.DeleteInstanceCall.Receives.Instance = param2
	if f.DeleteInstanceCall.Stub != nil {
		return f.DeleteInstanceCall.Stub(param1, param2)
	}
	return f.DeleteInstanceCall.Returns.Error
}
func (f *InstancesClient) ListInstances(param1 string) ([]*compute.Instance, error) {
	f.ListInstancesCall.Lock()
	defer f.ListInstancesCall.Unlock()
	f.ListInstancesCall.CallCount++
	f.ListInstancesCall.Receives.Zone = param1
	if f.ListInstancesCall.Stub != nil {
		return f.ListInstancesCall.Stub(param1)
	}
	return f.ListInstancesCall.Returns.InstanceSlice, f.ListInstancesCall.Returns.Error
}
