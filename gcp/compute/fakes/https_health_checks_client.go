package fakes

import gcpcompute "google.golang.org/api/compute/v1"

type HttpsHealthChecksClient struct {
	ListHttpsHealthChecksCall struct {
		CallCount int
		Receives  struct {
			Filter string
		}
		Returns struct {
			Output *gcpcompute.HttpsHealthCheckList
			Error  error
		}
	}

	DeleteHttpsHealthCheckCall struct {
		CallCount int
		Receives  struct {
			HttpsHealthCheck string
		}
		Returns struct {
			Error error
		}
	}
}

func (n *HttpsHealthChecksClient) ListHttpsHealthChecks(filter string) (*gcpcompute.HttpsHealthCheckList, error) {
	n.ListHttpsHealthChecksCall.CallCount++
	n.ListHttpsHealthChecksCall.Receives.Filter = filter

	return n.ListHttpsHealthChecksCall.Returns.Output, n.ListHttpsHealthChecksCall.Returns.Error
}

func (n *HttpsHealthChecksClient) DeleteHttpsHealthCheck(httpsHealthCheck string) error {
	n.DeleteHttpsHealthCheckCall.CallCount++
	n.DeleteHttpsHealthCheckCall.Receives.HttpsHealthCheck = httpsHealthCheck

	return n.DeleteHttpsHealthCheckCall.Returns.Error
}
