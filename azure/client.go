package azure

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-04-01/network"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/resources"
	"github.com/Azure/go-autorest/autorest"
)

type client struct {
	rgClient       resources.GroupsClient
	sgClient       network.ApplicationSecurityGroupsClient
	autorestClient autorest.Client
}

func NewClient(rg resources.GroupsClient, sg network.ApplicationSecurityGroupsClient) client {
	return client{
		rgClient:       rg,
		sgClient:       sg,
		autorestClient: rg.Client,
	}
}

func (c client) ListGroups() ([]string, error) {
	ctx := context.Background()
	groups := []string{}

	for list, err := c.rgClient.ListComplete(ctx, "", nil); list.NotDone(); err = list.Next() {
		if err != nil {
			return []string{}, fmt.Errorf("List Complete Resource Groups: %s", err)
		}

		groups = append(groups, *list.Value().Name)
	}

	return groups, nil
}

func (c client) DeleteGroup(name string) error {
	ctx := context.Background()

	future, err := c.rgClient.Delete(ctx, name)
	if err != nil {
		return err
	}

	err = future.WaitForCompletionRef(ctx, c.autorestClient)
	if err != nil {
		return fmt.Errorf("Waiting for completion: %s", err)
	}

	return nil
}

func (c client) ListAppSecurityGroups(rgName string) ([]string, error) {
	ctx := context.Background()
	groups := []string{}

	for list, err := c.sgClient.ListComplete(ctx, rgName); list.NotDone(); err = list.Next() {
		if err != nil {
			return groups, fmt.Errorf("List Complete App Security Groups: %s", err)
		}

		groups = append(groups, *list.Value().Name)
	}

	return groups, nil
}

func (c client) DeleteAppSecurityGroup(rgName, name string) error {
	ctx := context.Background()

	future, err := c.sgClient.Delete(ctx, rgName, name)
	if err != nil {
		return err
	}

	err = future.WaitForCompletionRef(ctx, c.autorestClient)
	if err != nil {
		return fmt.Errorf("Waiting for completion: %s", err)
	}

	return nil
}
