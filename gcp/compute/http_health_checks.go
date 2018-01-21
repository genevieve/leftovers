package compute

import (
	"fmt"

	gcpcompute "google.golang.org/api/compute/v1"
)

type httpHealthChecksClient interface {
	ListHttpHealthChecks(filter string) (*gcpcompute.HttpHealthCheckList, error)
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
	httpHealthChecks, err := i.client.ListHttpHealthChecks(filter)
	if err != nil {
		return fmt.Errorf("Listing http health checks: %s", err)
	}

	for _, h := range httpHealthChecks.Items {
		proceed := i.logger.Prompt(fmt.Sprintf("Are you sure you want to delete http health check %s?", h.Name))
		if !proceed {
			continue
		}

		if err := i.client.DeleteHttpHealthCheck(h.Name); err != nil {
			i.logger.Printf("ERROR deleting http health check %s: %s\n", h.Name, err)
		} else {
			i.logger.Printf("SUCCESS deleting http health check %s\n", h.Name)
		}
	}

	return nil
}
