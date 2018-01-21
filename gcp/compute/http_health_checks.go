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

func (i HttpHealthChecks) Delete(filter string) error {
	httpHealthChecks, err := i.client.ListHttpHealthChecks()
	if err != nil {
		return fmt.Errorf("Listing http health checks: %s", err)
	}

	for _, check := range httpHealthChecks.Items {
		n := check.Name

		if !strings.Contains(n, filter) {
			continue
		}

		proceed := i.logger.Prompt(fmt.Sprintf("Are you sure you want to delete http health check %s?", n))
		if !proceed {
			continue
		}

		if err := i.client.DeleteHttpHealthCheck(n); err != nil {
			i.logger.Printf("ERROR deleting http health check %s: %s\n", n, err)
		} else {
			i.logger.Printf("SUCCESS deleting http health check %s\n", n)
		}
	}

	return nil
}
