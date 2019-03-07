package fakes

import (
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
)

type ComputeInstanceLister struct {
	ListCall struct {
		CallCount int
		Returns   struct {
			ComputeInstances []servers.Server
			Error            error
		}
	}
}

func (lister *ComputeInstanceLister) List() ([]servers.Server, error) {
	lister.ListCall.CallCount++

	return lister.ListCall.Returns.ComputeInstances, lister.ListCall.Returns.Error
}
