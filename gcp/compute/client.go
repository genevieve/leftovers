package compute

import (
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	gcpcompute "google.golang.org/api/compute/v1"
)

type client struct {
	project string

	service           *gcpcompute.Service
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
		service:           service,
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

func (c client) DeleteBackendService(backendService string) error {
	op, err := c.backendServices.Delete(c.project, backendService).Do()
	if err != nil {
		return err
	}

	return c.waitOnDelete(op)
}

func (c client) ListDisks(zone string) (*gcpcompute.DiskList, error) {
	return c.disks.List(c.project, zone).Do()
}

func (c client) DeleteDisk(zone, disk string) error {
	op, err := c.disks.Delete(c.project, zone, disk).Do()
	if err != nil {
		return err
	}

	return c.waitOnDelete(op)
}

func (c client) ListInstances(zone string) (*gcpcompute.InstanceList, error) {
	return c.instances.List(c.project, zone).Do()
}

func (c client) DeleteInstance(zone, instance string) error {
	op, err := c.instances.Delete(c.project, zone, instance).Do()
	if err != nil {
		return err
	}

	return c.waitOnDelete(op)
}

func (c client) ListNetworks() (*gcpcompute.NetworkList, error) {
	return c.networks.List(c.project).Do()
}

func (c client) DeleteNetwork(network string) error {
	op, err := c.networks.Delete(c.project, network).Do()
	if err != nil {
		return err
	}

	return c.waitOnDelete(op)
}

func (c client) ListHttpHealthChecks() (*gcpcompute.HttpHealthCheckList, error) {
	return c.httpHealthChecks.List(c.project).Do()
}

func (c client) DeleteHttpHealthCheck(httpHealthCheck string) error {
	op, err := c.httpHealthChecks.Delete(c.project, httpHealthCheck).Do()
	if err != nil {
		return err
	}

	return c.waitOnDelete(op)
}

func (c client) ListHttpsHealthChecks() (*gcpcompute.HttpsHealthCheckList, error) {
	return c.httpsHealthChecks.List(c.project).Do()
}

func (c client) DeleteHttpsHealthCheck(httpsHealthCheck string) error {
	op, err := c.httpsHealthChecks.Delete(c.project, httpsHealthCheck).Do()
	if err != nil {
		return err
	}

	return c.waitOnDelete(op)
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

func (c client) waitOnDelete(op *gcpcompute.Operation) error {
	waiter := &operationWaiter{
		Op:      op,
		Service: c.service,
		Project: c.project,
	}

	state := &resource.StateChangeConf{
		Delay:      10 * time.Second,
		Timeout:    10 * time.Minute,
		MinTimeout: 2 * time.Second,
		Pending:    []string{"PENDING", "RUNNING"},
		Target:     []string{"DONE"},
		Refresh:    waiter.refreshFunc(),
	}

	opRaw, err := state.WaitForState()
	if err != nil {
		return err
	}

	if resultOp := opRaw.(*gcpcompute.Operation); resultOp.Error != nil {
		return ComputeOperationError(*resultOp.Error)
	}

	return nil
}
