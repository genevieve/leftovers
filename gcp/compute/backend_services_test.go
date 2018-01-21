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
		filter = "banana"

		backendServices = compute.NewBackendServices(client, logger)
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			logger.PromptCall.Returns.Proceed = true
			client.ListBackendServicesCall.Returns.Output = &gcpcompute.BackendServiceList{
				Items: []*gcpcompute.BackendService{{
					Name: "banana-backend-service",
				}},
			}
		})

		It("deletes backend services", func() {
			err := backendServices.Delete(filter)
			Expect(err).NotTo(HaveOccurred())

			Expect(client.ListBackendServicesCall.CallCount).To(Equal(1))

			Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete backend service banana-backend-service?"))

			Expect(client.DeleteBackendServiceCall.CallCount).To(Equal(1))
			Expect(client.DeleteBackendServiceCall.Receives.BackendService).To(Equal("banana-backend-service"))

			Expect(logger.PrintfCall.Messages).To(Equal([]string{"SUCCESS deleting backend service banana-backend-service\n"}))
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

		Context("when the backend service name does not contain the filter", func() {
			It("does not try to delete it", func() {
				err := backendServices.Delete("grape")
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PromptCall.CallCount).To(Equal(0))
				Expect(client.DeleteBackendServiceCall.CallCount).To(Equal(0))
			})
		})

		Context("when the client fails to delete the backend service", func() {
			BeforeEach(func() {
				client.DeleteBackendServiceCall.Returns.Error = errors.New("some error")
			})

			It("logs the error", func() {
				err := backendServices.Delete(filter)
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PrintfCall.Messages).To(Equal([]string{"ERROR deleting backend service banana-backend-service: some error\n"}))
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
