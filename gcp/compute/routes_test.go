package compute_test

import (
	"errors"

	"github.com/genevieve/leftovers/gcp/compute"
	"github.com/genevieve/leftovers/gcp/compute/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	gcpcompute "google.golang.org/api/compute/v1"
)

var _ = Describe("Routes", func() {
	var (
		client *fakes.RoutesClient
		logger *fakes.Logger

		routes compute.Routes
	)

	BeforeEach(func() {
		client = &fakes.RoutesClient{}
		logger = &fakes.Logger{}

		routes = compute.NewRoutes(client, logger)
	})

	Describe("List", func() {
		var filter string

		BeforeEach(func() {
			logger.PromptWithDetailsCall.Returns.Proceed = true
			client.ListRoutesCall.Returns.RouteSlice = []*gcpcompute.Route{{
				Name: "banana-route",
			}}
			filter = "banana"
		})

		It("lists, filters, and prompts for routes to delete", func() {
			list, err := routes.List(filter, false)
			Expect(err).NotTo(HaveOccurred())

			Expect(client.ListRoutesCall.CallCount).To(Equal(1))

			Expect(logger.PromptWithDetailsCall.CallCount).To(Equal(1))
			Expect(logger.PromptWithDetailsCall.Receives.ResourceType).To(Equal("Route"))
			Expect(logger.PromptWithDetailsCall.Receives.ResourceName).To(Equal("banana-route"))

			Expect(list).To(HaveLen(1))
		})

		Context("when the route name does not contain the filter, but the network does", func() {
			BeforeEach(func() {
				client.ListRoutesCall.Returns.RouteSlice = []*gcpcompute.Route{{
					Name:    "banana-route",
					Network: ".com/networks/kiwi-network",
				}}
				client.GetNetworkNameCall.Returns.Name = "kiwi-network"
			})
			It("returns the route in the list to delete", func() {
				list, err := routes.List("kiwi", false)
				Expect(err).NotTo(HaveOccurred())

				Expect(list).To(HaveLen(1))
			})
		})

		Context("when the client fails to list routes", func() {
			BeforeEach(func() {
				client.ListRoutesCall.Returns.Error = errors.New("some error")
			})

			It("returns the error", func() {
				_, err := routes.List(filter, false)
				Expect(err).To(MatchError("List Routes: some error"))
			})
		})

		Context("when the route name contains the word 'default'", func() {
			BeforeEach(func() {
				client.ListRoutesCall.Returns.RouteSlice = []*gcpcompute.Route{{
					Name: "default-route",
				}}
			})

			It("does not add it to the list", func() {
				list, err := routes.List("", false)
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PromptWithDetailsCall.CallCount).To(Equal(0))
				Expect(list).To(HaveLen(0))
			})
		})

		Context("when the route name does not contain the filter", func() {
			It("does not add it to the list", func() {
				list, err := routes.List("grape", false)
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PromptWithDetailsCall.CallCount).To(Equal(0))
				Expect(list).To(HaveLen(0))
			})
		})

		Context("when the user says no to the prompt", func() {
			BeforeEach(func() {
				logger.PromptWithDetailsCall.Returns.Proceed = false
			})

			It("does not add it to the list", func() {
				list, err := routes.List((filter), false)
				Expect(err).NotTo(HaveOccurred())

				Expect(list).To(HaveLen(0))
			})
		})
	})
})
