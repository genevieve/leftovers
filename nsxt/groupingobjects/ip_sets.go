package groupingobjects

import (
	"context"
	"fmt"
	"github.com/genevieve/leftovers/common"
)

type IPSets struct {
	client groupingObjectsAPI
	ctx    context.Context
	logger logger
}

func NewIPSets(client groupingObjectsAPI, ctx context.Context, logger logger) IPSets {
	return IPSets{
		client: client,
		ctx:    ctx,
		logger: logger,
	}
}

func (i IPSets) List(filter string, regex bool) ([]common.Deletable, error) {
	i.logger.Debugln("Listing IP Sets...")
	result, _, err := i.client.ListIPSets(i.ctx, map[string]interface{}{})

	if err != nil {
		return []common.Deletable{}, fmt.Errorf("List IP Sets: %s", err)
	}

	var resources []common.Deletable
	for _, ipSet := range result.Results {
		resource := NewIPSet(i.client, i.ctx, ipSet.DisplayName, ipSet.Id)

		if !common.MatchRegex(ipSet.DisplayName, filter, regex) {
			continue
		}

		proceed := i.logger.PromptWithDetails(resource.Type(), resource.Name())
		if !proceed {
			continue
		}

		resources = append(resources, resource)
	}

	return resources, nil
}

func (i IPSets) Type() string {
	return "IP Set"
}
