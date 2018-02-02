package fakes

import gcpcompute "google.golang.org/api/compute/v1"

type GlobalAddressesClient struct {
	ListGlobalAddressesCall struct {
		CallCount int
		Returns   struct {
			Output *gcpcompute.AddressList
			Error  error
		}
	}

	DeleteGlobalAddressCall struct {
		CallCount int
		Receives  struct {
			Address string
		}
		Returns struct {
			Error error
		}
	}
}

func (a *GlobalAddressesClient) ListGlobalAddresses() (*gcpcompute.AddressList, error) {
	a.ListGlobalAddressesCall.CallCount++

	return a.ListGlobalAddressesCall.Returns.Output, a.ListGlobalAddressesCall.Returns.Error
}

func (a *GlobalAddressesClient) DeleteGlobalAddress(address string) error {
	a.DeleteGlobalAddressCall.CallCount++
	a.DeleteGlobalAddressCall.Receives.Address = address

	return a.DeleteGlobalAddressCall.Returns.Error
}
