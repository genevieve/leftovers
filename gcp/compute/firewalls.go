package compute

import (
	"fmt"
	"strings"

	gcpcompute "google.golang.org/api/compute/v1"
)

type firewallsClient interface {
	ListFirewalls() (*gcpcompute.FirewallList, error)
	DeleteFirewall(firewall string) error
}

type Firewalls struct {
	client firewallsClient
	logger logger
}

func NewFirewalls(client firewallsClient, logger logger) Firewalls {
	return Firewalls{
		client: client,
		logger: logger,
	}
}

func (i Firewalls) Delete(filter string) error {
	firewalls, err := i.client.ListFirewalls()
	if err != nil {
		return fmt.Errorf("Listing firewalls: %s", err)
	}

	for _, f := range firewalls.Items {
		n := f.Name

		if !strings.Contains(n, filter) {
			continue
		}

		proceed := i.logger.Prompt(fmt.Sprintf("Are you sure you want to delete firewall %s?", n))
		if !proceed {
			continue
		}

		if err := i.client.DeleteFirewall(n); err != nil {
			i.logger.Printf("ERROR deleting firewall %s: %s\n", n, err)
		} else {
			i.logger.Printf("SUCCESS deleting firewall %s\n", n)
		}
	}

	return nil
}
