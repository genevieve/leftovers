package compute_test

import (
	"errors"

	"github.com/genevievelesperance/leftovers/gcp/compute"
	"github.com/genevievelesperance/leftovers/gcp/compute/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	gcpcompute "google.golang.org/api/compute/v1"
)

var _ = Describe("Addresses", func() {
	var (
		client  *fakes.AddressesClient
		logger  *fakes.Logger
		regions map[string]string

		addresses compute.Addresses
	)

	BeforeEach(func() {
		client = &fakes.AddressesClient{}
		logger = &fakes.Logger{}
		regions = map[string]string{
			"https://region-1": "region-1",
		}

		addresses = compute.NewAddresses(client, logger, regions)
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			logger.PromptCall.Returns.Proceed = true
			client.ListAddressesCall.Returns.Output = &gcpcompute.AddressList{
				Items: []*gcpcompute.Address{{
					Name:   "banana",
					Region: "https://region-1",
				}},
			}
		})

		It("deletes addresses", func() {
			err := addresses.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(client.ListAddressesCall.CallCount).To(Equal(1))
			Expect(client.ListAddressesCall.Receives.Region).To(Equal("region-1"))

			Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete address banana?"))

			Expect(client.DeleteAddressCall.CallCount).To(Equal(1))
			Expect(client.DeleteAddressCall.Receives.Region).To(Equal("region-1"))
			Expect(client.DeleteAddressCall.Receives.Address).To(Equal("banana"))

			Expect(logger.PrintfCall.Messages).To(Equal([]string{"SUCCESS deleting address banana\n"}))
		})

		Context("when the client fails to list addresses", func() {
			BeforeEach(func() {
				client.ListAddressesCall.Returns.Error = errors.New("some error")
			})

			It("returns the error", func() {
				err := addresses.Delete()
				Expect(err).To(MatchError("Listing addresses for region region-1: some error"))
			})
		})

		Context("when the client fails to delete the address", func() {
			BeforeEach(func() {
				client.DeleteAddressCall.Returns.Error = errors.New("some error")
			})

			It("logs the error", func() {
				err := addresses.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PrintfCall.Messages).To(Equal([]string{"ERROR deleting address banana: some error\n"}))
			})
		})

		Context("when the user says no to the prompt", func() {
			BeforeEach(func() {
				logger.PromptCall.Returns.Proceed = false
			})

			It("does not delete the address", func() {
				err := addresses.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(client.DeleteAddressCall.CallCount).To(Equal(0))
			})
		})
	})
})
