package compute

import (
	"fmt"
	"github.com/genevieve/leftovers/common"
	gcpcompute "google.golang.org/api/compute/v1"
)

//go:generate faux --interface addressesClient --output fakes/addresses_client.go
type addressesClient interface {
	ListAddresses(region string) ([]*gcpcompute.Address, error)
	DeleteAddress(region, address string) error
}

type Addresses struct {
	client  addressesClient
	logger  logger
	regions map[string]string
}

func NewAddresses(client addressesClient, logger logger, regions map[string]string) Addresses {
	return Addresses{
		client:  client,
		logger:  logger,
		regions: regions,
	}
}

func (a Addresses) List(filter string, regex bool) ([]common.Deletable, error) {
	addresses := []*gcpcompute.Address{}
	for _, region := range a.regions {
		a.logger.Debugf("Listing Addresses for Region %s...\n", region)
		l, err := a.client.ListAddresses(region)
		if err != nil {
			return nil, fmt.Errorf("List Addresses for Region %s: %s", region, err)
		}

		addresses = append(addresses, l...)
	}

	var resources []common.Deletable
	for _, address := range addresses {
		resource := NewAddress(a.client, address.Name, a.regions[address.Region], len(address.Users))

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

func (a Addresses) Type() string {
	return "address"
}
