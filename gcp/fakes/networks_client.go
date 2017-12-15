package fakes

import compute "google.golang.org/api/compute/v1"

type NetworksClient struct {
	ListCall struct {
		CallCount int
		Receives  struct {
			Project string
		}
		Returns struct {
			Output *compute.NetworkList
			Error  error
		}
	}

	DeleteCall struct {
		CallCount int
		Receives  struct {
			Project string
			Network string
		}
		Returns struct {
			Output *compute.Operation
			Error  error
		}
	}
}

func (n *NetworksClient) List(project string) (*compute.NetworkList, error) {
	n.ListCall.CallCount++
	n.ListCall.Receives.Project = project

	return n.ListCall.Returns.Output, n.ListCall.Returns.Error
}

func (n *NetworksClient) Delete(project, network string) (*compute.Operation, error) {
	n.DeleteCall.CallCount++
	n.DeleteCall.Receives.Project = project
	n.DeleteCall.Receives.Network = network

	return n.DeleteCall.Returns.Output, n.DeleteCall.Returns.Error
}
