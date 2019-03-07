package openstack

import "fmt"

type ComputeInstance struct {
	name    string
	id      string
	deleter ComputeInstanceDeleter
}

func NewComputeInstance(name string, id string, deleter ComputeInstanceDeleter) ComputeInstance {
	return ComputeInstance{name, id, deleter}
}

func (ci ComputeInstance) Name() string {
	return fmt.Sprintf("%s %s", ci.name, ci.id)
}

func (ci ComputeInstance) Type() string {
	return "Compute Instance"
}

func (ci ComputeInstance) Delete() error {
	return ci.deleter.Delete(ci.id)
}
