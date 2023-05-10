package groupingobjects_test

import (
	"context"
	"errors"

	"github.com/genevieve/leftovers/nsxt/groupingobjects"
	"github.com/genevieve/leftovers/nsxt/groupingobjects/fakes"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/vmware/go-vmware-nsxt/manager"
)

var _ = Describe("NS Groups", func() {
	var (
		client   *fakes.GroupingObjectsAPI
		logger   *fakes.Logger
		ctx      context.Context
		nsGroups groupingobjects.NSGroups
	)

	BeforeEach(func() {
		client = &fakes.GroupingObjectsAPI{}
		logger = &fakes.Logger{}

		ctx = context.WithValue(context.Background(), "fruit", "pineapple")

		logger.PromptWithDetailsCall.Returns.Proceed = true

		nsGroups = groupingobjects.NewNSGroups(client, ctx, logger)
	})

	Describe("List", func() {
		var filter string

		BeforeEach(func() {
			client.ListNSGroupsCall.Returns.NsGroupListResult = manager.NsGroupListResult{
				Results: []manager.NsGroup{
					manager.NsGroup{
						Id:          "pineapple-123",
						DisplayName: "pineapple",
					},
					manager.NsGroup{
						Id:          "cherimoya-456",
						DisplayName: "cherimoya",
					},
				},
			}

			filter = "pineapple"
		})

		It("lists, filters, and prompts for ns groups to delete", func() {
			list, err := nsGroups.List(filter)
			Expect(err).NotTo(HaveOccurred())

			Expect(client.ListNSGroupsCall.CallCount).To(Equal(1))
			Expect(client.ListNSGroupsCall.Receives.Context).To(Equal(ctx))

			Expect(logger.PromptWithDetailsCall.CallCount).To(Equal(1))
			Expect(logger.PromptWithDetailsCall.Receives.ResourceType).To(Equal("NS Group"))
			Expect(logger.PromptWithDetailsCall.Receives.ResourceName).To(Equal("pineapple"))

			Expect(list).To(HaveLen(1))
			Expect(list[0].Name()).NotTo(Equal("cherimoya"))
		})

		Context("when the client fails to list ns groups", func() {
			BeforeEach(func() {
				client.ListNSGroupsCall.Returns.Error = errors.New("PC LOAD LETTER")
			})

			It("returns the error", func() {
				_, err := nsGroups.List(filter)
				Expect(err).To(MatchError("List NS Groups: PC LOAD LETTER"))
			})
		})
	})
})
