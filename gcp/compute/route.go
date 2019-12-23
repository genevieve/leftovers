package compute

import (
	"fmt"
)

type Route struct {
	client      routesClient
	name        string
	clearerName string
}

func NewRoute(client routesClient, name, network string) Route {
	clearerName := name

	networkName := client.GetNetworkName(network)
	if len(networkName) > 0 {
		clearerName = fmt.Sprintf("%s (%s)", name, networkName)
	}

	return Route{
		client:      client,
		name:        name,
		clearerName: clearerName,
	}
}

func (r Route) Delete() error {
	err := r.client.DeleteRoute(r.name)
	if err != nil {
		return fmt.Errorf("Delete: %s", err)
	}

	return nil
}

func (r Route) Name() string {
	return r.clearerName
}

func (r Route) Type() string {
	return "Route"
}
