package fakes

import gcpcompute "google.golang.org/api/compute/v1"

type RoutersClient struct {
	ListRoutersCall struct {
		CallCount int
		Returns   struct {
			Output []*gcpcompute.Router
			Error  error
		}
	}

	DeleteRouterCall struct {
		CallCount int
		Receives  struct {
			Router string
		}
		Returns struct {
			Error error
		}
	}
}

func (n *RoutersClient) ListRouters() ([]*gcpcompute.Router, error) {
	n.ListRoutersCall.CallCount++

	return n.ListRoutersCall.Returns.Output, n.ListRoutersCall.Returns.Error
}

func (n *RoutersClient) DeleteRouter(router string) error {
	n.DeleteRouterCall.CallCount++
	n.DeleteRouterCall.Receives.Router = router

	return n.DeleteRouterCall.Returns.Error
}
