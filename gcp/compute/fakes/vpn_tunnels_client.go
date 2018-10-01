package fakes

import gcpcompute "google.golang.org/api/compute/v1"

type VpnTunnelsClient struct {
	ListVpnTunnelsCall struct {
		CallCount int
		Receives  struct {
			Region string
		}
		Returns struct {
			Output []*gcpcompute.VpnTunnel
			Error  error
		}
	}

	DeleteVpnTunnelCall struct {
		CallCount int
		Receives  struct {
			VpnTunnel string
			Region    string
		}
		Returns struct {
			Error error
		}
	}
}

func (u *VpnTunnelsClient) ListVpnTunnels(region string) ([]*gcpcompute.VpnTunnel, error) {
	u.ListVpnTunnelsCall.CallCount++
	u.ListVpnTunnelsCall.Receives.Region = region

	return u.ListVpnTunnelsCall.Returns.Output, u.ListVpnTunnelsCall.Returns.Error
}

func (u *VpnTunnelsClient) DeleteVpnTunnel(region, vpnTunnel string) error {
	u.DeleteVpnTunnelCall.CallCount++
	u.DeleteVpnTunnelCall.Receives.Region = region
	u.DeleteVpnTunnelCall.Receives.VpnTunnel = vpnTunnel

	return u.DeleteVpnTunnelCall.Returns.Error
}
