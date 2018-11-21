package compute_test

import (
	"errors"

	"github.com/genevieve/leftovers/openstack/compute"
	"github.com/genevieve/leftovers/openstack/compute/fakes"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/floatingips"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("FloatingIPs", func() {
	var (
		client *fakes.FloatingIPsClient
		logger *fakes.Logger

		floatingIPs compute.FloatingIPs
	)

	BeforeEach(func() {
		client = &fakes.FloatingIPsClient{}
		logger = &fakes.Logger{}

		logger.PromptWithDetailsCall.Returns.Proceed = true

		floatingIPs = compute.NewFloatingIPs(client, logger)
	})

	Describe("List", func() {
		var filter string

		BeforeEach(func() {
			client.ListFloatingIPsCall.Returns.Output = []floatingips.FloatingIP{{
				IP:   "the-ip",
				Pool: "the-pool",
			}}
			filter = "banana"
		})

		It("lists and prompts for floating ips to delete", func() {
			list, err := floatingIPs.List(filter)
			Expect(err).NotTo(HaveOccurred())

			Expect(client.ListFloatingIPsCall.CallCount).To(Equal(1))

			Expect(logger.PromptWithDetailsCall.CallCount).To(Equal(1))
			Expect(logger.PromptWithDetailsCall.Receives.Type).To(Equal("Floating IP"))
			Expect(logger.PromptWithDetailsCall.Receives.Name).To(Equal("the-ip"))

			Expect(list).To(HaveLen(1))
		})

		Context("when the client fails to list floating ips", func() {
			BeforeEach(func() {
				client.ListFloatingIPsCall.Returns.Error = errors.New("some error")
			})

			It("returns the error", func() {
				_, err := floatingIPs.List(filter)
				Expect(err).To(MatchError("List Floating IPs: some error"))
			})
		})

		Context("when the user says no to the prompt", func() {
			BeforeEach(func() {
				logger.PromptWithDetailsCall.Returns.Proceed = false
			})

			It("does not add it to the list", func() {
				list, err := floatingIPs.List(filter)
				Expect(err).NotTo(HaveOccurred())

				Expect(list).To(HaveLen(0))
			})
		})
	})
})
