package azure

import (
	"fmt"
	"strings"

	"github.com/genevieve/leftovers/common"
)

type groupsClient interface {
	ListGroups() ([]string, error)
	DeleteGroup(string) error
}

type Groups struct {
	client groupsClient
	logger logger
}

func NewGroups(client groupsClient, logger logger) Groups {
	return Groups{
		client: client,
		logger: logger,
	}
}

func (g Groups) List(filter string) ([]common.Deletable, error) {
	g.logger.Debugln("Listing Resource Groups")
	groups, err := g.client.ListGroups()
	if err != nil {
		return []common.Deletable{}, fmt.Errorf("Listing Resource Groups: %s", err)
	}

	var resources []common.Deletable
	for _, group := range groups {
		r := NewGroup(g.client, group)

		if !strings.Contains(r.Name(), filter) {
			continue
		}

		proceed := g.logger.PromptWithDetails(r.Type(), r.Name())
		if !proceed {
			continue
		}

		resources = append(resources, r)
	}

	return resources, nil
}

func (g Groups) Type() string {
	return "resource-group"
}
