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

var _ = Describe("NS Services", func() {
	var (
		client     *fakes.GroupingObjectsAPI
		logger     *fakes.Logger
		ctx        context.Context
		nsServices groupingobjects.NSServices
	)

	BeforeEach(func() {
		client = &fakes.GroupingObjectsAPI{}
		logger = &fakes.Logger{}

		ctx = context.WithValue(context.Background(), "fruit", "pineapple")

		logger.PromptWithDetailsCall.Returns.Proceed = true

		nsServices = groupingobjects.NewNSServices(client, ctx, logger)
	})

	Describe("List", func() {
		var filter string

		BeforeEach(func() {
			client.ListNSServicesCall.Returns.NsServiceListResult = manager.NsServiceListResult{
				Results: []manager.NsService{
					manager.NsService{
						Id:          "pineapple-123",
						DisplayName: "pineapple",
					},
					manager.NsService{
						Id:          "cherimoya-456",
						DisplayName: "cherimoya",
					},
				},
			}

			filter = "pineapple"
		})

		It("lists, filters, and prompts for ns services to delete", func() {
			list, err := nsServices.List(filter)
			Expect(err).NotTo(HaveOccurred())

			Expect(client.ListNSServicesCall.CallCount).To(Equal(1))
			Expect(client.ListNSServicesCall.Receives.Context).To(Equal(ctx))

			Expect(logger.PromptWithDetailsCall.CallCount).To(Equal(1))
			Expect(logger.PromptWithDetailsCall.Receives.ResourceType).To(Equal("NS Service"))
			Expect(logger.PromptWithDetailsCall.Receives.ResourceName).To(Equal("pineapple"))

			Expect(list).To(HaveLen(1))
			Expect(list[0].Name()).NotTo(Equal("cherimoya"))
		})

		Context("when the client fails to list ns services", func() {
			BeforeEach(func() {
				client.ListNSServicesCall.Returns.Error = errors.New("PC LOAD LETTER")
			})

			It("returns the error", func() {
				_, err := nsServices.List(filter)
				Expect(err).To(MatchError("List NS Services: PC LOAD LETTER"))
			})
		})
	})
})
