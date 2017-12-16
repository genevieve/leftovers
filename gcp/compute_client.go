package gcp

import (
	compute "google.golang.org/api/compute/v1"
)

type computeClient struct {
	project  string
	networks *compute.NetworksService
	disks    *compute.DisksService
	zones    *compute.ZonesService
}

func NewComputeClient(project string, service *compute.Service) computeClient {
	return computeClient{
		project:  project,
		networks: service.Networks,
		disks:    service.Disks,
		zones:    service.Zones,
	}
}

func (c computeClient) ListZones() ([]string, error) {
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

func (c computeClient) ListNetworks() (*compute.NetworkList, error) {
	return c.networks.List(c.project).Do()
}

func (c computeClient) DeleteNetwork(network string) (*compute.Operation, error) {
	return c.networks.Delete(c.project, network).Do()
}

func (c computeClient) ListDisks() (*compute.DiskList, error) {
	return c.disks.List(c.project, "us-west1-a").Do()
}

func (c computeClient) DeleteDisk(zone, disk string) (*compute.Operation, error) {
	return c.disks.Delete(c.project, "us-west1-a", disk).Do()
}
