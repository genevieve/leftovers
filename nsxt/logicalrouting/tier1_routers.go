package logicalrouting

import (
	"context"
	"fmt"
	"github.com/genevieve/leftovers/common"
)

type Tier1Routers struct {
	client logicalRoutingAPI
	ctx    context.Context
	logger logger
}

func NewTier1Routers(client logicalRoutingAPI, ctx context.Context, logger logger) Tier1Routers {
	return Tier1Routers{
		client: client,
		ctx:    ctx,
		logger: logger,
	}
}

func (t Tier1Routers) List(filter string, regex bool) ([]common.Deletable, error) {
	t.logger.Debugln("Listing Tier 1 Routers...")
	result, _, err := t.client.ListLogicalRouters(t.ctx, map[string]interface{}{
		"routerType": "TIER1",
	})

	if err != nil {
		return []common.Deletable{}, fmt.Errorf("List Tier 1 Routers: %s", err)
	}

	var resources []common.Deletable
	for _, router := range result.Results {
		if router.SystemOwned {
			continue
		}

		resource := NewTier1Router(t.client, t.ctx, router.DisplayName, router.Id)

		if !common.ResourceMatches(router.DisplayName, filter, regex) {
			continue
		}

		proceed := t.logger.PromptWithDetails(resource.Type(), resource.Name())
		if !proceed {
			continue
		}

		resources = append(resources, resource)
	}

	return resources, nil
}

func (t Tier1Routers) Type() string {
	return "Tier 1 Router"
}
