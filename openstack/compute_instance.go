package openstack

import "fmt"

type ComputeInstance struct {
	name   string
	id     string
	client ComputeClient
}

func NewComputeInstance(name string, id string, client ComputeClient) ComputeInstance {
	return ComputeInstance{
		name:   fmt.Sprintf("%s %s", name, id),
		id:     id,
		client: client,
	}
}

func (ci ComputeInstance) Name() string {
	return ci.name
}

func (ComputeInstance) Type() string {
	return "Compute Instance"
}

func (ci ComputeInstance) Delete() error {
	return ci.client.Delete(ci.id)
}
