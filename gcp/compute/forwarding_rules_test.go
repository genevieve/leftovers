package compute_test

import (
	"errors"

	"github.com/genevievelesperance/leftovers/gcp/compute"
	"github.com/genevievelesperance/leftovers/gcp/compute/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	gcpcompute "google.golang.org/api/compute/v1"
)

var _ = Describe("ForwardingRules", func() {
	var (
		client  *fakes.ForwardingRulesClient
		logger  *fakes.Logger
		regions map[string]string
		filter  string

		forwardingRules compute.ForwardingRules
	)

	BeforeEach(func() {
		client = &fakes.ForwardingRulesClient{}
		logger = &fakes.Logger{}
		regions = map[string]string{
			"https://region-1": "region-1",
		}
		filter = "grape"

		forwardingRules = compute.NewForwardingRules(client, logger, regions)
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			logger.PromptCall.Returns.Proceed = true
			client.ListForwardingRulesCall.Returns.Output = &gcpcompute.ForwardingRuleList{
				Items: []*gcpcompute.ForwardingRule{{
					Name:   "banana",
					Region: "https://region-1",
				}},
			}
		})

		It("deletes forwarding rules", func() {
			err := forwardingRules.Delete(filter)
			Expect(err).NotTo(HaveOccurred())

			Expect(client.ListForwardingRulesCall.CallCount).To(Equal(1))
			Expect(client.ListForwardingRulesCall.Receives.Region).To(Equal("region-1"))

			Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete forwarding rule banana?"))

			Expect(client.DeleteForwardingRuleCall.CallCount).To(Equal(1))
			Expect(client.DeleteForwardingRuleCall.Receives.Region).To(Equal("region-1"))
			Expect(client.DeleteForwardingRuleCall.Receives.ForwardingRule).To(Equal("banana"))

			Expect(logger.PrintfCall.Messages).To(Equal([]string{"SUCCESS deleting forwarding rule banana\n"}))
		})

		Context("when the client fails to list forwarding rules", func() {
			BeforeEach(func() {
				client.ListForwardingRulesCall.Returns.Error = errors.New("some error")
			})

			It("returns the error", func() {
				err := forwardingRules.Delete(filter)
				Expect(err).To(MatchError("Listing forwarding rules for region region-1: some error"))
			})
		})

		Context("when the client fails to delete the forwarding rule", func() {
			BeforeEach(func() {
				client.DeleteForwardingRuleCall.Returns.Error = errors.New("some error")
			})

			It("logs the error", func() {
				err := forwardingRules.Delete(filter)
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PrintfCall.Messages).To(Equal([]string{"ERROR deleting forwarding rule banana: some error\n"}))
			})
		})

		Context("when the user says no to the prompt", func() {
			BeforeEach(func() {
				logger.PromptCall.Returns.Proceed = false
			})

			It("does not delete the forwarding rule", func() {
				err := forwardingRules.Delete(filter)
				Expect(err).NotTo(HaveOccurred())

				Expect(client.DeleteForwardingRuleCall.CallCount).To(Equal(0))
			})
		})
	})
})
