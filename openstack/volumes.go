package openstack

import (
	"fmt"

	"github.com/genevieve/leftovers/common"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumes"
)

//go:generate faux --interface volumesClient --output fakes/volumes_client.go
type volumesClient interface {
	List() ([]volumes.Volume, error)
	Delete(volumeID string) error
}

type Volumes struct {
	client volumesClient
	logger logger
}

func NewVolumes(client volumesClient, logger logger) Volumes {
	return Volumes{
		client: client,
		logger: logger,
	}
}

func (v Volumes) List(filter string) ([]common.Deletable, error) {
	v.logger.Debugln("Listing Volumes...")

	result, err := v.client.List()
	if err != nil {
		return nil, fmt.Errorf("List Volumes: %s", err)
	}

	var resources []common.Deletable
	for _, volume := range result {
		r := NewVolume(volume.Name, volume.ID, v.client)

		proceed := v.logger.PromptWithDetails(r.Type(), r.Name())
		if !proceed {
			continue
		}

		resources = append(resources, r)
	}

	return resources, nil
}

func (Volumes) Type() string {
	return "Volume"
}
