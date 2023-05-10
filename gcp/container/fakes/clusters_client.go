package fakes

import (
	"sync"

	"google.golang.org/api/container/v1"
)

type ClustersClient struct {
	DeleteClusterCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Zone    string
			Cluster string
		}
		Returns struct {
			Error error
		}
		Stub func(string, string) error
	}
	ListClustersCall struct {
		mutex     sync.Mutex
		CallCount int
		Returns   struct {
			ListClustersResponse *container.ListClustersResponse
			Error                error
		}
		Stub func() (*container.ListClustersResponse, error)
	}
}

func (f *ClustersClient) DeleteCluster(param1 string, param2 string) error {
	f.DeleteClusterCall.mutex.Lock()
	defer f.DeleteClusterCall.mutex.Unlock()
	f.DeleteClusterCall.CallCount++
	f.DeleteClusterCall.Receives.Zone = param1
	f.DeleteClusterCall.Receives.Cluster = param2
	if f.DeleteClusterCall.Stub != nil {
		return f.DeleteClusterCall.Stub(param1, param2)
	}
	return f.DeleteClusterCall.Returns.Error
}
func (f *ClustersClient) ListClusters() (*container.ListClustersResponse, error) {
	f.ListClustersCall.mutex.Lock()
	defer f.ListClustersCall.mutex.Unlock()
	f.ListClustersCall.CallCount++
	if f.ListClustersCall.Stub != nil {
		return f.ListClustersCall.Stub()
	}
	return f.ListClustersCall.Returns.ListClustersResponse, f.ListClustersCall.Returns.Error
}
