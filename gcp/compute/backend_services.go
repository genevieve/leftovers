package compute

import (
	"fmt"
	"github.com/genevieve/leftovers/common"
	gcpcompute "google.golang.org/api/compute/v1"
)

//go:generate faux --interface backendServicesClient --output fakes/backend_services_client.go
type backendServicesClient interface {
	ListBackendServices() ([]*gcpcompute.BackendService, error)
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

func (b BackendServices) List(filter string, regex bool) ([]common.Deletable, error) {
	b.logger.Debugln("Listing Backend Services...")
	backendServices, err := b.client.ListBackendServices()
	if err != nil {
		return nil, fmt.Errorf("List Backend Services: %s", err)
	}

	var resources []common.Deletable
	for _, backend := range backendServices {
		resource := NewBackendService(b.client, backend.Name)

		if !common.ResourceMatches(backend.Name, filter, regex) {
			continue
		}

		proceed := b.logger.PromptWithDetails(resource.Type(), resource.Name())
		if !proceed {
			continue
		}

		resources = append(resources, resource)
	}

	return resources, nil
}

func (b BackendServices) Type() string {
	return "backend-service"
}
