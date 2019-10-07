package fakes

import (
	"sync"

	gcpcompute "google.golang.org/api/compute/v1"
)

type HttpsHealthChecksClient struct {
	DeleteHttpsHealthCheckCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			HttpsHealthCheck string
		}
		Returns struct {
			Error error
		}
		Stub func(string) error
	}
	ListHttpsHealthChecksCall struct {
		sync.Mutex
		CallCount int
		Returns   struct {
			HttpsHealthCheckSlice []*gcpcompute.HttpsHealthCheck
			Error                 error
		}
		Stub func() ([]*gcpcompute.HttpsHealthCheck, error)
	}
}

func (f *HttpsHealthChecksClient) DeleteHttpsHealthCheck(param1 string) error {
	f.DeleteHttpsHealthCheckCall.Lock()
	defer f.DeleteHttpsHealthCheckCall.Unlock()
	f.DeleteHttpsHealthCheckCall.CallCount++
	f.DeleteHttpsHealthCheckCall.Receives.HttpsHealthCheck = param1
	if f.DeleteHttpsHealthCheckCall.Stub != nil {
		return f.DeleteHttpsHealthCheckCall.Stub(param1)
	}
	return f.DeleteHttpsHealthCheckCall.Returns.Error
}
func (f *HttpsHealthChecksClient) ListHttpsHealthChecks() ([]*gcpcompute.HttpsHealthCheck, error) {
	f.ListHttpsHealthChecksCall.Lock()
	defer f.ListHttpsHealthChecksCall.Unlock()
	f.ListHttpsHealthChecksCall.CallCount++
	if f.ListHttpsHealthChecksCall.Stub != nil {
		return f.ListHttpsHealthChecksCall.Stub()
	}
	return f.ListHttpsHealthChecksCall.Returns.HttpsHealthCheckSlice, f.ListHttpsHealthChecksCall.Returns.Error
}
