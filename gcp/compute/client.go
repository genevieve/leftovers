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

func (c client) ListZones() ([]string, error) {
	list, err := c.zones.List(c.project).Do()
	if err != nil {
		return []string{}, err
	}

	var zones []string
	for _, z := range list.Items {
		zones = append(zones, z.Name)
	}
	return zones, nil
}

func (c client) ListNetworks() (*gcp.NetworkList, error) {
	return c.networks.List(c.project).Do()
}

func (c client) DeleteNetwork(network string) (*gcp.Operation, error) {
	return c.networks.Delete(c.project, network).Do()
}

func (c client) ListDisks() (*gcp.DiskList, error) {
	return c.disks.List(c.project, "us-west1-a").Do()
}

func (c client) DeleteDisk(zone, disk string) (*gcp.Operation, error) {
	return c.disks.Delete(c.project, "us-west1-a", disk).Do()
}
