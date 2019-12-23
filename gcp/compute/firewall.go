package compute

import "fmt"

type Firewall struct {
	client      firewallsClient
	name        string
	clearerName string
}

func NewFirewall(client firewallsClient, name, network string) Firewall {
	clearerName := name

	networkName := client.GetNetworkName(network)
	if len(networkName) > 0 {
		clearerName = fmt.Sprintf("%s (%s)", name, networkName)
	}

	return Firewall{
		client:      client,
		name:        name,
		clearerName: clearerName,
	}
}

func (f Firewall) Delete() error {
	err := f.client.DeleteFirewall(f.name)
	if err != nil {
		return fmt.Errorf("Delete: %s", err)
	}

	return nil
}

func (f Firewall) Name() string {
	return f.clearerName
}

func (f Firewall) Type() string {
	return "Firewall"
}
