package azure_test

import (
	"errors"

	"github.com/genevieve/leftovers/azure"
	"github.com/genevieve/leftovers/azure/fakes"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("AppSecurityGroups", func() {
	var (
		client *fakes.AppSecurityGroupsClient
		logger *fakes.Logger
		filter string
		rgName string

		groups azure.AppSecurityGroups
	)

	BeforeEach(func() {
		client = &fakes.AppSecurityGroupsClient{}
		logger = &fakes.Logger{}
		filter = "banana"
		rgName = "resource-group"

		groups = azure.NewAppSecurityGroups(client, rgName, logger)
	})

	Describe("List", func() {
		BeforeEach(func() {
			logger.PromptWithDetailsCall.Returns.Proceed = true
			client.ListAppSecurityGroupsCall.Returns.StringSlice = []string{"banana-group", "kiwi-group"}
		})

		It("returns a list of resources to delete", func() {
			items, err := groups.List(filter)
			Expect(err).NotTo(HaveOccurred())

			Expect(client.ListAppSecurityGroupsCall.CallCount).To(Equal(1))
			Expect(client.ListAppSecurityGroupsCall.Receives.RgName).To(Equal(rgName))

			Expect(logger.PromptWithDetailsCall.CallCount).To(Equal(1))

			Expect(items).To(HaveLen(1))
		})

		Context("when client fails to list the resource", func() {
			BeforeEach(func() {
				client.ListAppSecurityGroupsCall.Returns.Error = errors.New("some error")
			})

			It("returns the error", func() {
				_, err := groups.List(filter)
				Expect(err).To(MatchError("Listing Application Security Groups: some error"))
			})
		})

		Context("when the user responds no to the prompt", func() {
			BeforeEach(func() {
				logger.PromptWithDetailsCall.Returns.Proceed = false
			})

			It("does not return it in the list", func() {
				items, err := groups.List(filter)
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PromptWithDetailsCall.Receives.ResourceType).To(Equal("Application Security Group"))
				Expect(logger.PromptWithDetailsCall.Receives.ResourceName).To(Equal("banana-group"))

				Expect(items).To(HaveLen(0))
			})
		})

		Context("when the resource group name does not contain the filter", func() {
			It("does not return it in the list", func() {
				items, err := groups.List("grape")
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PromptWithDetailsCall.CallCount).To(Equal(0))
				Expect(items).To(HaveLen(0))
			})
		})
	})
})
