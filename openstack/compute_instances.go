package openstack

import (
	"github.com/genevieve/leftovers/common"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
)

type ComputeInstanceLister interface {
	List() ([]servers.Server, error)
}

type ComputeInstanceDeleter interface {
	Delete(instanceID string) error
}

type ComputeInstances struct {
	provider ComputeInstanceProvider
	logger   logger
}

type ComputeInstanceProvider interface {
	GetComputeInstanceLister() ComputeInstanceLister
	GetComputeInstanceDeleter() ComputeInstanceDeleter
}

func NewComputeInstances(computeInstanceProvider ComputeInstanceProvider, logger logger) ComputeInstances {
	return ComputeInstances{
		provider: computeInstanceProvider,
		logger:   logger,
	}
}

func (ci ComputeInstances) List() ([]common.Deletable, error) {
	computeInstances, err := ci.provider.GetComputeInstanceLister().List()
	if err != nil {
		return nil, err
	}
	var deletables []common.Deletable
	for _, instance := range computeInstances {
		deletable := NewComputeInstance(instance.Name, instance.ID, ci.provider.GetComputeInstanceDeleter())
		if ci.logger.PromptWithDetails(deletable.Type(), deletable.Name()) {
			deletables = append(deletables, deletable)
		}
	}
	return deletables, nil
}

func (ci ComputeInstances) Type() string {
	return "Compute Instance"
}
