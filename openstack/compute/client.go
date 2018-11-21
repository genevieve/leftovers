package compute

import (
	"github.com/gophercloud/gophercloud"

	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/floatingips"
)

type client struct {
	logger  logger
	service *gophercloud.ServiceClient
}

func NewClient(service *gophercloud.ServiceClient, logger logger) client {
	return client{
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
