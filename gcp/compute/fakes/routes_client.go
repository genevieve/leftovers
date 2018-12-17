package fakes

import gcpcompute "google.golang.org/api/compute/v1"

type RoutesClient struct {
	ListRoutesCall struct {
		CallCount int
		Returns struct {
			Output []*gcpcompute.Route
			Error error
		}
	}

	DeleteRouteCall struct {
		CallCount int
		Receives struct {
			Route string
		}
		Returns struct {
			Error error
		}
	}
}

func (n *RoutesClient) ListRoutes() ([]*gcpcompute.Route, error) {
	n.ListRoutesCall.CallCount++

	return n.ListRoutesCall.Returns.Output, n.ListRoutesCall.Returns.Error
}

func (n *RoutesClient) DeleteRoute(route string) error {
	n.DeleteRouteCall.CallCount++
	n.DeleteRouteCall.Receives.Route = route

	return n.DeleteRouteCall.Returns.Error
}