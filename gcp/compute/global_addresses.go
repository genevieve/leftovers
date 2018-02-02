package compute

import (
	"fmt"
	"strings"
	"sync"

	gcpcompute "google.golang.org/api/compute/v1"
)

type globalAddressesClient interface {
	ListGlobalAddresses() (*gcpcompute.AddressList, error)
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

func (a GlobalAddresses) List(filter string) (map[string]string, error) {
	addresses, err := a.client.ListGlobalAddresses()
	if err != nil {
		return nil, fmt.Errorf("Listing global addresses: %s", err)
	}

	delete := map[string]string{}
	for _, address := range addresses.Items {
		if len(address.Users) > 0 {
			continue
		}

		if !strings.Contains(address.Name, filter) {
			continue
		}

		proceed := a.logger.Prompt(fmt.Sprintf("Are you sure you want to delete global address %s?", address.Name))
		if !proceed {
			continue
		}

		delete[address.Name] = ""
	}

	return delete, nil
}

func (a GlobalAddresses) Delete(addrs map[string]string) {
	var wg sync.WaitGroup

	for name, _ := range addrs {
		wg.Add(1)

		go func(name string) {
			err := a.client.DeleteGlobalAddress(name)

			if err != nil {
				a.logger.Printf("ERROR deleting global address %s: %s\n", name, err)
			} else {
				a.logger.Printf("SUCCESS deleting global address %s\n", name)
			}

			wg.Done()
		}(name)
	}

	wg.Wait()
}
