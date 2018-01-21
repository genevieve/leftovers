package fakes

import gcpcompute "google.golang.org/api/compute/v1"

type HttpHealthChecksClient struct {
	ListHttpHealthChecksCall struct {
		CallCount int
		Receives  struct {
			Filter string
		}
		Returns struct {
			Output *gcpcompute.HttpHealthCheckList
			Error  error
		}
	}

	DeleteHttpHealthCheckCall struct {
		CallCount int
		Receives  struct {
			HttpHealthCheck string
		}
		Returns struct {
			Error error
		}
	}
}

func (n *HttpHealthChecksClient) ListHttpHealthChecks(filter string) (*gcpcompute.HttpHealthCheckList, error) {
	n.ListHttpHealthChecksCall.CallCount++
	n.ListHttpHealthChecksCall.Receives.Filter = filter

	return n.ListHttpHealthChecksCall.Returns.Output, n.ListHttpHealthChecksCall.Returns.Error
}

func (n *HttpHealthChecksClient) DeleteHttpHealthCheck(httpHealthCheck string) error {
	n.DeleteHttpHealthCheckCall.CallCount++
	n.DeleteHttpHealthCheckCall.Receives.HttpHealthCheck = httpHealthCheck

	return n.DeleteHttpHealthCheckCall.Returns.Error
}
