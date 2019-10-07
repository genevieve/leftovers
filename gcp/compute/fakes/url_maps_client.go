package fakes

import (
	"sync"

	gcpcompute "google.golang.org/api/compute/v1"
)

type UrlMapsClient struct {
	DeleteUrlMapCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			UrlMap string
		}
		Returns struct {
			Error error
		}
		Stub func(string) error
	}
	ListUrlMapsCall struct {
		sync.Mutex
		CallCount int
		Returns   struct {
			UrlMapList *gcpcompute.UrlMapList
			Error      error
		}
		Stub func() (*gcpcompute.UrlMapList, error)
	}
}

func (f *UrlMapsClient) DeleteUrlMap(param1 string) error {
	f.DeleteUrlMapCall.Lock()
	defer f.DeleteUrlMapCall.Unlock()
	f.DeleteUrlMapCall.CallCount++
	f.DeleteUrlMapCall.Receives.UrlMap = param1
	if f.DeleteUrlMapCall.Stub != nil {
		return f.DeleteUrlMapCall.Stub(param1)
	}
	return f.DeleteUrlMapCall.Returns.Error
}
func (f *UrlMapsClient) ListUrlMaps() (*gcpcompute.UrlMapList, error) {
	f.ListUrlMapsCall.Lock()
	defer f.ListUrlMapsCall.Unlock()
	f.ListUrlMapsCall.CallCount++
	if f.ListUrlMapsCall.Stub != nil {
		return f.ListUrlMapsCall.Stub()
	}
	return f.ListUrlMapsCall.Returns.UrlMapList, f.ListUrlMapsCall.Returns.Error
}
