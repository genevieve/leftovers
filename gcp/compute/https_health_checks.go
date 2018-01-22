package compute

import (
	"fmt"
	"strings"

	gcpcompute "google.golang.org/api/compute/v1"
)

type httpsHealthChecksClient interface {
	ListHttpsHealthChecks() (*gcpcompute.HttpsHealthCheckList, error)
	DeleteHttpsHealthCheck(httpsHealthCheck string) error
}

type HttpsHealthChecks struct {
	client httpsHealthChecksClient
	logger logger
}

func NewHttpsHealthChecks(client httpsHealthChecksClient, logger logger) HttpsHealthChecks {
	return HttpsHealthChecks{
		client: client,
		logger: logger,
	}
}

func (h HttpsHealthChecks) List(filter string) (map[string]string, error) {
	delete := map[string]string{}

	checks, err := h.client.ListHttpsHealthChecks()
	if err != nil {
		return delete, fmt.Errorf("Listing https health checks: %s", err)
	}

	for _, check := range checks.Items {
		if !strings.Contains(check.Name, filter) {
			continue
		}

		proceed := h.logger.Prompt(fmt.Sprintf("Are you sure you want to delete https health check %s?", check.Name))
		if !proceed {
			continue
		}

		delete[check.Name] = ""
	}

	return delete, nil
}

func (h HttpsHealthChecks) Delete(checks map[string]string) {
	for name, _ := range checks {
		err := h.client.DeleteHttpsHealthCheck(name)

		if err != nil {
			h.logger.Printf("ERROR deleting https health check %s: %s\n", name, err)
		} else {
			h.logger.Printf("SUCCESS deleting https health check %s\n", name)
		}
	}
}
