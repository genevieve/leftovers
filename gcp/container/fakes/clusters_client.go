package fakes

import (
	gcpcontainer "google.golang.org/api/container/v1"
)

type ClustersClient struct {
	ListClustersCall struct {
		CallCount int
		Receives  struct {
			Zone string
		}
		Returns struct {
			Output *gcpcontainer.ListClustersResponse
			Error  error
		}
	}
	DeleteClusterCall struct {
		CallCount int
		Receives  struct {
			Zone    string
			Cluster string
		}
		Returns struct {
			Error error
		}
	}
}

func (c *ClustersClient) ListClusters(zone string) (*gcpcontainer.ListClustersResponse, error) {
	c.ListClustersCall.CallCount++
	c.ListClustersCall.Receives.Zone = zone

	return c.ListClustersCall.Returns.Output, c.ListClustersCall.Returns.Error
}

func (c *ClustersClient) DeleteCluster(zone string, cluster string) error {
	c.DeleteClusterCall.CallCount++
	c.DeleteClusterCall.Receives.Zone = zone
	c.DeleteClusterCall.Receives.Cluster = cluster

	return c.DeleteClusterCall.Returns.Error
}
