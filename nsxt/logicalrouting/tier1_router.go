package logicalrouting

import (
	"context"
	"fmt"
)

type Tier1Router struct {
	client logicalRoutingAPI
	ctx    context.Context
	id     string
	name   string
}

func NewTier1Router(client logicalRoutingAPI, ctx context.Context, name, id string) Tier1Router {
	return Tier1Router{
		client: client,
		ctx:    ctx,
		name:   name,
		id:     id,
	}
}

func (t Tier1Router) Delete() error {
	// For each logical router, delete it's link ports?
	options := map[string]interface{}{
		"force":           true,
		"logicalRouterId": t.id,
	}
	ports, _, err := t.client.ListLogicalRouterPorts(t.ctx, options)
	if err != nil {
		return fmt.Errorf("List Logical Router Ports: %s", err)
	}

	options = map[string]interface{}{
		"force": true,
	}
	for _, p := range ports.Results {
		_, err := t.client.DeleteLogicalRouterPort(t.ctx, p.Id, options)
		if err != nil {
			return fmt.Errorf("Delete Logical Router Port: %s", err)
		}
	}

	_, err = t.client.DeleteLogicalRouter(t.ctx, t.id, options)
	if err != nil {
		return fmt.Errorf("Delete: %s", err)
	}

	return nil
}

func (t Tier1Router) Name() string {
	return t.name
}

func (t Tier1Router) Type() string {
	return "Tier 1 Router"
}
