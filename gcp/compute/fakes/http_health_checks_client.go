package fakes

import (
	"sync"

	gcpcompute "google.golang.org/api/compute/v1"
)

type HttpHealthChecksClient struct {
	DeleteHttpHealthCheckCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			HttpHealthCheck string
		}
		Returns struct {
			Error error
		}
		Stub func(string) error
	}
	ListHttpHealthChecksCall struct {
		sync.Mutex
		CallCount int
		Returns   struct {
			HttpHealthCheckSlice []*gcpcompute.HttpHealthCheck
			Error                error
		}
		Stub func() ([]*gcpcompute.HttpHealthCheck, error)
	}
}

func (f *HttpHealthChecksClient) DeleteHttpHealthCheck(param1 string) error {
	f.DeleteHttpHealthCheckCall.Lock()
	defer f.DeleteHttpHealthCheckCall.Unlock()
	f.DeleteHttpHealthCheckCall.CallCount++
	f.DeleteHttpHealthCheckCall.Receives.HttpHealthCheck = param1
	if f.DeleteHttpHealthCheckCall.Stub != nil {
		return f.DeleteHttpHealthCheckCall.Stub(param1)
	}
	return f.DeleteHttpHealthCheckCall.Returns.Error
}
func (f *HttpHealthChecksClient) ListHttpHealthChecks() ([]*gcpcompute.HttpHealthCheck, error) {
	f.ListHttpHealthChecksCall.Lock()
	defer f.ListHttpHealthChecksCall.Unlock()
	f.ListHttpHealthChecksCall.CallCount++
	if f.ListHttpHealthChecksCall.Stub != nil {
		return f.ListHttpHealthChecksCall.Stub()
	}
	return f.ListHttpHealthChecksCall.Returns.HttpHealthCheckSlice, f.ListHttpHealthChecksCall.Returns.Error
}
