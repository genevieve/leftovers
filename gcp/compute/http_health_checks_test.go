package compute_test

import (
	"errors"

	"github.com/genevievelesperance/leftovers/gcp/compute"
	"github.com/genevievelesperance/leftovers/gcp/compute/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	gcpcompute "google.golang.org/api/compute/v1"
)

var _ = Describe("HttpHealthChecks", func() {
	var (
		client *fakes.HttpHealthChecksClient
		logger *fakes.Logger

		httpHealthChecks compute.HttpHealthChecks
	)

	BeforeEach(func() {
		client = &fakes.HttpHealthChecksClient{}
		logger = &fakes.Logger{}

		httpHealthChecks = compute.NewHttpHealthChecks(client, logger)
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			logger.PromptCall.Returns.Proceed = true
			client.ListHttpHealthChecksCall.Returns.Output = &gcpcompute.HttpHealthCheckList{
				Items: []*gcpcompute.HttpHealthCheck{{
					Name: "banana",
				}},
			}
		})

		It("deletes http health checks", func() {
			err := httpHealthChecks.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(client.ListHttpHealthChecksCall.CallCount).To(Equal(1))

			Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete http health check banana?"))

			Expect(client.DeleteHttpHealthCheckCall.CallCount).To(Equal(1))
			Expect(client.DeleteHttpHealthCheckCall.Receives.HttpHealthCheck).To(Equal("banana"))

			Expect(logger.PrintfCall.Messages).To(Equal([]string{"SUCCESS deleting http health check banana\n"}))
		})

		Context("when the client fails to list http health checks", func() {
			BeforeEach(func() {
				client.ListHttpHealthChecksCall.Returns.Error = errors.New("some error")
			})

			It("returns the error", func() {
				err := httpHealthChecks.Delete()
				Expect(err).To(MatchError("Listing http health checks: some error"))
			})
		})

		Context("when the client fails to delete the http health check", func() {
			BeforeEach(func() {
				client.DeleteHttpHealthCheckCall.Returns.Error = errors.New("some error")
			})

			It("logs the error", func() {
				err := httpHealthChecks.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PrintfCall.Messages).To(Equal([]string{"ERROR deleting http health check banana: some error\n"}))
			})
		})

		Context("when the user says no to the prompt", func() {
			BeforeEach(func() {
				logger.PromptCall.Returns.Proceed = false
			})

			It("does not delete the http health check", func() {
				err := httpHealthChecks.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(client.DeleteHttpHealthCheckCall.CallCount).To(Equal(0))
			})
		})
	})
})
