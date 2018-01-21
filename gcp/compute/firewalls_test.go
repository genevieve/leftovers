package compute_test

import (
	"errors"

	"github.com/genevievelesperance/leftovers/gcp/compute"
	"github.com/genevievelesperance/leftovers/gcp/compute/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	gcpcompute "google.golang.org/api/compute/v1"
)

var _ = Describe("Firewalls", func() {
	var (
		client *fakes.FirewallsClient
		logger *fakes.Logger
		filter string

		firewalls compute.Firewalls
	)

	BeforeEach(func() {
		client = &fakes.FirewallsClient{}
		logger = &fakes.Logger{}
		filter = "banana"

		firewalls = compute.NewFirewalls(client, logger)
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			logger.PromptCall.Returns.Proceed = true
			client.ListFirewallsCall.Returns.Output = &gcpcompute.FirewallList{
				Items: []*gcpcompute.Firewall{{
					Name: "banana-firewall",
				}},
			}
		})

		It("deletes firewalls", func() {
			err := firewalls.Delete(filter)
			Expect(err).NotTo(HaveOccurred())

			Expect(client.ListFirewallsCall.CallCount).To(Equal(1))

			Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete firewall banana-firewall?"))

			Expect(client.DeleteFirewallCall.CallCount).To(Equal(1))
			Expect(client.DeleteFirewallCall.Receives.Firewall).To(Equal("banana-firewall"))

			Expect(logger.PrintfCall.Messages).To(Equal([]string{"SUCCESS deleting firewall banana-firewall\n"}))
		})

		Context("when the client fails to list firewalls", func() {
			BeforeEach(func() {
				client.ListFirewallsCall.Returns.Error = errors.New("some error")
			})

			It("returns the error", func() {
				err := firewalls.Delete(filter)
				Expect(err).To(MatchError("Listing firewalls: some error"))
			})
		})

		Context("when the firewall name does not contain the filter", func() {
			It("does not delete it", func() {
				err := firewalls.Delete("grape")
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PromptCall.CallCount).To(Equal(0))
				Expect(client.DeleteFirewallCall.CallCount).To(Equal(0))
			})
		})

		Context("when the client fails to delete the firewall", func() {
			BeforeEach(func() {
				client.DeleteFirewallCall.Returns.Error = errors.New("some error")
			})

			It("logs the error", func() {
				err := firewalls.Delete(filter)
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PrintfCall.Messages).To(Equal([]string{"ERROR deleting firewall banana-firewall: some error\n"}))
			})
		})

		Context("when the user says no to the prompt", func() {
			BeforeEach(func() {
				logger.PromptCall.Returns.Proceed = false
			})

			It("does not delete it", func() {
				err := firewalls.Delete(filter)
				Expect(err).NotTo(HaveOccurred())

				Expect(client.DeleteFirewallCall.CallCount).To(Equal(0))
			})
		})
	})
})
