package fakes

import (
	"sync"

	gcpcompute "google.golang.org/api/compute/v1"
)

type FirewallsClient struct {
	DeleteFirewallCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Firewall string
		}
		Returns struct {
			Error error
		}
		Stub func(string) error
	}
	GetNetworkNameCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Url string
		}
		Returns struct {
			Name string
		}
		Stub func(string) string
	}
	ListFirewallsCall struct {
		sync.Mutex
		CallCount int
		Returns   struct {
			FirewallSlice []*gcpcompute.Firewall
			Error         error
		}
		Stub func() ([]*gcpcompute.Firewall, error)
	}
}

func (f *FirewallsClient) DeleteFirewall(param1 string) error {
	f.DeleteFirewallCall.Lock()
	defer f.DeleteFirewallCall.Unlock()
	f.DeleteFirewallCall.CallCount++
	f.DeleteFirewallCall.Receives.Firewall = param1
	if f.DeleteFirewallCall.Stub != nil {
		return f.DeleteFirewallCall.Stub(param1)
	}
	return f.DeleteFirewallCall.Returns.Error
}
func (f *FirewallsClient) GetNetworkName(param1 string) string {
	f.GetNetworkNameCall.Lock()
	defer f.GetNetworkNameCall.Unlock()
	f.GetNetworkNameCall.CallCount++
	f.GetNetworkNameCall.Receives.Url = param1
	if f.GetNetworkNameCall.Stub != nil {
		return f.GetNetworkNameCall.Stub(param1)
	}
	return f.GetNetworkNameCall.Returns.Name
}
func (f *FirewallsClient) ListFirewalls() ([]*gcpcompute.Firewall, error) {
	f.ListFirewallsCall.Lock()
	defer f.ListFirewallsCall.Unlock()
	f.ListFirewallsCall.CallCount++
	if f.ListFirewallsCall.Stub != nil {
		return f.ListFirewallsCall.Stub()
	}
	return f.ListFirewallsCall.Returns.FirewallSlice, f.ListFirewallsCall.Returns.Error
}
