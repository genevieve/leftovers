package gcp_test

import (
	"github.com/genevievelesperance/leftovers/gcp"
	"github.com/genevievelesperance/leftovers/gcp/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	compute "google.golang.org/api/compute/v1"
)

var _ = Describe("Networks", func() {
	var (
		client  *fakes.NetworksClient
		logger  *fakes.Logger
		project string

		networks gcp.Networks
	)

	BeforeEach(func() {
		client = &fakes.NetworksClient{}
		logger = &fakes.Logger{}
		project = "the-project"

		networks = gcp.NewNetworks(client, logger, project)
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			logger.PromptCall.Returns.Proceed = true
			client.ListCall.Returns.Output = &compute.NetworkList{
				Items: []*compute.Network{{
					Name: "banana",
				}},
			}
		})

		It("deletes networks", func() {
			err := networks.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(client.ListCall.CallCount).To(Equal(1))
			Expect(client.ListCall.Receives.Project).To(Equal("the-project"))

			Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete network banana?"))

			Expect(client.DeleteCall.CallCount).To(Equal(1))
			Expect(client.DeleteCall.Receives.Project).To(Equal("the-project"))
			Expect(client.DeleteCall.Receives.Network).To(Equal("banana"))

			Expect(logger.PrintfCall.Messages).To(Equal([]string{"SUCCESS deleting network banana\n"}))
		})
	})
})
