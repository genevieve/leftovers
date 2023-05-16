package acceptance

import (
	"context"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
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
		Logger:         app.NewLogger(os.Stdin, os.Stdout, true, false),
	}
}

func (a AzureAcceptance) CreateResourceGroup(name string) {
	credential, err := azidentity.NewClientSecretCredential(a.TenantId, a.ClientId, a.ClientSecret, nil)
	Expect(err).NotTo(HaveOccurred())

	clientFactory, err := armresources.NewClientFactory(a.SubscriptionId, credential, nil)
	Expect(err).NotTo(HaveOccurred())

	groupsClient := clientFactory.NewResourceGroupsClient()

	location := "West US"
	group := armresources.ResourceGroup{Name: &name, Location: &location}
	_, err = groupsClient.CreateOrUpdate(context.Background(), name, group, nil)
	Expect(err).NotTo(HaveOccurred())
}
