package fakes

import gcpcompute "google.golang.org/api/compute/v1"

type TargetVpnGatewaysClient struct {
	ListTargetVpnGatewaysCall struct {
		CallCount int
		Receives  struct {
			Region string
		}
		Returns struct {
			Output []*gcpcompute.TargetVpnGateway
			Error  error
		}
	}

	DeleteTargetVpnGatewayCall struct {
		CallCount int
		Receives  struct {
			TargetVpnGateway string
			Region           string
		}
		Returns struct {
			Error error
		}
	}
}

func (u *TargetVpnGatewaysClient) ListTargetVpnGateways(region string) ([]*gcpcompute.TargetVpnGateway, error) {
	u.ListTargetVpnGatewaysCall.CallCount++
	u.ListTargetVpnGatewaysCall.Receives.Region = region

	return u.ListTargetVpnGatewaysCall.Returns.Output, u.ListTargetVpnGatewaysCall.Returns.Error
}

func (u *TargetVpnGatewaysClient) DeleteTargetVpnGateway(region, targetVpnGateway string) error {
	u.DeleteTargetVpnGatewayCall.CallCount++
	u.DeleteTargetVpnGatewayCall.Receives.Region = region
	u.DeleteTargetVpnGatewayCall.Receives.TargetVpnGateway = targetVpnGateway

	return u.DeleteTargetVpnGatewayCall.Returns.Error
}
