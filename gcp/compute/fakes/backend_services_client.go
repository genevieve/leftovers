package fakes

import (
	"sync"

	gcpcompute "google.golang.org/api/compute/v1"
)

type BackendServicesClient struct {
	DeleteBackendServiceCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			BackendService string
		}
		Returns struct {
			Error error
		}
		Stub func(string) error
	}
	ListBackendServicesCall struct {
		sync.Mutex
		CallCount int
		Returns   struct {
			BackendServiceSlice []*gcpcompute.BackendService
			Error               error
		}
		Stub func() ([]*gcpcompute.BackendService, error)
	}
}

func (f *BackendServicesClient) DeleteBackendService(param1 string) error {
	f.DeleteBackendServiceCall.Lock()
	defer f.DeleteBackendServiceCall.Unlock()
	f.DeleteBackendServiceCall.CallCount++
	f.DeleteBackendServiceCall.Receives.BackendService = param1
	if f.DeleteBackendServiceCall.Stub != nil {
		return f.DeleteBackendServiceCall.Stub(param1)
	}
	return f.DeleteBackendServiceCall.Returns.Error
}
func (f *BackendServicesClient) ListBackendServices() ([]*gcpcompute.BackendService, error) {
	f.ListBackendServicesCall.Lock()
	defer f.ListBackendServicesCall.Unlock()
	f.ListBackendServicesCall.CallCount++
	if f.ListBackendServicesCall.Stub != nil {
		return f.ListBackendServicesCall.Stub()
	}
	return f.ListBackendServicesCall.Returns.BackendServiceSlice, f.ListBackendServicesCall.Returns.Error
}
