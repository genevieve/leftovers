package fakes

import (
	"sync"

	gcp "google.golang.org/api/compute/v1"
)

type TargetHttpProxiesClient struct {
	DeleteTargetHttpProxyCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			TargetHttpProxy string
		}
		Returns struct {
			Error error
		}
		Stub func(string) error
	}
	ListTargetHttpProxiesCall struct {
		sync.Mutex
		CallCount int
		Returns   struct {
			TargetHttpProxyList *gcp.TargetHttpProxyList
			Error               error
		}
		Stub func() (*gcp.TargetHttpProxyList, error)
	}
}

func (f *TargetHttpProxiesClient) DeleteTargetHttpProxy(param1 string) error {
	f.DeleteTargetHttpProxyCall.Lock()
	defer f.DeleteTargetHttpProxyCall.Unlock()
	f.DeleteTargetHttpProxyCall.CallCount++
	f.DeleteTargetHttpProxyCall.Receives.TargetHttpProxy = param1
	if f.DeleteTargetHttpProxyCall.Stub != nil {
		return f.DeleteTargetHttpProxyCall.Stub(param1)
	}
	return f.DeleteTargetHttpProxyCall.Returns.Error
}
func (f *TargetHttpProxiesClient) ListTargetHttpProxies() (*gcp.TargetHttpProxyList, error) {
	f.ListTargetHttpProxiesCall.Lock()
	defer f.ListTargetHttpProxiesCall.Unlock()
	f.ListTargetHttpProxiesCall.CallCount++
	if f.ListTargetHttpProxiesCall.Stub != nil {
		return f.ListTargetHttpProxiesCall.Stub()
	}
	return f.ListTargetHttpProxiesCall.Returns.TargetHttpProxyList, f.ListTargetHttpProxiesCall.Returns.Error
}
