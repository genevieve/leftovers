package acceptance

import (
	"fmt"
	"net/http"
	"os"

	"github.com/genevieve/leftovers/app"
	. "github.com/onsi/gomega"
	nsxt "github.com/vmware/go-vmware-nsxt"
	"github.com/vmware/go-vmware-nsxt/manager"
)

type NSXTAcceptance struct {
	ManagerHost string
	User        string
	Password    string
	EdgeCluster string
	NSXTClient  *nsxt.APIClient
	Logger      *app.Logger
}

func NewNSXTAcceptance() NSXTAcceptance {
	managerHost := os.Getenv("NSXT_MANAGER_HOST")
	Expect(managerHost).NotTo(Equal(""), "Missing $NSXT_MANAGER_HOST.")

	nsxtUser := os.Getenv("NSXT_USER")
	Expect(nsxtUser).NotTo(Equal(""), "Missing $NSXT_USER.")

	nsxtPassword := os.Getenv("NSXT_PASSWORD")
	Expect(nsxtPassword).NotTo(Equal(""), "Missing $NSXT_PASSWORD.")

	nsxtClient, err := nsxt.NewAPIClient(&nsxt.Configuration{
		BasePath: fmt.Sprintf("https://%s/api/v1", managerHost),
		UserName: nsxtUser,
		Password: nsxtPassword,
		Host:     managerHost,
		Insecure: true,
		RetriesConfiguration: nsxt.ClientRetriesConfiguration{
			MaxRetries:    1,
			RetryMinDelay: 100,
			RetryMaxDelay: 500,
		},
	})
	Expect(err).NotTo(HaveOccurred())

	edgeCluster := os.Getenv("NSXT_EDGE_CLUSTER")
	Expect(edgeCluster).NotTo(Equal(""), "Missing $NSXT_EDGE_CLUSTER.")

	return NSXTAcceptance{
		ManagerHost: managerHost,
		User:        nsxtUser,
		Password:    nsxtPassword,
		NSXTClient:  nsxtClient,
		EdgeCluster: edgeCluster,
		Logger:      app.NewLogger(os.Stdin, os.Stdout, true),
	}
}

// adapted from the nsxt terraform provider's edge cluster data source read method
func (n *NSXTAcceptance) GetEdgeClusterID(name string) string {
	clusters, _, err := n.NSXTClient.NetworkTransportApi.ListEdgeClusters(n.NSXTClient.Context, nil)
	Expect(err).NotTo(HaveOccurred())

	var match manager.EdgeCluster
	for _, cluster := range clusters.Results {
		if cluster.DisplayName == name {
			match = cluster
			break
		}
	}
	Expect(match.DisplayName).NotTo(Equal(""), "Could not find edge cluster with specified name")

	return match.Id
}

func (n *NSXTAcceptance) CreateT1Router(name string) manager.LogicalRouter {
	logicalRouter := manager.LogicalRouter{
		DisplayName:   name,
		RouterType:    "TIER1",
		EdgeClusterId: n.GetEdgeClusterID(n.EdgeCluster),
	}

	logicalRouter, resp, err := n.NSXTClient.LogicalRoutingAndServicesApi.CreateLogicalRouter(n.NSXTClient.Context, logicalRouter)
	Expect(err).NotTo(HaveOccurred())

	Expect(resp.StatusCode).To(Equal(http.StatusCreated))

	return logicalRouter
}
