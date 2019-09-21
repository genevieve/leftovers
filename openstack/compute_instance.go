package openstack

import "fmt"

type ComputeInstance struct {
	name          string
	id            string
	computeClient ComputeClient
}

func NewComputeInstance(name string, id string, computeClient ComputeClient) ComputeInstance {
	return ComputeInstance{
		name:          fmt.Sprintf("%s %s", name, id),
		id:            id,
		computeClient: computeClient,
	}
}

func (ci ComputeInstance) Name() string {
	return ci.name
}

func (ci ComputeInstance) Type() string {
	return "Compute Instance"
}

func (ci ComputeInstance) Delete() error {
	return ci.computeClient.Delete(ci.id)

}
