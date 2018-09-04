package fakes

import gcpcompute "google.golang.org/api/compute/v1"

type InstanceGroupManagersClient struct {
	ListInstanceGroupManagersCall struct {
		CallCount int
		Receives  struct {
			Zone string
		}
		Returns struct {
			Output []*gcpcompute.InstanceGroupManager
			Error  error
		}
	}

	DeleteInstanceGroupManagerCall struct {
		CallCount int
		Receives  struct {
			Zone                 string
			InstanceGroupManager string
		}
		Returns struct {
			Error error
		}
	}
}

func (n *InstanceGroupManagersClient) ListInstanceGroupManagers(zone string) ([]*gcpcompute.InstanceGroupManager, error) {
	n.ListInstanceGroupManagersCall.CallCount++
	n.ListInstanceGroupManagersCall.Receives.Zone = zone

	return n.ListInstanceGroupManagersCall.Returns.Output, n.ListInstanceGroupManagersCall.Returns.Error
}

func (n *InstanceGroupManagersClient) DeleteInstanceGroupManager(zone, instanceGroupManager string) error {
	n.DeleteInstanceGroupManagerCall.CallCount++
	n.DeleteInstanceGroupManagerCall.Receives.Zone = zone
	n.DeleteInstanceGroupManagerCall.Receives.InstanceGroupManager = instanceGroupManager

	return n.DeleteInstanceGroupManagerCall.Returns.Error
}
