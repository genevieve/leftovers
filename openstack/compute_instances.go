package openstack

import (
	"github.com/genevieve/leftovers/common"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
)

type ComputeInstances struct {
	computeClient ComputeClient
	logger        logger
}

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

func (ci ComputeInstances) List() ([]common.Deletable, error) {
	computeInstances, err := ci.computeClient.List()
	if err != nil {
		return nil, err
	}

	var deletables []common.Deletable
	for _, instance := range computeInstances {
		deletable := NewComputeInstance(instance.Name, instance.ID, ci.computeClient)
		if ci.logger.PromptWithDetails(deletable.Type(), deletable.Name()) {
			deletables = append(deletables, deletable)
		}
	}

	return deletables, nil
}

func (ci ComputeInstances) Type() string {
	return "Compute Instance"
}
