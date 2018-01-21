package compute

import (
	"fmt"
	"strings"

	gcpcompute "google.golang.org/api/compute/v1"
)

type backendServicesClient interface {
	ListBackendServices() (*gcpcompute.BackendServiceList, error)
	DeleteBackendService(backendService string) error
}

type BackendServices struct {
	client backendServicesClient
	logger logger
}

func NewBackendServices(client backendServicesClient, logger logger) BackendServices {
	return BackendServices{
		client: client,
		logger: logger,
	}
}

func (i BackendServices) Delete(filter string) error {
	backendServices, err := i.client.ListBackendServices()
	if err != nil {
		return fmt.Errorf("Listing backend services: %s", err)
	}

	for _, b := range backendServices.Items {
		n := b.Name

		if !strings.Contains(n, filter) {
			continue
		}

		proceed := i.logger.Prompt(fmt.Sprintf("Are you sure you want to delete backend service %s?", b.Name))
		if !proceed {
			continue
		}

		if err := i.client.DeleteBackendService(b.Name); err != nil {
			i.logger.Printf("ERROR deleting backend service %s: %s\n", b.Name, err)
		} else {
			i.logger.Printf("SUCCESS deleting backend service %s\n", b.Name)
		}
	}

	return nil
}
