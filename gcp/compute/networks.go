package compute

import (
	"fmt"

	gcp "google.golang.org/api/compute/v1"
)

type networksClient interface {
	ListNetworks() (*gcp.NetworkList, error)
	DeleteNetwork(network string) (*gcp.Operation, error)
}

type Networks struct {
	client networksClient
	logger logger
}

func NewNetworks(client networksClient, logger logger) Networks {
	return Networks{
		client: client,
		logger: logger,
	}
}

func (e Networks) Delete() error {
	networks, err := e.client.ListNetworks()
	if err != nil {
		return fmt.Errorf("Listing networks: %s", err)
	}

	for _, t := range networks.Items {
		n := t.Name

		proceed := e.logger.Prompt(fmt.Sprintf("Are you sure you want to delete network %s?", n))
		if !proceed {
			continue
		}

		_, err = e.client.DeleteNetwork(n)
		if err == nil {
			e.logger.Printf("SUCCESS deleting network %s\n", n)
		} else {
			e.logger.Printf("ERROR deleting network %s: %s\n", n, err)
		}
	}

	return nil
}
