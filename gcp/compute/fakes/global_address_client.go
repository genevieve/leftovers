package fakes

import (
	"sync"

	gcpcompute "google.golang.org/api/compute/v1"
)

type GlobalAddressesClient struct {
	DeleteGlobalAddressCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Address string
		}
		Returns struct {
			Error error
		}
		Stub func(string) error
	}
	ListGlobalAddressesCall struct {
		sync.Mutex
		CallCount int
		Returns   struct {
			AddressSlice []*gcpcompute.Address
			Error        error
		}
		Stub func() ([]*gcpcompute.Address, error)
	}
}

func (f *GlobalAddressesClient) DeleteGlobalAddress(param1 string) error {
	f.DeleteGlobalAddressCall.Lock()
	defer f.DeleteGlobalAddressCall.Unlock()
	f.DeleteGlobalAddressCall.CallCount++
	f.DeleteGlobalAddressCall.Receives.Address = param1
	if f.DeleteGlobalAddressCall.Stub != nil {
		return f.DeleteGlobalAddressCall.Stub(param1)
	}
	return f.DeleteGlobalAddressCall.Returns.Error
}
func (f *GlobalAddressesClient) ListGlobalAddresses() ([]*gcpcompute.Address, error) {
	f.ListGlobalAddressesCall.Lock()
	defer f.ListGlobalAddressesCall.Unlock()
	f.ListGlobalAddressesCall.CallCount++
	if f.ListGlobalAddressesCall.Stub != nil {
		return f.ListGlobalAddressesCall.Stub()
	}
	return f.ListGlobalAddressesCall.Returns.AddressSlice, f.ListGlobalAddressesCall.Returns.Error
}
