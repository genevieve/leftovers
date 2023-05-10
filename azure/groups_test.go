package azure_test

import (
	"errors"

	"github.com/genevieve/leftovers/azure"
	"github.com/genevieve/leftovers/azure/fakes"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Groups", func() {
	var (
		client *fakes.GroupsClient
		logger *fakes.Logger
		filter string

		groups azure.Groups
	)

	BeforeEach(func() {
		client = &fakes.GroupsClient{}
		logger = &fakes.Logger{}
		filter = "banana"

		groups = azure.NewGroups(client, logger)
	})

	Describe("List", func() {
		BeforeEach(func() {
			logger.PromptWithDetailsCall.Returns.Proceed = true
			client.ListGroupsCall.Returns.Groups = []string{"banana-group", "kiwi-group"}
		})

		It("returns a list of resource groups to delete", func() {
			items, err := groups.List(filter)
			Expect(err).NotTo(HaveOccurred())

			Expect(client.ListGroupsCall.CallCount).To(Equal(1))
			Expect(logger.PromptWithDetailsCall.CallCount).To(Equal(1))

			Expect(items).To(HaveLen(1))
		})

		Context("when client fails to list resource groups", func() {
			BeforeEach(func() {
				client.ListGroupsCall.Returns.Err = errors.New("some error")
			})

			It("returns the error", func() {
				_, err := groups.List(filter)
				Expect(err).To(MatchError("Listing Resource Groups: some error"))
			})
		})

		Context("when the user responds no to the prompt", func() {
			BeforeEach(func() {
				logger.PromptWithDetailsCall.Returns.Proceed = false
			})

			It("does not return it in the list", func() {
				items, err := groups.List(filter)
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PromptWithDetailsCall.Receives.ResourceType).To(Equal("Resource Group"))
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
