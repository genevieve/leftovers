package compute_test

import (
	"errors"

	"github.com/genevievelesperance/leftovers/gcp/compute"
	"github.com/genevievelesperance/leftovers/gcp/compute/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	gcpcompute "google.golang.org/api/compute/v1"
)

var _ = Describe("TargetPools", func() {
	var (
		client  *fakes.TargetPoolsClient
		logger  *fakes.Logger
		regions map[string]string

		targetPools compute.TargetPools
	)

	BeforeEach(func() {
		client = &fakes.TargetPoolsClient{}
		logger = &fakes.Logger{}
		regions = map[string]string{
			"https://region-1": "region-1",
		}

		targetPools = compute.NewTargetPools(client, logger, regions)
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			logger.PromptCall.Returns.Proceed = true
			client.ListTargetPoolsCall.Returns.Output = &gcpcompute.TargetPoolList{
				Items: []*gcpcompute.TargetPool{{
					Name:   "banana",
					Region: "https://region-1",
				}},
			}
		})

		It("deletes target pools", func() {
			err := targetPools.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(client.ListTargetPoolsCall.CallCount).To(Equal(1))
			Expect(client.ListTargetPoolsCall.Receives.Region).To(Equal("region-1"))

			Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete target pool banana?"))

			Expect(client.DeleteTargetPoolCall.CallCount).To(Equal(1))
			Expect(client.DeleteTargetPoolCall.Receives.Region).To(Equal("region-1"))
			Expect(client.DeleteTargetPoolCall.Receives.TargetPool).To(Equal("banana"))

			Expect(logger.PrintfCall.Messages).To(Equal([]string{"SUCCESS deleting target pool banana\n"}))
		})

		Context("when the client fails to list target pools", func() {
			BeforeEach(func() {
				client.ListTargetPoolsCall.Returns.Error = errors.New("some error")
			})

			It("returns the error", func() {
				err := targetPools.Delete()
				Expect(err).To(MatchError("Listing target pools for region region-1: some error"))
			})
		})

		Context("when the client fails to delete the target pool", func() {
			BeforeEach(func() {
				client.DeleteTargetPoolCall.Returns.Error = errors.New("some error")
			})

			It("logs the error", func() {
				err := targetPools.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PrintfCall.Messages).To(Equal([]string{"ERROR deleting target pool banana: some error\n"}))
			})
		})

		Context("when the user says no to the prompt", func() {
			BeforeEach(func() {
				logger.PromptCall.Returns.Proceed = false
			})

			It("does not delete the target pool", func() {
				err := targetPools.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(client.DeleteTargetPoolCall.CallCount).To(Equal(0))
			})
		})
	})
})
