package fakes

import (
	"sync"

	gcpcompute "google.golang.org/api/compute/v1"
)

type SubnetworksClient struct {
	DeleteSubnetworkCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Region  string
			Network string
		}
		Returns struct {
			Error error
		}
		Stub func(string, string) error
	}
	ListSubnetworksCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Region string
		}
		Returns struct {
			SubnetworkSlice []*gcpcompute.Subnetwork
			Error           error
		}
		Stub func(string) ([]*gcpcompute.Subnetwork, error)
	}
}

func (f *SubnetworksClient) DeleteSubnetwork(param1 string, param2 string) error {
	f.DeleteSubnetworkCall.Lock()
	defer f.DeleteSubnetworkCall.Unlock()
	f.DeleteSubnetworkCall.CallCount++
	f.DeleteSubnetworkCall.Receives.Region = param1
	f.DeleteSubnetworkCall.Receives.Network = param2
	if f.DeleteSubnetworkCall.Stub != nil {
		return f.DeleteSubnetworkCall.Stub(param1, param2)
	}
	return f.DeleteSubnetworkCall.Returns.Error
}
func (f *SubnetworksClient) ListSubnetworks(param1 string) ([]*gcpcompute.Subnetwork, error) {
	f.ListSubnetworksCall.Lock()
	defer f.ListSubnetworksCall.Unlock()
	f.ListSubnetworksCall.CallCount++
	f.ListSubnetworksCall.Receives.Region = param1
	if f.ListSubnetworksCall.Stub != nil {
		return f.ListSubnetworksCall.Stub(param1)
	}
	return f.ListSubnetworksCall.Returns.SubnetworkSlice, f.ListSubnetworksCall.Returns.Error
}
