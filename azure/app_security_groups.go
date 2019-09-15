package azure

import (
	"fmt"
	"strings"

	"github.com/genevieve/leftovers/common"
)

type appSecurityGroupsClient interface {
	ListAppSecurityGroups(rgName string) ([]string, error)
	DeleteAppSecurityGroup(rgName string, name string) error
}

type AppSecurityGroups struct {
	client appSecurityGroupsClient
	rgName string
	logger logger
}

func NewAppSecurityGroups(client appSecurityGroupsClient, rgName string, logger logger) AppSecurityGroups {
	return AppSecurityGroups{
		client: client,
		rgName: rgName,
		logger: logger,
	}
}

func (g AppSecurityGroups) List(filter string) ([]common.Deletable, error) {
	groups, err := g.client.ListAppSecurityGroups(g.rgName)
	if err != nil {
		return nil, fmt.Errorf("Listing Application Security Groups: %s", err)
	}

	var resources []common.Deletable
	for _, group := range groups {
		r := NewAppSecurityGroup(g.client, g.rgName, group)

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

func (g AppSecurityGroups) Type() string {
	return "app-security-group"
}
