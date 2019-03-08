package openstack

import "fmt"

type Volume struct {
	name          string
	id            string
	volumesClient VolumesClient
}

func NewVolume(name string, id string, volumesClient VolumesClient) Volume {
	return Volume{
		name:          name,
		id:            id,
		volumesClient: volumesClient,
	}
}

func (volume Volume) Name() string {
	return fmt.Sprintf("%s %s", volume.name, volume.id)
}
func (volume Volume) Type() string {
	return "Volume"
}
func (volume Volume) Delete() error {
	return volume.volumesClient.Delete(volume.id)
}
