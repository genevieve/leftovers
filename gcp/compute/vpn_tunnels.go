package compute

import (
	"fmt"

	"github.com/genevieve/leftovers/common"
	gcpcompute "google.golang.org/api/compute/v1"
)

//go:generate faux --interface vpnTunnelsClient --output fakes/vpn_tunnels_client.go
type vpnTunnelsClient interface {
	ListVpnTunnels(region string) ([]*gcpcompute.VpnTunnel, error)
	DeleteVpnTunnel(region, vpnTunnel string) error
}

type VpnTunnels struct {
	client  vpnTunnelsClient
	logger  logger
	regions map[string]string
}

func NewVpnTunnels(client vpnTunnelsClient, logger logger, regions map[string]string) VpnTunnels {
	return VpnTunnels{
		client:  client,
		logger:  logger,
		regions: regions,
	}
}

func (v VpnTunnels) List(filter string, regex bool) ([]common.Deletable, error) {
	tunnels := []*gcpcompute.VpnTunnel{}

	for _, region := range v.regions {
		v.logger.Debugf("Listing Vpn Tunnels for Region %s...\n", region)
		l, err := v.client.ListVpnTunnels(region)
		if err != nil {
			return nil, fmt.Errorf("List Vpn Tunnels: %s", err)
		}

		tunnels = append(tunnels, l...)
	}

	var resources []common.Deletable

	for _, t := range tunnels {
		resource := NewVpnTunnel(v.client, t.Name, v.regions[t.Region])

		if !common.ResourceMatches(resource.Name(), filter, regex) {
			continue
		}

		proceed := v.logger.PromptWithDetails(resource.Type(), resource.Name())
		if !proceed {
			continue
		}

		resources = append(resources, resource)
	}

	return resources, nil
}

func (VpnTunnels) Type() string {
	return "vpn-tunnel"
}
