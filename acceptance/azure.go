package acceptance

import (
	"os"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/resources"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/genevieve/leftovers/app"
	. "github.com/onsi/gomega"
)

type AzureAcceptance struct {
	SubscriptionId string
	TenantId       string
	ClientId       string
	ClientSecret   string
	Logger         *app.Logger
}

func NewAzureAcceptance() AzureAcceptance {
	subscriptionId := os.Getenv("BBL_AZURE_SUBSCRIPTION_ID")
	Expect(subscriptionId).NotTo(Equal(""), "Missing $BBL_AZURE_SUBSCRIPTION_ID.")

	tenantId := os.Getenv("BBL_AZURE_TENANT_ID")
	Expect(tenantId).NotTo(Equal(""), "Missing $BBL_AZURE_TENANT_ID.")

	clientId := os.Getenv("BBL_AZURE_CLIENT_ID")
	Expect(clientId).NotTo(Equal(""), "Missing $BBL_AZURE_CLIENT_ID.")

	clientSecret := os.Getenv("BBL_AZURE_CLIENT_SECRET")
	Expect(clientSecret).NotTo(Equal(""), "Missing $BBL_AZURE_CLIENT_SECRET.")

	return AzureAcceptance{
		SubscriptionId: subscriptionId,
		TenantId:       tenantId,
		ClientId:       clientId,
		ClientSecret:   clientSecret,
		Logger:         app.NewLogger(os.Stdin, os.Stdout, true),
	}
}

func (a AzureAcceptance) CreateResourceGroup(name string) {
	oauthConfig, err := adal.NewOAuthConfig(azure.PublicCloud.ActiveDirectoryEndpoint, a.TenantId)
	Expect(err).NotTo(HaveOccurred())

	servicePrincipalToken, err := adal.NewServicePrincipalToken(*oauthConfig, a.ClientId, a.ClientSecret, azure.PublicCloud.ResourceManagerEndpoint)
	Expect(err).NotTo(HaveOccurred())

	groupsClient := resources.NewGroupsClient(a.SubscriptionId)
	groupsClient.Authorizer = autorest.NewBearerAuthorizer(servicePrincipalToken)
	groupsClient.Sender = autorest.CreateSender(autorest.AsIs())

	location := "West US"
	_, err = groupsClient.CreateOrUpdate(name, resources.Group{Name: &name, Location: &location})
	Expect(err).NotTo(HaveOccurred())
}
