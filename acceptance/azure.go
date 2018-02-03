package acceptance

import (
	"os"
	"strings"

	"github.com/Azure/azure-sdk-for-go/arm/resources/resources"
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

func NewAzureAcceptance() *AzureAcceptance {
	return &AzureAcceptance{}
}

func (a *AzureAcceptance) ReadyToTest() bool {
	iaas := os.Getenv("LEFTOVERS_ACCEPTANCE")
	if iaas == "" {
		return false
	}

	if strings.ToLower(iaas) != "azure" {
		return false
	}

	a.SubscriptionId = os.Getenv("BBL_AZURE_SUBSCRIPTION_ID")
	a.TenantId = os.Getenv("BBL_AZURE_TENANT_ID")
	a.ClientId = os.Getenv("BBL_AZURE_CLIENT_ID")
	a.ClientSecret = os.Getenv("BBL_AZURE_CLIENT_SECRET")

	logger := app.NewLogger(os.Stdin, os.Stdout, true)
	a.Logger = logger

	return true
}

func (a *AzureAcceptance) CreateResourceGroup(name string) {
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
