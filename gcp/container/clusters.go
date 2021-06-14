package container

import (
	"fmt"

	"github.com/genevieve/leftovers/common"
	gcpcontainer "google.golang.org/api/container/v1"
)

type Clusters struct {
	client clustersClient
	zones  map[string]string
	logger logger
}

//go:generate faux --interface clustersClient --output fakes/clusters_client.go
type clustersClient interface {
	ListClusters(zone string) (*gcpcontainer.ListClustersResponse, error)
	DeleteCluster(zone, cluster string) error
}

func NewClusters(client clustersClient, zones map[string]string, logger logger) Clusters {
	return Clusters{
		client: client,
		zones:  zones,
		logger: logger,
	}
}

func (c Clusters) List(filter string, regex bool) ([]common.Deletable, error) {
	clusters := []*gcpcontainer.Cluster{}
	for _, zone := range c.zones {
		c.logger.Debugf("Listing Clusters for Zone %s...\n", zone)
		resp, err := c.client.ListClusters(zone)
		if err != nil {
			return nil, fmt.Errorf("List Clusters for Zone %s: %s", zone, err)
		}
		clusters = append(clusters, resp.Clusters...)
	}

	deletables := []common.Deletable{}
	for _, cluster := range clusters {
		resource := NewCluster(c.client, cluster.Zone, cluster.Name)

		if !common.MatchRegex(resource.Name(), filter, regex) {
			continue
		}

		proceed := c.logger.PromptWithDetails(resource.Type(), resource.Name())
		if !proceed {
			continue
		}

		deletables = append(deletables, resource)
	}

	return deletables, nil
}

func (c Clusters) Type() string {
	return "cluster"
}
