package openstack

import (
	"github.com/genevieve/leftovers/common"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumes"
)

type VolumesLister interface {
	List() ([]volumes.Volume, error)
}

type VolumesDeleter interface {
	Delete(volumeID string) error
}

type VolumesServiceProvider interface {
	GetVolumesLister() VolumesLister
	GetVolumesDeleter() VolumesDeleter
}

type Volumes struct {
	volumesServiceProvider VolumesServiceProvider
	logger                 logger
}

func NewVolumes(volumesServiceProvider VolumesServiceProvider, logger logger) Volumes {
	return Volumes{volumesServiceProvider, logger}
}

func (volumes Volumes) Type() string {
	return "Volume"
}

func (volumes Volumes) List() ([]common.Deletable, error) {
	serviceProvider := volumes.volumesServiceProvider
	result, err := serviceProvider.GetVolumesLister().List()

	if err != nil {
		return nil, err
	}

	var deletables []common.Deletable
	for _, volume := range result {
		deletable := NewVolume(volume.Name, volume.ID, serviceProvider.GetVolumesDeleter())
		confirm := volumes.logger.PromptWithDetails(deletable.Type(), deletable.Name())

		if confirm {
			deletables = append(deletables, NewVolume(volume.Name, volume.ID, serviceProvider.GetVolumesDeleter()))
		}
	}

	return deletables, nil
}
