package azure

import (
	"fmt"

	"github.com/Azure/azure-sdk-for-go/arm/resources/resources"
	"github.com/Azure/go-autorest/autorest"
)

type groupsClient interface {
	List(filter string, top *int32) (resources.GroupListResult, error)
	Delete(name string, channel <-chan struct{}) (<-chan autorest.Response, <-chan error)
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

func (r Groups) Delete() error {
	groups, err := r.client.List("", nil)
	if err != nil {
		return fmt.Errorf("Listing resource groups: %s", err)
	}

	for _, g := range *groups.Value {
		n := *g.Name

		proceed := r.logger.Prompt(fmt.Sprintf("Are you sure you want to delete resource group %s?", n))
		if !proceed {
			continue
		}

		_, errChan := r.client.Delete(n, nil)
		err = <-errChan
		if err == nil {
			r.logger.Printf("SUCCESS deleting resource group %s\n", n)
		} else {
			r.logger.Printf("ERROR deleting resource group %s: %s\n", n, err)
		}
	}

	return nil
}
