package fakes

import (
	"sync"

	gcpcompute "google.golang.org/api/compute/v1"
)

type VpnTunnelsClient struct {
	DeleteVpnTunnelCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Region    string
			VpnTunnel string
		}
		Returns struct {
			Error error
		}
		Stub func(string, string) error
	}
	ListVpnTunnelsCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Region string
		}
		Returns struct {
			VpnTunnelSlice []*gcpcompute.VpnTunnel
			Error          error
		}
		Stub func(string) ([]*gcpcompute.VpnTunnel, error)
	}
}

func (f *VpnTunnelsClient) DeleteVpnTunnel(param1 string, param2 string) error {
	f.DeleteVpnTunnelCall.Lock()
	defer f.DeleteVpnTunnelCall.Unlock()
	f.DeleteVpnTunnelCall.CallCount++
	f.DeleteVpnTunnelCall.Receives.Region = param1
	f.DeleteVpnTunnelCall.Receives.VpnTunnel = param2
	if f.DeleteVpnTunnelCall.Stub != nil {
		return f.DeleteVpnTunnelCall.Stub(param1, param2)
	}
	return f.DeleteVpnTunnelCall.Returns.Error
}
func (f *VpnTunnelsClient) ListVpnTunnels(param1 string) ([]*gcpcompute.VpnTunnel, error) {
	f.ListVpnTunnelsCall.Lock()
	defer f.ListVpnTunnelsCall.Unlock()
	f.ListVpnTunnelsCall.CallCount++
	f.ListVpnTunnelsCall.Receives.Region = param1
	if f.ListVpnTunnelsCall.Stub != nil {
		return f.ListVpnTunnelsCall.Stub(param1)
	}
	return f.ListVpnTunnelsCall.Returns.VpnTunnelSlice, f.ListVpnTunnelsCall.Returns.Error
}
