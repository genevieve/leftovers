package compute

import (
	"fmt"
	"strings"

	gcpcompute "google.golang.org/api/compute/v1"
)

type addressesClient interface {
	ListAddresses(region string) (*gcpcompute.AddressList, error)
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

func (o Addresses) Delete(filter string) error {
	var addrs []*gcpcompute.Address
	for _, region := range o.regions {
		l, err := o.client.ListAddresses(region)
		if err != nil {
			return fmt.Errorf("Listing addresses for region %s: %s", region, err)
		}
		addrs = append(addrs, l.Items...)
	}

	for _, a := range addrs {
		if len(a.Users) > 0 {
			continue
		}

		n := a.Name

		if !strings.Contains(n, filter) {
			continue
		}

		proceed := o.logger.Prompt(fmt.Sprintf("Are you sure you want to delete address %s?", n))
		if !proceed {
			continue
		}

		regionName := o.regions[a.Region]
		if err := o.client.DeleteAddress(regionName, n); err != nil {
			o.logger.Printf("ERROR deleting address %s: %s\n", n, err)
		} else {
			o.logger.Printf("SUCCESS deleting address %s\n", n)
		}
	}

	return nil
}
