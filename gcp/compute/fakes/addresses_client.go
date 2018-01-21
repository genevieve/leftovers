package fakes

import gcpcompute "google.golang.org/api/compute/v1"

type AddressesClient struct {
	ListAddressesCall struct {
		CallCount int
		Receives  struct {
			Region string
			Filter string
		}
		Returns struct {
			Output *gcpcompute.AddressList
			Error  error
		}
	}

	DeleteAddressCall struct {
		CallCount int
		Receives  struct {
			Region  string
			Address string
		}
		Returns struct {
			Error error
		}
	}
}

func (n *AddressesClient) ListAddresses(region, filter string) (*gcpcompute.AddressList, error) {
	n.ListAddressesCall.CallCount++
	n.ListAddressesCall.Receives.Region = region
	n.ListAddressesCall.Receives.Filter = filter

	return n.ListAddressesCall.Returns.Output, n.ListAddressesCall.Returns.Error
}

func (n *AddressesClient) DeleteAddress(region, address string) error {
	n.DeleteAddressCall.CallCount++
	n.DeleteAddressCall.Receives.Address = address
	n.DeleteAddressCall.Receives.Region = region

	return n.DeleteAddressCall.Returns.Error
}
