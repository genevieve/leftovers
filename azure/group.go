package azure

import (
	"fmt"
)

type Group struct {
	client groupsClient
	name   string
}

// Group represents an Azure resource group.
func NewGroup(client groupsClient, name string) Group {
	return Group{
		client: client,
		name:   name,
	}
}

// Delete deletes an Azure resource group and all other Azure
// resources in the resource group.
func (g Group) Delete() error {
	err := g.client.DeleteGroup(g.name)
	if err != nil {
		return fmt.Errorf("Delete: %s", err)
	}

	return nil
}

func (g Group) Name() string {
	return g.name
}

func (g Group) Type() string {
	return "Resource Group"
}
