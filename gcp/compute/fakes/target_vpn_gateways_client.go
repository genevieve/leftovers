package fakes

import (
	"sync"

	gcpcompute "google.golang.org/api/compute/v1"
)

type TargetVpnGatewaysClient struct {
	DeleteTargetVpnGatewayCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Region           string
			TargetVpnGateway string
		}
		Returns struct {
			Error error
		}
		Stub func(string, string) error
	}
	ListTargetVpnGatewaysCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Region string
		}
		Returns struct {
			TargetVpnGatewaySlice []*gcpcompute.TargetVpnGateway
			Error                 error
		}
		Stub func(string) ([]*gcpcompute.TargetVpnGateway, error)
	}
}

func (f *TargetVpnGatewaysClient) DeleteTargetVpnGateway(param1 string, param2 string) error {
	f.DeleteTargetVpnGatewayCall.Lock()
	defer f.DeleteTargetVpnGatewayCall.Unlock()
	f.DeleteTargetVpnGatewayCall.CallCount++
	f.DeleteTargetVpnGatewayCall.Receives.Region = param1
	f.DeleteTargetVpnGatewayCall.Receives.TargetVpnGateway = param2
	if f.DeleteTargetVpnGatewayCall.Stub != nil {
		return f.DeleteTargetVpnGatewayCall.Stub(param1, param2)
	}
	return f.DeleteTargetVpnGatewayCall.Returns.Error
}
func (f *TargetVpnGatewaysClient) ListTargetVpnGateways(param1 string) ([]*gcpcompute.TargetVpnGateway, error) {
	f.ListTargetVpnGatewaysCall.Lock()
	defer f.ListTargetVpnGatewaysCall.Unlock()
	f.ListTargetVpnGatewaysCall.CallCount++
	f.ListTargetVpnGatewaysCall.Receives.Region = param1
	if f.ListTargetVpnGatewaysCall.Stub != nil {
		return f.ListTargetVpnGatewaysCall.Stub(param1)
	}
	return f.ListTargetVpnGatewaysCall.Returns.TargetVpnGatewaySlice, f.ListTargetVpnGatewaysCall.Returns.Error
}
