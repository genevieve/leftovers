package openstack

import (
	"github.com/genevieve/leftovers/common"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumes"
)

type VolumesClient interface {
	List() ([]volumes.Volume, error)
	Delete(volumeID string) error
}

type Volumes struct {
	volumesClient VolumesClient
	logger        logger
}

func NewVolumes(volumesClient VolumesClient, logger logger) Volumes {
	return Volumes{
		volumesClient: volumesClient,
		logger:        logger,
	}
}

func (volumes Volumes) Type() string {
	return "Volume"
}

func (volumes Volumes) List() ([]common.Deletable, error) {
	result, err := volumes.volumesClient.List()

	if err != nil {
		return nil, err
	}

	var deletables []common.Deletable
	for _, volume := range result {
		deletable := NewVolume(volume.Name, volume.ID, volumes.volumesClient)
		confirm := volumes.logger.PromptWithDetails(deletable.Type(), deletable.Name())

		if confirm {
			deletables = append(deletables, deletable)
		}
	}

	return deletables, nil
}
