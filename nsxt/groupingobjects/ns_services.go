package groupingobjects

import (
	"context"
	"fmt"
	"github.com/genevieve/leftovers/common"
)

type NSServices struct {
	client groupingObjectsAPI
	ctx    context.Context
	logger logger
}

func NewNSServices(client groupingObjectsAPI, ctx context.Context, logger logger) NSServices {
	return NSServices{
		client: client,
		ctx:    ctx,
		logger: logger,
	}
}

func (n NSServices) List(filter string, regex bool) ([]common.Deletable, error) {
	n.logger.Debugln("Listing NS Services...")
	result, _, err := n.client.ListNSServices(n.ctx, map[string]interface{}{})

	if err != nil {
		return []common.Deletable{}, fmt.Errorf("List NS Services: %s", err)
	}

	var resources []common.Deletable
	for _, nsService := range result.Results {
		resource := NewNSService(n.client, n.ctx, nsService.DisplayName, nsService.Id)

		if !common.ResourceMatches(nsService.DisplayName, filter, regex) {
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

func (i NSServices) Type() string {
	return "NS Service"
}
