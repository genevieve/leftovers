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

func (b BackendServices) List(filter string) (map[string]string, error) {
	backendServices, err := b.client.ListBackendServices()
	if err != nil {
		return nil, fmt.Errorf("Listing backend services: %s", err)
	}

	delete := map[string]string{}
	for _, backend := range backendServices.Items {
		if !strings.Contains(backend.Name, filter) {
			continue
		}

		proceed := b.logger.Prompt(fmt.Sprintf("Are you sure you want to delete backend service %s?", backend.Name))
		if !proceed {
			continue
		}

		delete[backend.Name] = ""
	}

	return delete, nil
}

func (b BackendServices) delete(resources []BackendService) {
	for _, resource := range resources {
		err := resource.Delete()

		if err != nil {
			b.logger.Printf("ERROR deleting backend service %s: %s\n", resource.name, err)
		} else {
			b.logger.Printf("SUCCESS deleting backend service %s\n", resource.name)
		}
	}
}

func (b BackendServices) Delete(backendServices map[string]string) {
	var resources []BackendService
	for name, _ := range backendServices {
		resources = append(resources, NewBackendService(b.client, name))
	}

	b.delete(resources)
}
