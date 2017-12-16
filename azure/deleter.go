package azure

import (
	"log"

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

func NewDeleter(logger logger, clientId, clientSecret, subscriptionId, tenantId string) Deleter {
	if clientId == "" {
		log.Fatal("Missing AZURE_CLIENT_ID.")
	}
	if clientSecret == "" {
		log.Fatal("Missing AZURE_CLIENT_SECRET.")
	}
	if subscriptionId == "" {
		log.Fatal("Missing AZURE_SUBSCRIPTION_ID.")
	}
	if tenantId == "" {
		log.Fatal("Missing AZURE_TENANT_ID.")
	}

	oauthConfig, err := adal.NewOAuthConfig(azurelib.PublicCloud.ActiveDirectoryEndpoint, tenantId)
	if err != nil {
		log.Fatalf("Creating oauth config: %s\n", err)
	}

	servicePrincipalToken, err := adal.NewServicePrincipalToken(*oauthConfig, clientId, clientSecret, azurelib.PublicCloud.ResourceManagerEndpoint)
	if err != nil {
		log.Fatalf("Creating service principal token: %s\n", err)
	}

	gc := resources.NewGroupsClient(subscriptionId)
	gc.ManagementClient.Authorizer = autorest.NewBearerAuthorizer(servicePrincipalToken)

	return Deleter{
		resources: []resource{NewGroups(gc, logger)},
	}
}
