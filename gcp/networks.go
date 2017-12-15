package gcp

import (
	"fmt"

	compute "google.golang.org/api/compute/v1"
)

type networksClient interface {
	List(project string) (*compute.NetworkList, error)
	Delete(project, network string) (*compute.Operation, error)
}

type Networks struct {
	client  networksClient
	logger  logger
	project string
}

func NewNetworks(client networksClient, logger logger, project string) Networks {
	return Networks{
		client:  client,
		logger:  logger,
		project: project,
	}
}

func (e Networks) Delete() error {
	networks, err := e.client.List(e.project)
	if err != nil {
		return fmt.Errorf("Listing networks: %s", err)
	}

	for _, t := range networks.Items {
		n := t.Name

		proceed := e.logger.Prompt(fmt.Sprintf("Are you sure you want to delete network %s?", n))
		if !proceed {
			continue
		}

		_, err = e.client.Delete(e.project, n)
		if err == nil {
			e.logger.Printf("SUCCESS deleting network %s\n", n)
		} else {
			e.logger.Printf("ERROR deleting network %s: %s\n", n, err)
		}
	}

	return nil
}
