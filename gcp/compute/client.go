package compute

import (
	gcp "google.golang.org/api/compute/v1"
)

type client struct {
	project  string
	networks *gcp.NetworksService
	disks    *gcp.DisksService
	zones    *gcp.ZonesService
}

func NewClient(project string, service *gcp.Service) client {
	return client{
		project:  project,
		networks: service.Networks,
		disks:    service.Disks,
		zones:    service.Zones,
	}
}

func (c client) ListZones() (map[string]string, error) {
	zones := map[string]string{}

	list, err := c.zones.List(c.project).Do()
	if err != nil {
		return zones, err
	}

	for _, z := range list.Items {
		zones[z.SelfLink] = z.Name
	}
	return zones, nil
}

func (c client) ListNetworks() (*gcp.NetworkList, error) {
	return c.networks.List(c.project).Do()
}

func (c client) DeleteNetwork(network string) (*gcp.Operation, error) {
	return c.networks.Delete(c.project, network).Do()
}

func (c client) ListDisks(zone string) (*gcp.DiskList, error) {
	return c.disks.List(c.project, zone).Do()
}

func (c client) DeleteDisk(zone, disk string) (*gcp.Operation, error) {
	return c.disks.Delete(c.project, zone, disk).Do()
}
