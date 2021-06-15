package compute

import (
	"fmt"
	"github.com/genevieve/leftovers/common"
	gcpcompute "google.golang.org/api/compute/v1"
)

//go:generate faux --interface globalAddressesClient --output fakes/global_address_client.go
type globalAddressesClient interface {
	ListGlobalAddresses() ([]*gcpcompute.Address, error)
	DeleteGlobalAddress(address string) error
}

type GlobalAddresses struct {
	client globalAddressesClient
	logger logger
}

func NewGlobalAddresses(client globalAddressesClient, logger logger) GlobalAddresses {
	return GlobalAddresses{
		client: client,
		logger: logger,
	}
}

func (a GlobalAddresses) List(filter string, regex bool) ([]common.Deletable, error) {
	a.logger.Debugln("Listing Global Addresses...")
	addresses, err := a.client.ListGlobalAddresses()
	if err != nil {
		return nil, fmt.Errorf("List Global Addresses: %s", err)
	}

	var resources []common.Deletable
	for _, address := range addresses {
		resource := NewGlobalAddress(a.client, address.Name)

		if !common.ResourceMatches(address.Name, filter, regex) {
			continue
		}

		proceed := a.logger.PromptWithDetails(resource.Type(), resource.Name())
		if !proceed {
			continue
		}

		resources = append(resources, resource)
	}

	return resources, nil
}

func (a GlobalAddresses) Type() string {
	return "global-address"
}
