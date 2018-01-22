package fakes

import gcpcompute "google.golang.org/api/compute/v1"

type TargetHttpProxiesClient struct {
	ListTargetHttpProxiesCall struct {
		CallCount int
		Returns   struct {
			Output *gcpcompute.TargetHttpProxyList
			Error  error
		}
	}

	DeleteTargetHttpProxyCall struct {
		CallCount int
		Receives  struct {
			TargetHttpProxy string
		}
		Returns struct {
			Error error
		}
	}
}

func (t *TargetHttpProxiesClient) ListTargetHttpProxies() (*gcpcompute.TargetHttpProxyList, error) {
	t.ListTargetHttpProxiesCall.CallCount++

	return t.ListTargetHttpProxiesCall.Returns.Output, t.ListTargetHttpProxiesCall.Returns.Error
}

func (t *TargetHttpProxiesClient) DeleteTargetHttpProxy(targetHttpProxy string) error {
	t.DeleteTargetHttpProxyCall.CallCount++
	t.DeleteTargetHttpProxyCall.Receives.TargetHttpProxy = targetHttpProxy

	return t.DeleteTargetHttpProxyCall.Returns.Error
}
