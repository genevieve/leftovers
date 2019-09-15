package azure

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/resources"
	"github.com/Azure/go-autorest/autorest"
)

type client struct {
	groupsClient   resources.GroupsClient
	autorestClient autorest.Client
}

func NewClient(groupsClient resources.GroupsClient) client {
	return client{
		groupsClient:   groupsClient,
		autorestClient: groupsClient.Client,
	}
}

func (c client) ListGroups() ([]string, error) {
	ctx := context.Background()
	groups := []string{}

	for list, err := c.groupsClient.ListComplete(ctx, "", nil); list.NotDone(); err = list.Next() {
		if err != nil {
			return []string{}, fmt.Errorf("List Groups: %s", err)
		}

		groups = append(groups, *list.Value().Name)
	}

	return groups, nil
}

func (c client) DeleteGroup(name string) error {
	ctx := context.Background()

	future, err := c.groupsClient.Delete(ctx, name)
	if err != nil {
		return err
	}

	err = future.WaitForCompletionRef(ctx, c.autorestClient)
	if err != nil {
		return fmt.Errorf("Waiting for completion: %s", err)
	}

	return nil
}
