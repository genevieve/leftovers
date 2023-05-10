package logicalrouting_test

import (
	"context"
	"errors"

	"github.com/genevieve/leftovers/nsxt/logicalrouting"
	"github.com/genevieve/leftovers/nsxt/logicalrouting/fakes"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tier 1 Router", func() {
	var (
		client *fakes.LogicalRoutingAPI
		ctx    context.Context
		name   string
		id     string

		tier1Router logicalrouting.Tier1Router
	)

	BeforeEach(func() {
		client = &fakes.LogicalRoutingAPI{}
		name = "ackee"
		id = "ackee-123"

		ctx = context.WithValue(context.Background(), "fruit", "ackee")

		tier1Router = logicalrouting.NewTier1Router(client, ctx, name, id)
	})

	Describe("Delete", func() {
		It("deletes the tier1 router", func() {
			err := tier1Router.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(client.DeleteLogicalRouterCall.CallCount).To(Equal(1))
			Expect(client.DeleteLogicalRouterCall.Receives.Id).To(Equal(id))
			Expect(client.DeleteLogicalRouterCall.Receives.Ctx).To(Equal(ctx))
			Expect(client.DeleteLogicalRouterCall.Receives.LocalVarOptionals).To(HaveKeyWithValue("force", true))
		})

		Context("when the client fails to delete the router", func() {
			BeforeEach(func() {
				client.DeleteLogicalRouterCall.Returns.Error = errors.New("insufficient funds")
			})

			It("returns the error", func() {
				err := tier1Router.Delete()
				Expect(err).To(MatchError("Delete: insufficient funds"))
			})
		})
	})

	Describe("Name", func() {
		It("returns the name", func() {
			Expect(tier1Router.Name()).To(Equal(name))
		})
	})

	Describe("Type", func() {
		It("returns the type", func() {
			Expect(tier1Router.Type()).To(Equal("Tier 1 Router"))
		})
	})
})
