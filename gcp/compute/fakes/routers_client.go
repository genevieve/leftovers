package fakes

import (
	"sync"

	gcpcompute "google.golang.org/api/compute/v1"
)

type RoutersClient struct {
	DeleteRouterCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Region string
			Router string
		}
		Returns struct {
			Error error
		}
		Stub func(string, string) error
	}
	ListRoutersCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Region string
		}
		Returns struct {
			RouterSlice []*gcpcompute.Router
			Error       error
		}
		Stub func(string) ([]*gcpcompute.Router, error)
	}
}

func (f *RoutersClient) DeleteRouter(param1 string, param2 string) error {
	f.DeleteRouterCall.Lock()
	defer f.DeleteRouterCall.Unlock()
	f.DeleteRouterCall.CallCount++
	f.DeleteRouterCall.Receives.Region = param1
	f.DeleteRouterCall.Receives.Router = param2
	if f.DeleteRouterCall.Stub != nil {
		return f.DeleteRouterCall.Stub(param1, param2)
	}
	return f.DeleteRouterCall.Returns.Error
}
func (f *RoutersClient) ListRouters(param1 string) ([]*gcpcompute.Router, error) {
	f.ListRoutersCall.Lock()
	defer f.ListRoutersCall.Unlock()
	f.ListRoutersCall.CallCount++
	f.ListRoutersCall.Receives.Region = param1
	if f.ListRoutersCall.Stub != nil {
		return f.ListRoutersCall.Stub(param1)
	}
	return f.ListRoutersCall.Returns.RouterSlice, f.ListRoutersCall.Returns.Error
}
