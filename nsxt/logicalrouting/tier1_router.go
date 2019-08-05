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
	options := map[string]interface{}{
		"force": true,
	}
	_, err := t.client.DeleteLogicalRouter(t.ctx, t.id, options)
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
