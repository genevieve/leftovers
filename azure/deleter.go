package azure

import (
	"errors"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/arm/resources/resources"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
	azurelib "github.com/Azure/go-autorest/autorest/azure"
)

type logger interface {
	Printf(m string, a ...interface{})
	Prompt(m string) bool
}

type resource interface {
	Delete() error
}

type Deleter struct {
	resources []resource
}

func (d Deleter) Delete() error {
	for _, r := range d.resources {
		if err := r.Delete(); err != nil {
			return err
		}
	}
	return nil
}

func NewDeleter(logger logger, clientId, clientSecret, subscriptionId, tenantId string) (Deleter, error) {
	if clientId == "" {
		return Deleter{}, errors.New("Missing BBL_AZURE_CLIENT_ID.")
	}
	if clientSecret == "" {
		return Deleter{}, errors.New("Missing BBL_AZURE_CLIENT_SECRET.")
	}
	if subscriptionId == "" {
		return Deleter{}, errors.New("Missing BBL_AZURE_SUBSCRIPTION_ID.")
	}
	if tenantId == "" {
		return Deleter{}, errors.New("Missing BBL_AZURE_TENANT_ID.")
	}

	oauthConfig, err := adal.NewOAuthConfig(azurelib.PublicCloud.ActiveDirectoryEndpoint, tenantId)
	if err != nil {
		return Deleter{}, fmt.Errorf("Creating oauth config: %s\n", err)
	}

	servicePrincipalToken, err := adal.NewServicePrincipalToken(*oauthConfig, clientId, clientSecret, azurelib.PublicCloud.ResourceManagerEndpoint)
	if err != nil {
		return Deleter{}, fmt.Errorf("Creating service principal token: %s\n", err)
	}

	gc := resources.NewGroupsClient(subscriptionId)
	gc.ManagementClient.Authorizer = autorest.NewBearerAuthorizer(servicePrincipalToken)

	return Deleter{
		resources: []resource{NewGroups(gc, logger)},
	}, nil
}
