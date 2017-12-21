package compute

import (
	"fmt"

	gcpcompute "google.golang.org/api/compute/v1"
)

type httpsHealthChecksClient interface {
	ListHttpsHealthChecks() (*gcpcompute.HttpsHealthCheckList, error)
	DeleteHttpsHealthCheck(httpsHealthCheck string) (*gcpcompute.Operation, error)
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

func (i HttpsHealthChecks) Delete() error {
	httpsHealthChecks, err := i.client.ListHttpsHealthChecks()
	if err != nil {
		return fmt.Errorf("Listing https health checks: %s", err)
	}

	for _, h := range httpsHealthChecks.Items {
		proceed := i.logger.Prompt(fmt.Sprintf("Are you sure you want to delete https health check %s?", h.Name))
		if !proceed {
			continue
		}

		if _, err := i.client.DeleteHttpsHealthCheck(h.Name); err != nil {
			i.logger.Printf("ERROR deleting https health check %s: %s\n", h.Name, err)
		} else {
			i.logger.Printf("SUCCESS deleting https health check %s\n", h.Name)
		}
	}

	return nil
}
