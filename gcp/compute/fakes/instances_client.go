package fakes

import gcpcompute "google.golang.org/api/compute/v1"

type InstancesClient struct {
	ListInstancesCall struct {
		CallCount int
		Receives  struct {
			Zone string
		}
		Returns struct {
			Output *gcpcompute.InstanceList
			Error  error
		}
	}

	DeleteInstanceCall struct {
		CallCount int
		Receives  struct {
			Zone     string
			Instance string
		}
		Returns struct {
			Error error
		}
	}
}

func (n *InstancesClient) ListInstances(zone string) (*gcpcompute.InstanceList, error) {
	n.ListInstancesCall.CallCount++
	n.ListInstancesCall.Receives.Zone = zone

	return n.ListInstancesCall.Returns.Output, n.ListInstancesCall.Returns.Error
}

func (n *InstancesClient) DeleteInstance(zone, instance string) error {
	n.DeleteInstanceCall.CallCount++
	n.DeleteInstanceCall.Receives.Zone = zone
	n.DeleteInstanceCall.Receives.Instance = instance

	return n.DeleteInstanceCall.Returns.Error
}
