package fakes

import (
	"sync"

	gcpcompute "google.golang.org/api/compute/v1"
)

type TargetHttpsProxiesClient struct {
	DeleteTargetHttpsProxyCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			TargetHttpsProxy string
		}
		Returns struct {
			Error error
		}
		Stub func(string) error
	}
	ListTargetHttpsProxiesCall struct {
		sync.Mutex
		CallCount int
		Returns   struct {
			TargetHttpsProxyList *gcpcompute.TargetHttpsProxyList
			Error                error
		}
		Stub func() (*gcpcompute.TargetHttpsProxyList, error)
	}
}

func (f *TargetHttpsProxiesClient) DeleteTargetHttpsProxy(param1 string) error {
	f.DeleteTargetHttpsProxyCall.Lock()
	defer f.DeleteTargetHttpsProxyCall.Unlock()
	f.DeleteTargetHttpsProxyCall.CallCount++
	f.DeleteTargetHttpsProxyCall.Receives.TargetHttpsProxy = param1
	if f.DeleteTargetHttpsProxyCall.Stub != nil {
		return f.DeleteTargetHttpsProxyCall.Stub(param1)
	}
	return f.DeleteTargetHttpsProxyCall.Returns.Error
}
func (f *TargetHttpsProxiesClient) ListTargetHttpsProxies() (*gcpcompute.TargetHttpsProxyList, error) {
	f.ListTargetHttpsProxiesCall.Lock()
	defer f.ListTargetHttpsProxiesCall.Unlock()
	f.ListTargetHttpsProxiesCall.CallCount++
	if f.ListTargetHttpsProxiesCall.Stub != nil {
		return f.ListTargetHttpsProxiesCall.Stub()
	}
	return f.ListTargetHttpsProxiesCall.Returns.TargetHttpsProxyList, f.ListTargetHttpsProxiesCall.Returns.Error
}
