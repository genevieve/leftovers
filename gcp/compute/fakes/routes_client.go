package fakes

import (
	"sync"

	gcpcompute "google.golang.org/api/compute/v1"
)

type RoutesClient struct {
	DeleteRouteCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Route string
		}
		Returns struct {
			Error error
		}
		Stub func(string) error
	}
	ListRoutesCall struct {
		sync.Mutex
		CallCount int
		Returns   struct {
			RouteSlice []*gcpcompute.Route
			Error      error
		}
		Stub func() ([]*gcpcompute.Route, error)
	}
}

func (f *RoutesClient) DeleteRoute(param1 string) error {
	f.DeleteRouteCall.Lock()
	defer f.DeleteRouteCall.Unlock()
	f.DeleteRouteCall.CallCount++
	f.DeleteRouteCall.Receives.Route = param1
	if f.DeleteRouteCall.Stub != nil {
		return f.DeleteRouteCall.Stub(param1)
	}
	return f.DeleteRouteCall.Returns.Error
}
func (f *RoutesClient) ListRoutes() ([]*gcpcompute.Route, error) {
	f.ListRoutesCall.Lock()
	defer f.ListRoutesCall.Unlock()
	f.ListRoutesCall.CallCount++
	if f.ListRoutesCall.Stub != nil {
		return f.ListRoutesCall.Stub()
	}
	return f.ListRoutesCall.Returns.RouteSlice, f.ListRoutesCall.Returns.Error
}
