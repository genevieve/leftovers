package fakes

import (
	"sync"

	gcp "google.golang.org/api/compute/v1"
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
	ListFirewallsCall struct {
		sync.Mutex
		CallCount int
		Returns   struct {
			FirewallSlice []*gcp.Firewall
			Error         error
		}
		Stub func() ([]*gcp.Firewall, error)
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
func (f *FirewallsClient) ListFirewalls() ([]*gcp.Firewall, error) {
	f.ListFirewallsCall.Lock()
	defer f.ListFirewallsCall.Unlock()
	f.ListFirewallsCall.CallCount++
	if f.ListFirewallsCall.Stub != nil {
		return f.ListFirewallsCall.Stub()
	}
	return f.ListFirewallsCall.Returns.FirewallSlice, f.ListFirewallsCall.Returns.Error
}
