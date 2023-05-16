package azure

import (
	"errors"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/fatih/color"
	"github.com/genevieve/leftovers/common"
	multierror "github.com/hashicorp/go-multierror"
)

type resource interface {
	List(filter string) ([]common.Deletable, error)
	Type() string
}

type Leftovers struct {
	logger   logger
	resource resource
}

// List will print all of the resources that match the provided filter.
func (l Leftovers) List(filter string) {
	l.logger.NoConfirm()

	list, err := l.resource.List(filter)
	if err != nil {
		l.logger.Println(color.YellowString(err.Error()))
	}

	for _, r := range list {
		l.logger.Println(fmt.Sprintf("[%s: %s]", r.Type(), r.Name()))
	}
}

// ListByType defaults to List as there is only one resource type.
func (l Leftovers) ListByType(filter, rType string) {
	l.List(filter)
}

// Types will print all the resource types that can
// be deleted on this IaaS.
func (l Leftovers) Types() {
	l.logger.Println(l.resource.Type())
}

// Delete will collect all resources that contain
// the provided filter in the resource's identifier, prompt
// you to confirm deletion (if enabled), and delete those
// that are selected.
func (l Leftovers) Delete(filter string) error {
	var (
		deletables []common.Deletable
		result     *multierror.Error
	)

	// TODO: If they say no to the Resource Group, prompt for individual resources.
	deletables, err := l.resource.List(filter)
	if err != nil {
		l.logger.Println(color.YellowString(err.Error()))
	}

	for _, d := range deletables {
		l.logger.Println(fmt.Sprintf("[%s: %s] Deleting...", d.Type(), d.Name()))

		err := d.Delete()
		if err != nil {
			err = fmt.Errorf("[%s: %s] %s", d.Type(), d.Name(), color.YellowString(err.Error()))
			result = multierror.Append(result, err)

			l.logger.Println(err.Error())
		} else {
			l.logger.Println(fmt.Sprintf("[%s: %s] %s", d.Type(), d.Name(), color.GreenString("Deleted!")))
		}
	}

	return result.ErrorOrNil()
}

// DeleteByType will collect all resources of the provied type that contain
// the provided filter in the resource's identifier, prompt
// you to confirm deletion (if enabled), and delete those
// that are selected.
func (l Leftovers) DeleteByType(filter, rType string) error {
	return l.Delete(filter)
}

// NewLeftovers returns a new Leftovers for Azure that can be used to list resources,
// list types, or delete resources for the provided account. It returns an error
// if the credentials provided are invalid.
func NewLeftovers(logger logger, clientId, clientSecret, subscriptionId, tenantId string) (Leftovers, error) {
	if clientId == "" {
		return Leftovers{}, errors.New("Missing client id.")
	}

	if clientSecret == "" {
		return Leftovers{}, errors.New("Missing client secret.")
	}

	if subscriptionId == "" {
		return Leftovers{}, errors.New("Missing subscription id.")
	}

	if tenantId == "" {
		return Leftovers{}, errors.New("Missing tenant id.")
	}

	credential, err := azidentity.NewClientSecretCredential(tenantId, clientId, clientSecret, nil)
	if err != nil {
		return Leftovers{}, fmt.Errorf("Creating service principal token: %s\n", err)
	}

	rg, err := armresources.NewResourceGroupsClient(subscriptionId, credential, nil)
	if err != nil {
		return Leftovers{}, fmt.Errorf("Creating resource groups client: %s\n", err)
	}

	sg, err := armnetwork.NewApplicationSecurityGroupsClient(subscriptionId, credential, nil)
	if err != nil {
		return Leftovers{}, fmt.Errorf("Creating application security groups client: %s\n", err)
	}

	client := NewClient(*rg, *sg)

	return Leftovers{
		logger:   logger,
		resource: NewGroups(client, logger),
	}, nil
}
