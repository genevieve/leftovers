package compute

import (
	gcpcompute "google.golang.org/api/compute/v1"
)

type client struct {
	project string

	disks     *gcpcompute.DisksService
	instances *gcpcompute.InstancesService
	networks  *gcpcompute.NetworksService
	zones     *gcpcompute.ZonesService
}

func NewClient(project string, service *gcpcompute.Service) client {
	return client{
		project:   project,
		disks:     service.Disks,
		instances: service.Instances,
		networks:  service.Networks,
		zones:     service.Zones,
	}
}

func (c client) ListDisks(zone string) (*gcpcompute.DiskList, error) {
	return c.disks.List(c.project, zone).Do()
}

func (c client) DeleteDisk(zone, disk string) (*gcpcompute.Operation, error) {
	return c.disks.Delete(c.project, zone, disk).Do()
}

func (c client) ListInstances(zone string) (*gcpcompute.InstanceList, error) {
	return c.instances.List(c.project, zone).Do()
}

func (c client) DeleteInstance(zone, instance string) (*gcpcompute.Operation, error) {
	return c.instances.Delete(c.project, zone, instance).Do()
}

func (c client) ListNetworks() (*gcpcompute.NetworkList, error) {
	return c.networks.List(c.project).Do()
}

func (c client) DeleteNetwork(network string) (*gcpcompute.Operation, error) {
	return c.networks.Delete(c.project, network).Do()
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
