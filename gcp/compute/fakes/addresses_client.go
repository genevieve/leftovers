package fakes

import gcpcompute "google.golang.org/api/compute/v1"

type AddressesClient struct {
	ListAddressesCall struct {
		CallCount int
		Receives  struct {
			Region string
		}
		Returns struct {
			Output []*gcpcompute.Address
			Error  error
		}
	}

	DeleteAddressCall struct {
		CallCount int
		Receives  struct {
			Address string
			Region  string
		}
		Returns struct {
			Error error
		}
	}
}

func (a *AddressesClient) ListAddresses(region string) ([]*gcpcompute.Address, error) {
	a.ListAddressesCall.CallCount++
	a.ListAddressesCall.Receives.Region = region

	return a.ListAddressesCall.Returns.Output, a.ListAddressesCall.Returns.Error
}

func (a *AddressesClient) DeleteAddress(region, address string) error {
	a.DeleteAddressCall.CallCount++
	a.DeleteAddressCall.Receives.Address = address
	a.DeleteAddressCall.Receives.Region = region

	return a.DeleteAddressCall.Returns.Error
}
