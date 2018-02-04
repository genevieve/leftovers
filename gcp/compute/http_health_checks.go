package compute

import (
	"fmt"
	"strings"

	gcpcompute "google.golang.org/api/compute/v1"
)

type httpHealthChecksClient interface {
	ListHttpHealthChecks() (*gcpcompute.HttpHealthCheckList, error)
	DeleteHttpHealthCheck(httpHealthCheck string) error
}

type HttpHealthChecks struct {
	client httpHealthChecksClient
	logger logger
}

func NewHttpHealthChecks(client httpHealthChecksClient, logger logger) HttpHealthChecks {
	return HttpHealthChecks{
		client: client,
		logger: logger,
	}
}

func (h HttpHealthChecks) List(filter string) (map[string]string, error) {
	checks, err := h.client.ListHttpHealthChecks()
	if err != nil {
		return nil, fmt.Errorf("Listing http health checks: %s", err)
	}

	delete := map[string]string{}
	for _, check := range checks.Items {
		if !strings.Contains(check.Name, filter) {
			continue
		}

		proceed := h.logger.Prompt(fmt.Sprintf("Are you sure you want to delete http health check %s?", check.Name))
		if !proceed {
			continue
		}

		delete[check.Name] = ""
	}

	return delete, nil
}

func (h HttpHealthChecks) Delete(checks map[string]string) {
	var resources []HttpHealthCheck
	for name, _ := range checks {
		resources = append(resources, NewHttpHealthCheck(h.client, name))
	}

	h.delete(resources)
}

func (h HttpHealthChecks) delete(resources []HttpHealthCheck) {
	for _, resource := range resources {
		err := resource.Delete()

		if err != nil {
			h.logger.Printf("%s\n", err)
		} else {
			h.logger.Printf("SUCCESS deleting http health check %s\n", resource.name)
		}
	}
}
