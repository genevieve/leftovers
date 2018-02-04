package compute

import (
	"fmt"
	"strings"

	gcp "google.golang.org/api/compute/v1"
)

type networksClient interface {
	ListNetworks() (*gcp.NetworkList, error)
	DeleteNetwork(network string) error
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

func (n Networks) List(filter string) (map[string]string, error) {
	networks, err := n.client.ListNetworks()
	if err != nil {
		return nil, fmt.Errorf("Listing networks: %s", err)
	}

	delete := map[string]string{}
	for _, network := range networks.Items {
		if !strings.Contains(network.Name, filter) {
			continue
		}

		proceed := n.logger.Prompt(fmt.Sprintf("Are you sure you want to delete network %s?", network.Name))
		if !proceed {
			continue
		}

		delete[network.Name] = ""
	}

	return delete, nil
}

func (n Networks) Delete(networks map[string]string) {
	var resources []Network
	for name, _ := range networks {
		resources = append(resources, NewNetwork(n.client, name))
	}

	n.delete(resources)
}

func (n Networks) delete(resources []Network) {
	for _, resource := range resources {
		err := resource.Delete()

		if err != nil {
			n.logger.Printf("%s\n", err)
		} else {
			n.logger.Printf("SUCCESS deleting network %s\n", resource.name)
		}
	}
}
