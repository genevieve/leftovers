package fakes

import gcpcompute "google.golang.org/api/compute/v1"

type InstanceGroupsClient struct {
	ListInstanceGroupsCall struct {
		CallCount int
		Receives  struct {
			Zone   string
			Filter string
		}
		Returns struct {
			Output *gcpcompute.InstanceGroupList
			Error  error
		}
	}

	DeleteInstanceGroupCall struct {
		CallCount int
		Receives  struct {
			Zone          string
			InstanceGroup string
		}
		Returns struct {
			Error error
		}
	}
}

func (n *InstanceGroupsClient) ListInstanceGroups(zone, filter string) (*gcpcompute.InstanceGroupList, error) {
	n.ListInstanceGroupsCall.CallCount++
	n.ListInstanceGroupsCall.Receives.Zone = zone
	n.ListInstanceGroupsCall.Receives.Filter = filter

	return n.ListInstanceGroupsCall.Returns.Output, n.ListInstanceGroupsCall.Returns.Error
}

func (n *InstanceGroupsClient) DeleteInstanceGroup(zone, instanceGroup string) error {
	n.DeleteInstanceGroupCall.CallCount++
	n.DeleteInstanceGroupCall.Receives.Zone = zone
	n.DeleteInstanceGroupCall.Receives.InstanceGroup = instanceGroup

	return n.DeleteInstanceGroupCall.Returns.Error
}
