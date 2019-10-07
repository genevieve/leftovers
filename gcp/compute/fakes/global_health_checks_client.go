package fakes

import (
	"sync"

	gcpcompute "google.golang.org/api/compute/v1"
)

type GlobalHealthChecksClient struct {
	DeleteGlobalHealthCheckCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			GlobalHealthCheck string
		}
		Returns struct {
			Error error
		}
		Stub func(string) error
	}
	ListGlobalHealthChecksCall struct {
		sync.Mutex
		CallCount int
		Returns   struct {
			HealthCheckSlice []*gcpcompute.HealthCheck
			Error            error
		}
		Stub func() ([]*gcpcompute.HealthCheck, error)
	}
}

func (f *GlobalHealthChecksClient) DeleteGlobalHealthCheck(param1 string) error {
	f.DeleteGlobalHealthCheckCall.Lock()
	defer f.DeleteGlobalHealthCheckCall.Unlock()
	f.DeleteGlobalHealthCheckCall.CallCount++
	f.DeleteGlobalHealthCheckCall.Receives.GlobalHealthCheck = param1
	if f.DeleteGlobalHealthCheckCall.Stub != nil {
		return f.DeleteGlobalHealthCheckCall.Stub(param1)
	}
	return f.DeleteGlobalHealthCheckCall.Returns.Error
}
func (f *GlobalHealthChecksClient) ListGlobalHealthChecks() ([]*gcpcompute.HealthCheck, error) {
	f.ListGlobalHealthChecksCall.Lock()
	defer f.ListGlobalHealthChecksCall.Unlock()
	f.ListGlobalHealthChecksCall.CallCount++
	if f.ListGlobalHealthChecksCall.Stub != nil {
		return f.ListGlobalHealthChecksCall.Stub()
	}
	return f.ListGlobalHealthChecksCall.Returns.HealthCheckSlice, f.ListGlobalHealthChecksCall.Returns.Error
}
