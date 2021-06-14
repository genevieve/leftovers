package openstack

import (
	"fmt"
	"github.com/genevieve/leftovers/common"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
)

type ComputeInstances struct {
	computeClient ComputeClient
	logger        logger
}

//go:generate faux --interface ComputeClient --output fakes/compute_client.go
type ComputeClient interface {
	List() ([]servers.Server, error)
	Delete(instanceID string) error
}

func NewComputeInstances(computeClient ComputeClient, logger logger) ComputeInstances {
	return ComputeInstances{
		computeClient: computeClient,
		logger:        logger,
	}
}

func (ci ComputeInstances) List(filter string, regex bool) ([]common.Deletable, error) {
	ci.logger.Debugln("Listing Compute Instances...")

	computeInstances, err := ci.computeClient.List()
	if err != nil {
		return nil, fmt.Errorf("List Compute Instances: %s", err)
	}

	var resources []common.Deletable
	for _, instance := range computeInstances {
		r := NewComputeInstance(instance.Name, instance.ID, ci.computeClient)

		if !common.MatchRegex(instance.Name, filter, regex) {
			continue
		}

		proceed := ci.logger.PromptWithDetails(r.Type(), r.Name())
		if !proceed {
			continue
		}

		resources = append(resources, r)
	}

	return resources, nil
}

func (ci ComputeInstances) Type() string {
	return "Compute Instance"
}
