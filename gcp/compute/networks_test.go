package compute_test

import (
	"errors"

	"github.com/genevievelesperance/leftovers/gcp/compute"
	"github.com/genevievelesperance/leftovers/gcp/compute/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	gcpcompute "google.golang.org/api/compute/v1"
)

var _ = Describe("Networks", func() {
	var (
		client *fakes.NetworksClient
		logger *fakes.Logger

		networks compute.Networks
	)

	BeforeEach(func() {
		client = &fakes.NetworksClient{}
		logger = &fakes.Logger{}

		networks = compute.NewNetworks(client, logger)
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			logger.PromptCall.Returns.Proceed = true
			client.ListNetworksCall.Returns.Output = &gcpcompute.NetworkList{
				Items: []*gcpcompute.Network{{
					Name: "banana",
				}},
			}
		})

		It("deletes networks", func() {
			err := networks.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(client.ListNetworksCall.CallCount).To(Equal(1))

			Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete network banana?"))

			Expect(client.DeleteNetworkCall.CallCount).To(Equal(1))
			Expect(client.DeleteNetworkCall.Receives.Network).To(Equal("banana"))

			Expect(logger.PrintfCall.Messages).To(Equal([]string{"SUCCESS deleting network banana\n"}))
		})

		Context("when the client fails to list networks", func() {
			BeforeEach(func() {
				client.ListNetworksCall.Returns.Error = errors.New("some error")
			})

			It("returns the error", func() {
				err := networks.Delete()
				Expect(err).To(MatchError("Listing networks: some error"))
			})
		})

		Context("when the client fails to delete the network", func() {
			BeforeEach(func() {
				client.DeleteNetworkCall.Returns.Error = errors.New("some error")
			})

			It("logs the error", func() {
				err := networks.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PrintfCall.Messages).To(Equal([]string{"ERROR deleting network banana: some error\n"}))
			})
		})

		Context("when the user says no to the prompt", func() {
			BeforeEach(func() {
				logger.PromptCall.Returns.Proceed = false
			})

			It("does not delete the network", func() {
				err := networks.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(client.DeleteNetworkCall.CallCount).To(Equal(0))
			})
		})
	})
})
