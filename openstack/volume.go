package openstack

import "fmt"

type Volume struct {
	name           string
	id             string
	VolumesDeleter VolumesDeleter
}

func NewVolume(name string, id string, volumesDeleter VolumesDeleter) Volume {
	return Volume{
		name:           name,
		id:             id,
		VolumesDeleter: volumesDeleter,
	}
}

func (volume Volume) Name() string {
	return fmt.Sprintf("%s %s", volume.name, volume.id)
}
func (volume Volume) Type() string {
	return "Volume"
}
func (volume Volume) Delete() error {
	return volume.VolumesDeleter.Delete(volume.id)
}
