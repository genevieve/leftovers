package fakes

import "github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/floatingips"

type FloatingIPsClient struct {
	ListFloatingIPsCall struct {
		CallCount int
		Returns   struct {
			Output []floatingips.FloatingIP
			Error  error
		}
	}

	DeleteFloatingIPCall struct {
		CallCount int
		Receives  struct {
			FloatingIP string
		}
		Returns struct {
			Error error
		}
	}
}

func (n *FloatingIPsClient) ListFloatingIPs() ([]floatingips.FloatingIP, error) {
	n.ListFloatingIPsCall.CallCount++

	return n.ListFloatingIPsCall.Returns.Output, n.ListFloatingIPsCall.Returns.Error
}

func (n *FloatingIPsClient) DeleteFloatingIP(ip string) error {
	n.DeleteFloatingIPCall.CallCount++
	n.DeleteFloatingIPCall.Receives.FloatingIP = ip

	return n.DeleteFloatingIPCall.Returns.Error
}
