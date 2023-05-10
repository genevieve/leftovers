package compute

import (
	"fmt"

	"github.com/genevieve/leftovers/common"
	gcpcompute "google.golang.org/api/compute/v1"
)

//go:generate faux --interface targetVpnGatewaysClient --output fakes/target_vpn_gateways_client.go
type targetVpnGatewaysClient interface {
	ListTargetVpnGateways(region string) ([]*gcpcompute.TargetVpnGateway, error)
	DeleteTargetVpnGateway(region, targetVpnGateway string) error
}

type TargetVpnGateways struct {
	client  targetVpnGatewaysClient
	logger  logger
	regions map[string]string
}

func NewTargetVpnGateways(client targetVpnGatewaysClient, logger logger, regions map[string]string) TargetVpnGateways {
	return TargetVpnGateways{
		client:  client,
		logger:  logger,
		regions: regions,
	}
}

func (t TargetVpnGateways) List(filter string, regex bool) ([]common.Deletable, error) {
	gateways := []*gcpcompute.TargetVpnGateway{}

	for _, region := range t.regions {
		t.logger.Debugf("Listing Target Vpn Gateways for region %s...\n", region)
		l, err := t.client.ListTargetVpnGateways(region)
		if err != nil {
			return nil, fmt.Errorf("List Target Vpn Gateways: %s", err)
		}

		gateways = append(gateways, l...)
	}

	var resources []common.Deletable

	for _, g := range gateways {
		resource := NewTargetVpnGateway(t.client, g.Name, t.regions[g.Region])

		if !common.ResourceMatches(resource.Name(), filter, regex) {
			continue
		}

		proceed := t.logger.PromptWithDetails(resource.Type(), resource.Name())
		if !proceed {
			continue
		}

		resources = append(resources, resource)
	}

	return resources, nil
}

func (TargetVpnGateways) Type() string {
	return "target-vpn-gateway"
}
