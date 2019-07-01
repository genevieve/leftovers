package fakes

import gcpcompute "google.golang.org/api/compute/v1"

type RoutersClient struct {
	ListRoutersCall struct {
		CallCount int
		Receives  struct {
			Region string
		}
		Returns struct {
			Output []*gcpcompute.Router
			Error  error
		}
	}

	DeleteRouterCall struct {
		CallCount int
		Receives  struct {
			Router string
			Region string
		}
		Returns struct {
			Error error
		}
	}
}

func (n *RoutersClient) ListRouters(region string) ([]*gcpcompute.Router, error) {
	n.ListRoutersCall.CallCount++
	n.ListRoutersCall.Receives.Region = region

	return n.ListRoutersCall.Returns.Output, n.ListRoutersCall.Returns.Error
}

func (n *RoutersClient) DeleteRouter(region, router string) error {
	n.DeleteRouterCall.CallCount++
	n.DeleteRouterCall.Receives.Region = region
	n.DeleteRouterCall.Receives.Router = router

	return n.DeleteRouterCall.Returns.Error
}
