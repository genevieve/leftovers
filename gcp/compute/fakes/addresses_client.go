package fakes

import (
	"sync"

	gcpcompute "google.golang.org/api/compute/v1"
)

type AddressesClient struct {
	DeleteAddressCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Region  string
			Address string
		}
		Returns struct {
			Error error
		}
		Stub func(string, string) error
	}
	ListAddressesCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Region string
		}
		Returns struct {
			AddressSlice []*gcpcompute.Address
			Error        error
		}
		Stub func(string) ([]*gcpcompute.Address, error)
	}
}

func (f *AddressesClient) DeleteAddress(param1 string, param2 string) error {
	f.DeleteAddressCall.Lock()
	defer f.DeleteAddressCall.Unlock()
	f.DeleteAddressCall.CallCount++
	f.DeleteAddressCall.Receives.Region = param1
	f.DeleteAddressCall.Receives.Address = param2
	if f.DeleteAddressCall.Stub != nil {
		return f.DeleteAddressCall.Stub(param1, param2)
	}
	return f.DeleteAddressCall.Returns.Error
}
func (f *AddressesClient) ListAddresses(param1 string) ([]*gcpcompute.Address, error) {
	f.ListAddressesCall.Lock()
	defer f.ListAddressesCall.Unlock()
	f.ListAddressesCall.CallCount++
	f.ListAddressesCall.Receives.Region = param1
	if f.ListAddressesCall.Stub != nil {
		return f.ListAddressesCall.Stub(param1)
	}
	return f.ListAddressesCall.Returns.AddressSlice, f.ListAddressesCall.Returns.Error
}
