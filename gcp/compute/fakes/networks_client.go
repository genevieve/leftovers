package fakes

import (
	"sync"

	compute "google.golang.org/api/compute/v1"
)

type NetworksClient struct {
	DeleteNetworkCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Network string
		}
		Returns struct {
			Error error
		}
		Stub func(string) error
	}
	ListNetworksCall struct {
		sync.Mutex
		CallCount int
		Returns   struct {
			NetworkSlice []*compute.Network
			Error        error
		}
		Stub func() ([]*compute.Network, error)
	}
}

func (f *NetworksClient) DeleteNetwork(param1 string) error {
	f.DeleteNetworkCall.Lock()
	defer f.DeleteNetworkCall.Unlock()
	f.DeleteNetworkCall.CallCount++
	f.DeleteNetworkCall.Receives.Network = param1
	if f.DeleteNetworkCall.Stub != nil {
		return f.DeleteNetworkCall.Stub(param1)
	}
	return f.DeleteNetworkCall.Returns.Error
}
func (f *NetworksClient) ListNetworks() ([]*compute.Network, error) {
	f.ListNetworksCall.Lock()
	defer f.ListNetworksCall.Unlock()
	f.ListNetworksCall.CallCount++
	if f.ListNetworksCall.Stub != nil {
		return f.ListNetworksCall.Stub()
	}
	return f.ListNetworksCall.Returns.NetworkSlice, f.ListNetworksCall.Returns.Error
}
