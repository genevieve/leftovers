package openstack

import "fmt"

type Volume struct {
	name   string
	id     string
	client volumesClient
}

func NewVolume(name string, id string, client volumesClient) Volume {
	return Volume{
		name:   name,
		id:     id,
		client: client,
	}
}

func (v Volume) Name() string {
	return fmt.Sprintf("%s %s", v.name, v.id)
}

func (Volume) Type() string {
	return "Volume"
}

func (v Volume) Delete() error {
	return v.client.Delete(v.id)
}
