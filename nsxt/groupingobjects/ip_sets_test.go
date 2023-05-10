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

var _ = Describe("IP Sets", func() {
	var (
		client *fakes.GroupingObjectsAPI
		logger *fakes.Logger
		ctx    context.Context
		ipSets groupingobjects.IPSets
	)

	BeforeEach(func() {
		client = &fakes.GroupingObjectsAPI{}
		logger = &fakes.Logger{}

		ctx = context.WithValue(context.Background(), "fruit", "pineapple")

		logger.PromptWithDetailsCall.Returns.Proceed = true

		ipSets = groupingobjects.NewIPSets(client, ctx, logger)
	})

	Describe("List", func() {
		var filter string

		BeforeEach(func() {
			client.ListIPSetsCall.Returns.IpSetListResult = manager.IpSetListResult{
				Results: []manager.IpSet{
					manager.IpSet{
						Id:          "pineapple-123",
						DisplayName: "pineapple",
					},
					manager.IpSet{
						Id:          "cherimoya-456",
						DisplayName: "cherimoya",
					},
				},
			}

			filter = "pineapple"
		})

		It("lists, filters, and prompts for ip sets to delete", func() {
			list, err := ipSets.List(filter)
			Expect(err).NotTo(HaveOccurred())

			Expect(client.ListIPSetsCall.CallCount).To(Equal(1))
			Expect(client.ListIPSetsCall.Receives.Context).To(Equal(ctx))

			Expect(logger.PromptWithDetailsCall.CallCount).To(Equal(1))
			Expect(logger.PromptWithDetailsCall.Receives.ResourceType).To(Equal("IP Set"))
			Expect(logger.PromptWithDetailsCall.Receives.ResourceName).To(Equal("pineapple"))

			Expect(list).To(HaveLen(1))
			Expect(list[0].Name()).NotTo(Equal("cherimoya"))
		})

		Context("when the client fails to list ip sets", func() {
			BeforeEach(func() {
				client.ListIPSetsCall.Returns.Error = errors.New("PC LOAD LETTER")
			})

			It("returns the error", func() {
				_, err := ipSets.List(filter)
				Expect(err).To(MatchError("List IP Sets: PC LOAD LETTER"))
			})
		})
	})
})
