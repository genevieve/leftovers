package fakes

import (
	"sync"

	gcpsql "google.golang.org/api/sqladmin/v1beta4"
)

type InstancesClient struct {
	DeleteInstanceCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			User string
		}
		Returns struct {
			Error error
		}
		Stub func(string) error
	}
	ListInstancesCall struct {
		sync.Mutex
		CallCount int
		Returns   struct {
			InstancesListResponse *gcpsql.InstancesListResponse
			Error                 error
		}
		Stub func() (*gcpsql.InstancesListResponse, error)
	}
}

func (f *InstancesClient) DeleteInstance(param1 string) error {
	f.DeleteInstanceCall.Lock()
	defer f.DeleteInstanceCall.Unlock()
	f.DeleteInstanceCall.CallCount++
	f.DeleteInstanceCall.Receives.User = param1
	if f.DeleteInstanceCall.Stub != nil {
		return f.DeleteInstanceCall.Stub(param1)
	}
	return f.DeleteInstanceCall.Returns.Error
}
func (f *InstancesClient) ListInstances() (*gcpsql.InstancesListResponse, error) {
	f.ListInstancesCall.Lock()
	defer f.ListInstancesCall.Unlock()
	f.ListInstancesCall.CallCount++
	if f.ListInstancesCall.Stub != nil {
		return f.ListInstancesCall.Stub()
	}
	return f.ListInstancesCall.Returns.InstancesListResponse, f.ListInstancesCall.Returns.Error
}
