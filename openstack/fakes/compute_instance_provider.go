package fakes

import "github.com/genevieve/leftovers/openstack"

type ComputeInstanceProvider struct {
	GetComputeInstanceListerCall struct {
		CallCount int
		Returns   struct {
			Lister openstack.ComputeInstanceLister
		}
	}

	GetComputeInstanceDeleterCall struct {
		CallCount int
		Returns   struct {
			Deleter openstack.ComputeInstanceDeleter
		}
	}
}

func (provider *ComputeInstanceProvider) GetComputeInstanceLister() openstack.ComputeInstanceLister {
	provider.GetComputeInstanceListerCall.CallCount++

	return provider.GetComputeInstanceListerCall.Returns.Lister
}

func (provider *ComputeInstanceProvider) GetComputeInstanceDeleter() openstack.ComputeInstanceDeleter {
	provider.GetComputeInstanceDeleterCall.CallCount++

	return provider.GetComputeInstanceDeleterCall.Returns.Deleter
}
