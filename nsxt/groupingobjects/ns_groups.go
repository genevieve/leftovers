package groupingobjects

import (
	"context"
	"fmt"
	"github.com/genevieve/leftovers/common"
)

type NSGroups struct {
	client groupingObjectsAPI
	ctx    context.Context
	logger logger
}

func NewNSGroups(client groupingObjectsAPI, ctx context.Context, logger logger) NSGroups {
	return NSGroups{
		client: client,
		ctx:    ctx,
		logger: logger,
	}
}

func (n NSGroups) List(filter string, regex bool) ([]common.Deletable, error) {
	n.logger.Debugln("Listing NS Groups...")
	result, _, err := n.client.ListNSGroups(n.ctx, map[string]interface{}{})

	if err != nil {
		return []common.Deletable{}, fmt.Errorf("List NS Groups: %s", err)
	}

	var resources []common.Deletable
	for _, nsGroup := range result.Results {
		resource := NewNSGroup(n.client, n.ctx, nsGroup.DisplayName, nsGroup.Id)

		if !common.ResourceMatches(nsGroup.DisplayName, filter, regex) {
			continue
		}

		proceed := n.logger.PromptWithDetails(resource.Type(), resource.Name())
		if !proceed {
			continue
		}

		resources = append(resources, resource)
	}

	return resources, nil
}

func (i NSGroups) Type() string {
	return "NS Group"
}
