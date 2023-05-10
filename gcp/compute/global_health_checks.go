package compute

import (
	"fmt"
	"github.com/genevieve/leftovers/common"
	gcpcompute "google.golang.org/api/compute/v1"
)

//go:generate faux --interface globalHealthChecksClient --output fakes/global_health_checks_client.go
type globalHealthChecksClient interface {
	ListGlobalHealthChecks() ([]*gcpcompute.HealthCheck, error)
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

func (h GlobalHealthChecks) List(filter string, regex bool) ([]common.Deletable, error) {
	h.logger.Debugln("Listing Global Health Checks...")
	checks, err := h.client.ListGlobalHealthChecks()
	if err != nil {
		return nil, fmt.Errorf("List Global Health Checks: %s", err)
	}

	var resources []common.Deletable
	for _, check := range checks {
		resource := NewGlobalHealthCheck(h.client, check.Name)

		if !common.ResourceMatches(check.Name, filter, regex) {
			continue
		}

		proceed := h.logger.PromptWithDetails(resource.Type(), resource.Name())
		if !proceed {
			continue
		}

		resources = append(resources, resource)
	}

	return resources, nil
}

func (h GlobalHealthChecks) Type() string {
	return "global-health-check"
}
