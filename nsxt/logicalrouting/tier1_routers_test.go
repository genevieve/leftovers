package logicalrouting_test

import (
	"context"
	"errors"

	"github.com/genevieve/leftovers/nsxt/logicalrouting"
	"github.com/genevieve/leftovers/nsxt/logicalrouting/fakes"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/vmware/go-vmware-nsxt/manager"
)

var _ = Describe("Tier 1 Routers", func() {
	var (
		client       *fakes.LogicalRoutingAPI
		logger       *fakes.Logger
		ctx          context.Context
		tier1Routers logicalrouting.Tier1Routers
	)

	BeforeEach(func() {
		client = &fakes.LogicalRoutingAPI{}
		logger = &fakes.Logger{}

		ctx = context.WithValue(context.Background(), "fruit", "soursop")

		logger.PromptWithDetailsCall.Returns.Proceed = true

		tier1Routers = logicalrouting.NewTier1Routers(client, ctx, logger)
	})

	Describe("List", func() {
		var filter string

		BeforeEach(func() {
			client.ListLogicalRoutersCall.Returns.LogicalRouterListResult = manager.LogicalRouterListResult{
				Results: []manager.LogicalRouter{
					manager.LogicalRouter{
						Id:          "soursop-123",
						DisplayName: "soursop",
					},
					manager.LogicalRouter{
						Id:          "cherimoya-456",
						DisplayName: "cherimoya",
					},
				},
			}

			filter = "soursop"
		})

		It("lists, filters, and prompts for tier 1 routers to delete", func() {
			list, err := tier1Routers.List(filter)
			Expect(err).NotTo(HaveOccurred())

			Expect(client.ListLogicalRoutersCall.CallCount).To(Equal(1))
			Expect(client.ListLogicalRoutersCall.Receives.Ctx).To(Equal(ctx))
			Expect(client.ListLogicalRoutersCall.Receives.LocalVarOptionals).To(HaveKeyWithValue("routerType", "TIER1"))

			Expect(logger.PromptWithDetailsCall.CallCount).To(Equal(1))
			Expect(logger.PromptWithDetailsCall.Receives.ResourceType).To(Equal("Tier 1 Router"))
			Expect(logger.PromptWithDetailsCall.Receives.ResourceName).To(Equal("soursop"))

			Expect(list).To(HaveLen(1))
			Expect(list[0].Name()).NotTo(Equal("cherimoya"))
		})

		Context("when system owned resources appear in the list returned by the client", func() {
			BeforeEach(func() {
				client.ListLogicalRoutersCall.Returns.LogicalRouterListResult = manager.LogicalRouterListResult{
					Results: []manager.LogicalRouter{
						manager.LogicalRouter{
							Id:          "soursop-123",
							DisplayName: "soursop",
						},
						manager.LogicalRouter{
							Id:          "soursop-456",
							DisplayName: "soursop-system",
							SystemOwned: true,
						},
					},
				}
			})

			It("skips system owned resources", func() {
				list, err := tier1Routers.List(filter)
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PromptWithDetailsCall.CallCount).To(Equal(1))
				Expect(logger.PromptWithDetailsCall.Receives.ResourceType).To(Equal("Tier 1 Router"))
				Expect(logger.PromptWithDetailsCall.Receives.ResourceName).To(Equal("soursop"))

				Expect(list).To(HaveLen(1))
				Expect(list[0].Name()).NotTo(Equal("soursop-system"))
			})
		})

		Context("when the client fails to list logical routers", func() {
			BeforeEach(func() {
				client.ListLogicalRoutersCall.Returns.Error = errors.New("PC LOAD LETTER")
			})

			It("returns the error", func() {
				_, err := tier1Routers.List(filter)
				Expect(err).To(MatchError("List Tier 1 Routers: PC LOAD LETTER"))
			})
		})
	})
})
