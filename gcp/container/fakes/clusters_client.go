package fakes

import (
	"sync"

	gcpcontainer "google.golang.org/api/container/v1"
)

type ClustersClient struct {
	DeleteClusterCall struct {
		sync.Mutex
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
		sync.Mutex
		CallCount int
		Receives  struct {
			Zone string
		}
		Returns struct {
			ListClustersResponse *gcpcontainer.ListClustersResponse
			Error                error
		}
		Stub func(string) (*gcpcontainer.ListClustersResponse, error)
	}
}

func (f *ClustersClient) DeleteCluster(param1 string, param2 string) error {
	f.DeleteClusterCall.Lock()
	defer f.DeleteClusterCall.Unlock()
	f.DeleteClusterCall.CallCount++
	f.DeleteClusterCall.Receives.Zone = param1
	f.DeleteClusterCall.Receives.Cluster = param2
	if f.DeleteClusterCall.Stub != nil {
		return f.DeleteClusterCall.Stub(param1, param2)
	}
	return f.DeleteClusterCall.Returns.Error
}
func (f *ClustersClient) ListClusters(param1 string) (*gcpcontainer.ListClustersResponse, error) {
	f.ListClustersCall.Lock()
	defer f.ListClustersCall.Unlock()
	f.ListClustersCall.CallCount++
	f.ListClustersCall.Receives.Zone = param1
	if f.ListClustersCall.Stub != nil {
		return f.ListClustersCall.Stub(param1)
	}
	return f.ListClustersCall.Returns.ListClustersResponse, f.ListClustersCall.Returns.Error
}
