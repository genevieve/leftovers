package compute

import (
	"fmt"
	"strings"

	"github.com/genevieve/leftovers/common"
	gcpcompute "google.golang.org/api/compute/v1"
)

//go:generate faux --interface firewallsClient --output fakes/firewalls_client.go
type firewallsClient interface {
	GetNetworkName(url string) (name string)
	ListFirewalls() ([]*gcpcompute.Firewall, error)
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

func (f Firewalls) List(filter string) ([]common.Deletable, error) {
	f.logger.Debugln("Listing Firewalls...")
	firewalls, err := f.client.ListFirewalls()
	if err != nil {
		return nil, fmt.Errorf("Listing Firewalls: %s", err)
	}

	var resources []common.Deletable
	for _, firewall := range firewalls {
		resource := NewFirewall(f.client, firewall.Name, firewall.Network)

		if strings.Contains(resource.Name(), "default") {
			continue
		}

		if !strings.Contains(resource.Name(), filter) {
			continue
		}

		proceed := f.logger.PromptWithDetails(resource.Type(), resource.Name())
		if !proceed {
			continue
		}

		resources = append(resources, resource)
	}

	return resources, nil
}

func (f Firewalls) Type() string {
	return "firewall"
}
