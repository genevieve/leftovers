package fakes

import gcpcompute "google.golang.org/api/compute/v1"

type HttpHealthChecksClient struct {
	ListHttpHealthChecksCall struct {
		CallCount int
		Returns   struct {
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
			Output *gcpcompute.Operation
			Error  error
		}
	}
}

func (n *HttpHealthChecksClient) ListHttpHealthChecks() (*gcpcompute.HttpHealthCheckList, error) {
	n.ListHttpHealthChecksCall.CallCount++

	return n.ListHttpHealthChecksCall.Returns.Output, n.ListHttpHealthChecksCall.Returns.Error
}

func (n *HttpHealthChecksClient) DeleteHttpHealthCheck(disk string) (*gcpcompute.Operation, error) {
	n.DeleteHttpHealthCheckCall.CallCount++
	n.DeleteHttpHealthCheckCall.Receives.HttpHealthCheck = disk

	return n.DeleteHttpHealthCheckCall.Returns.Output, n.DeleteHttpHealthCheckCall.Returns.Error
}
