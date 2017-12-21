package compute

import (
	gcpcompute "google.golang.org/api/compute/v1"
)

type client struct {
	project string

	backendServices   *gcpcompute.BackendServicesService
	disks             *gcpcompute.DisksService
	httpHealthChecks  *gcpcompute.HttpHealthChecksService
	httpsHealthChecks *gcpcompute.HttpsHealthChecksService
	instances         *gcpcompute.InstancesService
	networks          *gcpcompute.NetworksService
	zones             *gcpcompute.ZonesService
}

func NewClient(project string, service *gcpcompute.Service) client {
	return client{
		project:           project,
		backendServices:   service.BackendServices,
		disks:             service.Disks,
		httpHealthChecks:  service.HttpHealthChecks,
		httpsHealthChecks: service.HttpsHealthChecks,
		instances:         service.Instances,
		networks:          service.Networks,
		zones:             service.Zones,
	}
}

func (c client) ListBackendServices() (*gcpcompute.BackendServiceList, error) {
	return c.backendServices.List(c.project).Do()
}

func (c client) DeleteBackendService(backendService string) (*gcpcompute.Operation, error) {
	return c.backendServices.Delete(c.project, backendService).Do()
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

func (c client) ListHttpHealthChecks() (*gcpcompute.HttpHealthCheckList, error) {
	return c.httpHealthChecks.List(c.project).Do()
}

func (c client) DeleteHttpHealthCheck(httpHealthCheck string) (*gcpcompute.Operation, error) {
	return c.httpHealthChecks.Delete(c.project, httpHealthCheck).Do()
}

func (c client) ListHttpsHealthChecks() (*gcpcompute.HttpsHealthCheckList, error) {
	return c.httpsHealthChecks.List(c.project).Do()
}

func (c client) DeleteHttpsHealthCheck(httpsHealthCheck string) (*gcpcompute.Operation, error) {
	return c.httpsHealthChecks.Delete(c.project, httpsHealthCheck).Do()
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
