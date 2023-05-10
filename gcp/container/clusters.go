package container

import (
	"fmt"
	"strings"

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
	ListClusters() (*gcpcontainer.ListClustersResponse, error)
	DeleteCluster(zone, cluster string) error
}

func NewClusters(client clustersClient, zones map[string]string, logger logger) Clusters {
	return Clusters{
		client: client,
		zones:  zones,
		logger: logger,
	}
}

func (c Clusters) List(filter string) ([]common.Deletable, error) {
	c.logger.Debugf("Listing Clusters for all Zones...\n")
	resp, err := c.client.ListClusters()
	if err != nil {
		return nil, fmt.Errorf("List Clusters for all Zones: %w", err)
	}
	clusters := resp.Clusters

	deletables := []common.Deletable{}
	for _, cluster := range clusters {
		resource := NewCluster(c.client, cluster.Zone, cluster.Name)

		if !strings.Contains(resource.Name(), filter) {
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
