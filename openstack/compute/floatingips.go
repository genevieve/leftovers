package compute

import (
	"fmt"

	"github.com/genevieve/leftovers/common"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/floatingips"
)

type floatingipsClient interface {
	ListFloatingIPs() ([]floatingips.FloatingIP, error)
	DeleteFloatingIP(ip string) error
}

type FloatingIPs struct {
	client floatingipsClient
	logger logger
}

func NewFloatingIPs(client floatingipsClient, logger logger) FloatingIPs {
	return FloatingIPs{
		client: client,
		logger: logger,
	}
}

func (f FloatingIPs) List(filter string) ([]common.Deletable, error) {
	ips, err := f.client.ListFloatingIPs()
	if err != nil {
		return nil, fmt.Errorf("List Floating IPs: %s", err)
	}

	var resources []common.Deletable
	for _, i := range ips {
		resource := NewFloatingIP(f.client, i.IP)

		proceed := f.logger.PromptWithDetails(resource.Type(), resource.Name())
		if !proceed {
			continue
		}

		resources = append(resources, resource)
	}

	return resources, nil
}

func (a FloatingIPs) Type() string {
	return "floatingip"
}
