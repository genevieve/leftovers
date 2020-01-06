package fakes

import (
	"sync"

	gcpcompute "google.golang.org/api/compute/v1"
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
	GetNetworkNameCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Url string
		}
		Returns struct {
			Name string
		}
		Stub func(string) string
	}
	ListInstancesCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Zone string
		}
		Returns struct {
			InstanceSlice []*gcpcompute.Instance
			Error         error
		}
		Stub func(string) ([]*gcpcompute.Instance, error)
	}
	SetDiskAutoDeleteCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Zone     string
			Instance string
			Disk     string
		}
		Returns struct {
			Error error
		}
		Stub func(string, string, string) error
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
func (f *InstancesClient) GetNetworkName(param1 string) string {
	f.GetNetworkNameCall.Lock()
	defer f.GetNetworkNameCall.Unlock()
	f.GetNetworkNameCall.CallCount++
	f.GetNetworkNameCall.Receives.Url = param1
	if f.GetNetworkNameCall.Stub != nil {
		return f.GetNetworkNameCall.Stub(param1)
	}
	return f.GetNetworkNameCall.Returns.Name
}
func (f *InstancesClient) ListInstances(param1 string) ([]*gcpcompute.Instance, error) {
	f.ListInstancesCall.Lock()
	defer f.ListInstancesCall.Unlock()
	f.ListInstancesCall.CallCount++
	f.ListInstancesCall.Receives.Zone = param1
	if f.ListInstancesCall.Stub != nil {
		return f.ListInstancesCall.Stub(param1)
	}
	return f.ListInstancesCall.Returns.InstanceSlice, f.ListInstancesCall.Returns.Error
}
func (f *InstancesClient) SetDiskAutoDelete(param1 string, param2 string, param3 string) error {
	f.SetDiskAutoDeleteCall.Lock()
	defer f.SetDiskAutoDeleteCall.Unlock()
	f.SetDiskAutoDeleteCall.CallCount++
	f.SetDiskAutoDeleteCall.Receives.Zone = param1
	f.SetDiskAutoDeleteCall.Receives.Instance = param2
	f.SetDiskAutoDeleteCall.Receives.Disk = param3
	if f.SetDiskAutoDeleteCall.Stub != nil {
		return f.SetDiskAutoDeleteCall.Stub(param1, param2, param3)
	}
	return f.SetDiskAutoDeleteCall.Returns.Error
}
