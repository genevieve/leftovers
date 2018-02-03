package compute_test

import (
	"errors"

	"github.com/genevieve/leftovers/gcp/compute"
	"github.com/genevieve/leftovers/gcp/compute/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	gcpcompute "google.golang.org/api/compute/v1"
)

var _ = Describe("GlobalAddresses", func() {
	var (
		client *fakes.GlobalAddressesClient
		logger *fakes.Logger

		addresses compute.GlobalAddresses
	)

	BeforeEach(func() {
		client = &fakes.GlobalAddressesClient{}
		logger = &fakes.Logger{}

		logger.PromptCall.Returns.Proceed = true

		addresses = compute.NewGlobalAddresses(client, logger)
	})

	Describe("List", func() {
		var filter string

		BeforeEach(func() {
			client.ListGlobalAddressesCall.Returns.Output = &gcpcompute.AddressList{
				Items: []*gcpcompute.Address{{
					Name: "banana-address",
				}},
			}
			filter = "banana"
		})

		It("lists, filters, and prompts for addresses to delete", func() {
			list, err := addresses.List(filter)
			Expect(err).NotTo(HaveOccurred())

			Expect(client.ListGlobalAddressesCall.CallCount).To(Equal(1))

			Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete global address banana-address?"))

			Expect(list).To(HaveLen(1))
			Expect(list).To(HaveKeyWithValue("banana-address", ""))
		})

		Context("when the client fails to list addresses", func() {
			BeforeEach(func() {
				client.ListGlobalAddressesCall.Returns.Error = errors.New("some error")
			})

			It("returns the error", func() {
				_, err := addresses.List(filter)
				Expect(err).To(MatchError("Listing global addresses: some error"))
			})
		})

		Context("when the address name does not contain the filter", func() {
			It("does not add it to the list", func() {
				list, err := addresses.List("grape")
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PromptCall.CallCount).To(Equal(0))
				Expect(list).To(HaveLen(0))
			})
		})

		Context("when the address is in use", func() {
			BeforeEach(func() {
				client.ListGlobalAddressesCall.Returns.Output = &gcpcompute.AddressList{
					Items: []*gcpcompute.Address{{
						Name:  "banana-address",
						Users: []string{"a-virtual-machine"},
					}},
				}
			})

			It("does not add it to the list", func() {
				list, err := addresses.List(filter)
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PromptCall.CallCount).To(Equal(0))
				Expect(list).To(HaveLen(0))
			})
		})

		Context("when the user says no to the prompt", func() {
			BeforeEach(func() {
				logger.PromptCall.Returns.Proceed = false
			})

			It("does not add it to the list", func() {
				list, err := addresses.List(filter)
				Expect(err).NotTo(HaveOccurred())

				Expect(list).To(HaveLen(0))
			})
		})
	})

	Describe("Delete", func() {
		var list map[string]string

		BeforeEach(func() {
			list = map[string]string{"banana-address": ""}
		})

		It("deletes addresses", func() {
			addresses.Delete(list)

			Expect(client.DeleteGlobalAddressCall.CallCount).To(Equal(1))
			Expect(client.DeleteGlobalAddressCall.Receives.Address).To(Equal("banana-address"))

			Expect(logger.PrintfCall.Messages).To(Equal([]string{"SUCCESS deleting global address banana-address\n"}))
		})

		Context("when the client fails to delete the address", func() {
			BeforeEach(func() {
				client.DeleteGlobalAddressCall.Returns.Error = errors.New("some error")
			})

			It("logs the error", func() {
				addresses.Delete(list)

				Expect(logger.PrintfCall.Messages).To(Equal([]string{"ERROR deleting global address banana-address: some error\n"}))
			})
		})
	})
})
