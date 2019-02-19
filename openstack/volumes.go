package openstack

import (
	"github.com/genevieve/leftovers/common"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumes"
)

//go:generate counterfeiter . VolumesLister
type VolumesLister interface {
	List() ([]volumes.Volume, error)
}

//go:generate counterfeiter . VolumesDeleter
type VolumesDeleter interface {
	Delete(volumeID string) error
}

//go:generate counterfeiter . VolumesServiceProvider
type VolumesServiceProvider interface {
	GetVolumesLister() VolumesLister
	GetVolumesDeleter() VolumesDeleter
}

type Volumes struct {
	volumesServiceProvider VolumesServiceProvider
}

func NewVolumes(volumesServiceProvider VolumesServiceProvider) (Volumes, error) {
	return Volumes{volumesServiceProvider}, nil
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
		deletables = append(deletables, NewVolume(volume.Name, volume.ID, serviceProvider.GetVolumesDeleter()))
	}

	return deletables, nil
}
