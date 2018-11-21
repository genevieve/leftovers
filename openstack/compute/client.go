package compute

import (
	"github.com/gophercloud/gophercloud"

	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/floatingips"
)

type client struct {
	project string
	logger  logger

	service *gophercloud.ServiceClient
}

func NewClient(project string, service *gophercloud.ServiceClient, logger logger) client {
	return client{
		project: project,
		logger:  logger,
		service: service,
	}
}

func (c client) ListFloatingIPs() ([]floatingips.FloatingIP, error) {
	page, err := floatingips.List(c.service).AllPages()
	if err != nil {
		return nil, err
	}

	ips, err := floatingips.ExtractFloatingIPs(page)
	if err != nil {
		return nil, err
	}

	return ips, nil
}

func (c client) DeleteFloatingIP(ip string) error {
	return floatingips.Delete(c.service, ip).ExtractErr()
}
