package compute

import (
	"fmt"

	gcpcompute "google.golang.org/api/compute/v1"
)

type client struct {
	project string
	logger  logger

	service           *gcpcompute.Service
	addresses         *gcpcompute.AddressesService
	backendServices   *gcpcompute.BackendServicesService
	disks             *gcpcompute.DisksService
	httpHealthChecks  *gcpcompute.HttpHealthChecksService
	httpsHealthChecks *gcpcompute.HttpsHealthChecksService
	instances         *gcpcompute.InstancesService
	instanceGroups    *gcpcompute.InstanceGroupsService
	firewalls         *gcpcompute.FirewallsService
	forwardingRules   *gcpcompute.ForwardingRulesService
	networks          *gcpcompute.NetworksService
	targetPools       *gcpcompute.TargetPoolsService
	regions           *gcpcompute.RegionsService
	zones             *gcpcompute.ZonesService
}

func NewClient(project string, service *gcpcompute.Service, logger logger) client {
	return client{
		project:           project,
		logger:            logger,
		service:           service,
		addresses:         service.Addresses,
		backendServices:   service.BackendServices,
		disks:             service.Disks,
		httpHealthChecks:  service.HttpHealthChecks,
		httpsHealthChecks: service.HttpsHealthChecks,
		instances:         service.Instances,
		instanceGroups:    service.InstanceGroups,
		firewalls:         service.Firewalls,
		forwardingRules:   service.ForwardingRules,
		networks:          service.Networks,
		targetPools:       service.TargetPools,
		regions:           service.Regions,
		zones:             service.Zones,
	}
}

func (c client) ListAddresses(region, filter string) (*gcpcompute.AddressList, error) {
	call := c.addresses.List(c.project, region)
	if filter != "" {
		call.Filter(fmt.Sprintf("name eq %s", filter))
	}
	return call.Do()
}

func (c client) DeleteAddress(region, address string) error {
	op, err := c.addresses.Delete(c.project, region, address).Do()
	if err != nil {
		return err
	}

	return c.waitOnDelete(op)
}

func (c client) ListBackendServices(filter string) (*gcpcompute.BackendServiceList, error) {
	call := c.backendServices.List(c.project)
	if filter != "" {
		call.Filter(fmt.Sprintf("name eq %s", filter))
	}
	return call.Do()
}

func (c client) DeleteBackendService(backendService string) error {
	op, err := c.backendServices.Delete(c.project, backendService).Do()
	if err != nil {
		return err
	}

	return c.waitOnDelete(op)
}

func (c client) ListDisks(zone, filter string) (*gcpcompute.DiskList, error) {
	call := c.disks.List(c.project, zone)
	if filter != "" {
		call.Filter(fmt.Sprintf("name eq %s", filter))
	}
	return call.Do()
}

func (c client) DeleteDisk(zone, disk string) error {
	op, err := c.disks.Delete(c.project, zone, disk).Do()
	if err != nil {
		return err
	}

	return c.waitOnDelete(op)
}

func (c client) ListInstances(zone, filter string) (*gcpcompute.InstanceList, error) {
	call := c.instances.List(c.project, zone)
	if filter != "" {
		call.Filter(fmt.Sprintf("name eq %s", filter))
	}
	return call.Do()
}

func (c client) DeleteInstance(zone, instance string) error {
	op, err := c.instances.Delete(c.project, zone, instance).Do()
	if err != nil {
		return err
	}

	return c.waitOnDelete(op)
}

func (c client) ListInstanceGroups(zone, filter string) (*gcpcompute.InstanceGroupList, error) {
	call := c.instanceGroups.List(c.project, zone)
	if filter != "" {
		call.Filter(fmt.Sprintf("name eq %s", filter))
	}
	return call.Do()
}

func (c client) DeleteInstanceGroup(zone, instanceGroup string) error {
	op, err := c.instanceGroups.Delete(c.project, zone, instanceGroup).Do()
	if err != nil {
		return err
	}

	return c.waitOnDelete(op)
}

func (c client) ListHttpHealthChecks(filter string) (*gcpcompute.HttpHealthCheckList, error) {
	call := c.httpHealthChecks.List(c.project)
	if filter != "" {
		call.Filter(fmt.Sprintf("name eq %s", filter))
	}
	return call.Do()
}

func (c client) DeleteHttpHealthCheck(httpHealthCheck string) error {
	op, err := c.httpHealthChecks.Delete(c.project, httpHealthCheck).Do()
	if err != nil {
		return err
	}

	return c.waitOnDelete(op)
}

func (c client) ListHttpsHealthChecks(filter string) (*gcpcompute.HttpsHealthCheckList, error) {
	call := c.httpsHealthChecks.List(c.project)
	if filter != "" {
		call.Filter(fmt.Sprintf("name eq %s", filter))
	}
	return call.Do()
}

func (c client) DeleteHttpsHealthCheck(httpsHealthCheck string) error {
	op, err := c.httpsHealthChecks.Delete(c.project, httpsHealthCheck).Do()
	if err != nil {
		return err
	}

	return c.waitOnDelete(op)
}

func (c client) ListFirewalls(filter string) (*gcpcompute.FirewallList, error) {
	call := c.firewalls.List(c.project)
	if filter != "" {
		call.Filter(fmt.Sprintf("name eq %s", filter))
	}
	return call.Do()
}

func (c client) DeleteFirewall(firewall string) error {
	op, err := c.firewalls.Delete(c.project, firewall).Do()
	if err != nil {
		return err
	}

	return c.waitOnDelete(op)
}

func (c client) ListForwardingRules(region, filter string) (*gcpcompute.ForwardingRuleList, error) {
	call := c.forwardingRules.List(c.project, region)
	if filter != "" {
		call.Filter(fmt.Sprintf("name eq %s", filter))
	}
	return call.Do()
}

func (c client) DeleteForwardingRule(region, forwardingRule string) error {
	op, err := c.forwardingRules.Delete(c.project, region, forwardingRule).Do()
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

func (c client) ListTargetPools(region string) (*gcpcompute.TargetPoolList, error) {
	return c.targetPools.List(c.project, region).Do()
}

func (c client) DeleteTargetPool(region string, targetPool string) error {
	op, err := c.targetPools.Delete(c.project, region, targetPool).Do()
	if err != nil {
		return err
	}

	return c.waitOnDelete(op)
}

func (c client) ListRegions() (map[string]string, error) {
	regions := map[string]string{}

	list, err := c.regions.List(c.project).Do()
	if err != nil {
		return regions, err
	}

	for _, r := range list.Items {
		regions[r.SelfLink] = r.Name
	}
	return regions, nil
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
		op:      op,
		service: c.service,
		project: c.project,
		logger:  c.logger,
	}

	return waiter.Wait()
}
