package fakes

import gcpcompute "google.golang.org/api/compute/v1"

type BackendServicesClient struct {
	ListBackendServicesCall struct {
		CallCount int
		Receives  struct {
			Filter string
		}
		Returns struct {
			Output *gcpcompute.BackendServiceList
			Error  error
		}
	}

	DeleteBackendServiceCall struct {
		CallCount int
		Receives  struct {
			BackendService string
		}
		Returns struct {
			Error error
		}
	}
}

func (n *BackendServicesClient) ListBackendServices(filter string) (*gcpcompute.BackendServiceList, error) {
	n.ListBackendServicesCall.CallCount++
	n.ListBackendServicesCall.Receives.Filter = filter

	return n.ListBackendServicesCall.Returns.Output, n.ListBackendServicesCall.Returns.Error
}

func (n *BackendServicesClient) DeleteBackendService(backendService string) error {
	n.DeleteBackendServiceCall.CallCount++
	n.DeleteBackendServiceCall.Receives.BackendService = backendService

	return n.DeleteBackendServiceCall.Returns.Error
}
