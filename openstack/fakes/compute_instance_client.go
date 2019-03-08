package fakes

import "github.com/gophercloud/gophercloud/openstack/compute/v2/servers"

type ComputeInstanceClient struct {
	DeleteCall struct {
		CallCount int
		Returns   struct {
			Error error
		}
		Receives struct {
			InstanceID string
		}
	}
	ListCall struct {
		CallCount int
		Returns   struct {
			ComputeInstances []servers.Server
			Error            error
		}
	}
}

func (client *ComputeInstanceClient) Delete(instanceID string) error {
	client.DeleteCall.CallCount++
	client.DeleteCall.Receives.InstanceID = instanceID

	return client.DeleteCall.Returns.Error
}

func (client *ComputeInstanceClient) List() ([]servers.Server, error) {
	client.ListCall.CallCount++

	return client.ListCall.Returns.ComputeInstances, client.ListCall.Returns.Error
}
