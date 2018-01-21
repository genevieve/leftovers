package compute_test

import (
	"errors"

	"github.com/genevievelesperance/leftovers/gcp/compute"
	"github.com/genevievelesperance/leftovers/gcp/compute/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	gcpcompute "google.golang.org/api/compute/v1"
)

var _ = Describe("BackendServices", func() {
	var (
		client *fakes.BackendServicesClient
		logger *fakes.Logger
		filter string

		backendServices compute.BackendServices
	)

	BeforeEach(func() {
		client = &fakes.BackendServicesClient{}
		logger = &fakes.Logger{}
		filter = "grape"

		backendServices = compute.NewBackendServices(client, logger)
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			logger.PromptCall.Returns.Proceed = true
			client.ListBackendServicesCall.Returns.Output = &gcpcompute.BackendServiceList{
				Items: []*gcpcompute.BackendService{{
					Name: "banana",
				}},
			}
		})

		It("deletes backend services", func() {
			err := backendServices.Delete(filter)
			Expect(err).NotTo(HaveOccurred())

			Expect(client.ListBackendServicesCall.CallCount).To(Equal(1))

			Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete backend service banana?"))

			Expect(client.DeleteBackendServiceCall.CallCount).To(Equal(1))
			Expect(client.DeleteBackendServiceCall.Receives.BackendService).To(Equal("banana"))

			Expect(logger.PrintfCall.Messages).To(Equal([]string{"SUCCESS deleting backend service banana\n"}))
		})

		Context("when the client fails to list backend services", func() {
			BeforeEach(func() {
				client.ListBackendServicesCall.Returns.Error = errors.New("some error")
			})

			It("returns the error", func() {
				err := backendServices.Delete(filter)
				Expect(err).To(MatchError("Listing backend services: some error"))
			})
		})

		Context("when the client fails to delete the backend service", func() {
			BeforeEach(func() {
				client.DeleteBackendServiceCall.Returns.Error = errors.New("some error")
			})

			It("logs the error", func() {
				err := backendServices.Delete(filter)
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PrintfCall.Messages).To(Equal([]string{"ERROR deleting backend service banana: some error\n"}))
			})
		})

		Context("when the user says no to the prompt", func() {
			BeforeEach(func() {
				logger.PromptCall.Returns.Proceed = false
			})

			It("does not delete the backend service", func() {
				err := backendServices.Delete(filter)
				Expect(err).NotTo(HaveOccurred())

				Expect(client.DeleteBackendServiceCall.CallCount).To(Equal(0))
			})
		})
	})
})
