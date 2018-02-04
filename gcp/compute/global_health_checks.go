package compute

import (
	"fmt"
	"strings"

	gcpcompute "google.golang.org/api/compute/v1"
)

type globalHealthChecksClient interface {
	ListGlobalHealthChecks() (*gcpcompute.HealthCheckList, error)
	DeleteGlobalHealthCheck(globalHealthCheck string) error
}

type GlobalHealthChecks struct {
	client globalHealthChecksClient
	logger logger
}

func NewGlobalHealthChecks(client globalHealthChecksClient, logger logger) GlobalHealthChecks {
	return GlobalHealthChecks{
		client: client,
		logger: logger,
	}
}

func (h GlobalHealthChecks) List(filter string) (map[string]string, error) {
	checks, err := h.client.ListGlobalHealthChecks()
	if err != nil {
		return nil, fmt.Errorf("Listing global health checks: %s", err)
	}

	delete := map[string]string{}
	for _, check := range checks.Items {
		if !strings.Contains(check.Name, filter) {
			continue
		}

		proceed := h.logger.Prompt(fmt.Sprintf("Are you sure you want to delete global health check %s?", check.Name))
		if !proceed {
			continue
		}

		delete[check.Name] = ""
	}

	return delete, nil
}

func (h GlobalHealthChecks) Delete(checks map[string]string) {
	var resources []GlobalHealthCheck
	for name, _ := range checks {
		resources = append(resources, NewGlobalHealthCheck(h.client, name))
	}

	h.delete(resources)
}

func (g GlobalHealthChecks) delete(resources []GlobalHealthCheck) {
	for _, resource := range resources {
		err := resource.Delete()

		if err != nil {
			g.logger.Printf("%s\n", err)
		} else {
			g.logger.Printf("SUCCESS deleting global health check %s\n", resource.name)
		}
	}
}
