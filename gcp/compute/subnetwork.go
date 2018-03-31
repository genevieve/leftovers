package compute

import "fmt"

type Subnetwork struct {
	client      subnetworksClient
	name        string
	clearerName string
	region      string
}

func NewSubnetwork(client subnetworksClient, name, region, network string) Subnetwork {
	clearerName := name
	if network != "" {
		clearerName = fmt.Sprintf("%s (Network:%s)", name, network)
	}

	return Subnetwork{
		client:      client,
		name:        name,
		clearerName: clearerName,
		region:      region,
	}
}

func (s Subnetwork) Delete() error {
	err := s.client.DeleteSubnetwork(s.region, s.name)

	if err != nil {
		return fmt.Errorf("Delete: %s", err)
	}

	return nil
}

func (s Subnetwork) Name() string {
	return s.clearerName
}

func (s Subnetwork) Type() string {
	return "Subnetwork"
}
