package openstack

import "fmt"

type ComputeInstance struct {
	name          string
	id            string
	computeClient ComputeClient
}

func NewComputeInstance(name string, id string, computeClient ComputeClient) ComputeInstance {
	return ComputeInstance{
		name:          name,
		id:            id,
		computeClient: computeClient,
	}
}

func (ci ComputeInstance) Name() string {
	return fmt.Sprintf("%s %s", ci.name, ci.id)
}

func (ci ComputeInstance) Type() string {
	return "Compute Instance"
}

func (ci ComputeInstance) Delete() error {
	return ci.computeClient.Delete(ci.id)

}
