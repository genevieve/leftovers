package fakes

import compute "google.golang.org/api/compute/v1"

type SubnetworksClient struct {
	ListSubnetworksCall struct {
		CallCount int
		Receives  struct {
			Region string
		}
		Returns struct {
			Output []*compute.Subnetwork
			Error  error
		}
	}

	DeleteSubnetworkCall struct {
		CallCount int
		Receives  struct {
			Region     string
			Subnetwork string
		}
		Returns struct {
			Error error
		}
	}
}

func (n *SubnetworksClient) ListSubnetworks(region string) ([]*compute.Subnetwork, error) {
	n.ListSubnetworksCall.CallCount++
	n.ListSubnetworksCall.Receives.Region = region

	return n.ListSubnetworksCall.Returns.Output, n.ListSubnetworksCall.Returns.Error
}

func (n *SubnetworksClient) DeleteSubnetwork(region, subnetwork string) error {
	n.DeleteSubnetworkCall.CallCount++
	n.DeleteSubnetworkCall.Receives.Region = region
	n.DeleteSubnetworkCall.Receives.Subnetwork = subnetwork

	return n.DeleteSubnetworkCall.Returns.Error
}
