package azure

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
)

type client struct {
	rgClient armresources.ResourceGroupsClient
	sgClient armnetwork.ApplicationSecurityGroupsClient
}

func NewClient(rg armresources.ResourceGroupsClient, sg armnetwork.ApplicationSecurityGroupsClient) client {
	return client{
		rgClient: rg,
		sgClient: sg,
	}
}

func (c client) ListGroups() ([]string, error) {
	ctx := context.Background()
	groups := []string{}

	pager := c.rgClient.NewListPager(nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			return []string{}, fmt.Errorf("List Complete Resource Groups: %s", err)
		}
		for _, group := range nextResult.Value {
			groups = append(groups, *group.Name)
		}
	}

	return groups, nil
}

func (c client) DeleteGroup(name string) error {
	ctx := context.Background()

	poller, err := c.rgClient.BeginDelete(ctx, name, nil)
	if err != nil {
		return err
	}

	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c client) ListAppSecurityGroups(rgName string) ([]string, error) {
	ctx := context.Background()
	groups := []string{}

	pager := c.sgClient.NewListAllPager(nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			return groups, fmt.Errorf("List Complete App Security Groups: %s", err)
		}
		for _, group := range nextResult.Value {
			groups = append(groups, *group.Name)
		}
	}

	return groups, nil
}

func (c client) DeleteAppSecurityGroup(rgName, name string) error {
	ctx := context.Background()

	poller, err := c.sgClient.BeginDelete(ctx, rgName, name, nil)
	if err != nil {
		return err
	}

	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to pull the result %s", err)
	}

	return nil
}
