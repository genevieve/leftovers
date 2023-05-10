package compute

import (
	"fmt"
	"github.com/genevieve/leftovers/common"
	gcpcompute "google.golang.org/api/compute/v1"
)

//go:generate faux --interface subnetworksClient --output fakes/subnetworks_client.go
type subnetworksClient interface {
	ListSubnetworks(region string) ([]*gcpcompute.Subnetwork, error)
	DeleteSubnetwork(region, network string) error
}

type Subnetworks struct {
	client  subnetworksClient
	logger  logger
	regions map[string]string
}

func NewSubnetworks(client subnetworksClient, logger logger, regions map[string]string) Subnetworks {
	return Subnetworks{
		client:  client,
		logger:  logger,
		regions: regions,
	}
}

func (n Subnetworks) List(filter string, regex bool) ([]common.Deletable, error) {
	subnetworks := []*gcpcompute.Subnetwork{}
	for _, region := range n.regions {
		n.logger.Debugf("Listing Subnetworks for region %s...\n", region)
		l, err := n.client.ListSubnetworks(region)
		if err != nil {
			return nil, fmt.Errorf("List Subnetworks for region %s: %s", region, err)
		}

		subnetworks = append(subnetworks, l...)
	}

	var resources []common.Deletable
	for _, subnetwork := range subnetworks {
		resource := NewSubnetwork(n.client, subnetwork.Name, n.regions[subnetwork.Region], subnetwork.Network)

		if subnetwork.Name == "default" {
			continue
		}

		if !common.ResourceMatches(subnetwork.Name, filter, regex) {
			continue
		}

		proceed := n.logger.PromptWithDetails(resource.Type(), resource.Name())
		if !proceed {
			continue
		}

		resources = append(resources, resource)
	}

	return resources, nil
}

func (s Subnetworks) Type() string {
	return "subnetwork"
}
