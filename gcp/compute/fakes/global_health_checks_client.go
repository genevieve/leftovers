package fakes

import gcpcompute "google.golang.org/api/compute/v1"

type GlobalHealthChecksClient struct {
	ListGlobalHealthChecksCall struct {
		CallCount int
		Returns   struct {
			Output []*gcpcompute.HealthCheck
			Error  error
		}
	}

	DeleteGlobalHealthCheckCall struct {
		CallCount int
		Receives  struct {
			GlobalHealthCheck string
		}
		Returns struct {
			Error error
		}
	}
}

func (n *GlobalHealthChecksClient) ListGlobalHealthChecks() ([]*gcpcompute.HealthCheck, error) {
	n.ListGlobalHealthChecksCall.CallCount++

	return n.ListGlobalHealthChecksCall.Returns.Output, n.ListGlobalHealthChecksCall.Returns.Error
}

func (n *GlobalHealthChecksClient) DeleteGlobalHealthCheck(globalHealthCheck string) error {
	n.DeleteGlobalHealthCheckCall.CallCount++
	n.DeleteGlobalHealthCheckCall.Receives.GlobalHealthCheck = globalHealthCheck

	return n.DeleteGlobalHealthCheckCall.Returns.Error
}
