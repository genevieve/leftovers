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

func (i HttpsHealthChecks) Delete(filter string) error {
	httpsHealthChecks, err := i.client.ListHttpsHealthChecks()
	if err != nil {
		return fmt.Errorf("Listing https health checks: %s", err)
	}

	for _, check := range httpsHealthChecks.Items {
		n := check.Name

		if !strings.Contains(n, filter) {
			continue
		}

		proceed := i.logger.Prompt(fmt.Sprintf("Are you sure you want to delete https health check %s?", n))
		if !proceed {
			continue
		}

		if err := i.client.DeleteHttpsHealthCheck(n); err != nil {
			i.logger.Printf("ERROR deleting https health check %s: %s\n", n, err)
		} else {
			i.logger.Printf("SUCCESS deleting https health check %s\n", n)
		}
	}

	return nil
}
