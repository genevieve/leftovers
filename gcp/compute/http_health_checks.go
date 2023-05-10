package compute

import (
	"fmt"
	"github.com/genevieve/leftovers/common"
	gcpcompute "google.golang.org/api/compute/v1"
)

//go:generate faux --interface httpHealthChecksClient --output fakes/http_health_checks_client.go
type httpHealthChecksClient interface {
	ListHttpHealthChecks() ([]*gcpcompute.HttpHealthCheck, error)
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

func (h HttpHealthChecks) List(filter string, regex bool) ([]common.Deletable, error) {
	h.logger.Debugln("Listing Http Health Checks...")
	checks, err := h.client.ListHttpHealthChecks()
	if err != nil {
		return nil, fmt.Errorf("List Http Health Checks: %s", err)
	}

	var resources []common.Deletable
	for _, check := range checks {
		resource := NewHttpHealthCheck(h.client, check.Name)

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

func (h HttpHealthChecks) Type() string {
	return "http-health-check"
}
