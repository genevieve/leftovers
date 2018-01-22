package fakes

import gcpcompute "google.golang.org/api/compute/v1"

type TargetHttpsProxiesClient struct {
	ListTargetHttpsProxiesCall struct {
		CallCount int
		Returns   struct {
			Output *gcpcompute.TargetHttpsProxyList
			Error  error
		}
	}

	DeleteTargetHttpsProxyCall struct {
		CallCount int
		Receives  struct {
			TargetHttpsProxy string
		}
		Returns struct {
			Error error
		}
	}
}

func (t *TargetHttpsProxiesClient) ListTargetHttpsProxies() (*gcpcompute.TargetHttpsProxyList, error) {
	t.ListTargetHttpsProxiesCall.CallCount++

	return t.ListTargetHttpsProxiesCall.Returns.Output, t.ListTargetHttpsProxiesCall.Returns.Error
}

func (t *TargetHttpsProxiesClient) DeleteTargetHttpsProxy(targetHttpsProxy string) error {
	t.DeleteTargetHttpsProxyCall.CallCount++
	t.DeleteTargetHttpsProxyCall.Receives.TargetHttpsProxy = targetHttpsProxy

	return t.DeleteTargetHttpsProxyCall.Returns.Error
}
