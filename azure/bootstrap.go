package azure

import (
	"fmt"
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

func Bootstrap(logger logger, clientId, clientSecret, subscriptionId, tenantId string) {
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
		panic(err)
	}

	servicePrincipalToken, err := adal.NewServicePrincipalToken(*oauthConfig, clientId, clientSecret, azurelib.PublicCloud.ResourceManagerEndpoint)
	if err != nil {
		panic(err)
	}

	gc := resources.NewGroupsClient(subscriptionId)
	gc.ManagementClient.Authorizer = autorest.NewBearerAuthorizer(servicePrincipalToken)

	result, err := gc.List("", nil)
	if err != nil {
		panic(err)
	}

	for _, g := range *result.Value {
		fmt.Println(*g.Name)
	}
}
