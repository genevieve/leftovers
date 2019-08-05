package logicalrouting_test

import (
	"context"
	"errors"

	"github.com/genevieve/leftovers/nsxt/logicalrouting"
	"github.com/genevieve/leftovers/nsxt/logicalrouting/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vmware/go-vmware-nsxt/manager"
)

var _ = Describe("Tier 1 Router", func() {
	var (
		client *fakes.LogicalRoutingAndServicesAPI
		ctx    context.Context
		name   string
		id     string

		tier1Router logicalrouting.Tier1Router
	)

	BeforeEach(func() {
		client = &fakes.LogicalRoutingAndServicesAPI{}
		name = "ackee"
		id = "ackee-123"

		ctx = context.WithValue(context.Background(), "fruit", "ackee")

		client.ListLogicalRouterPortsCall.Returns.ListResult = manager.LogicalRouterPortListResult{
			Results: []manager.LogicalRouterPort{{Id: "grape"}},
		}

		tier1Router = logicalrouting.NewTier1Router(client, ctx, name, id)
	})

	Describe("Delete", func() {
		It("deletes the tier1 router and it's ports", func() {
			err := tier1Router.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(client.ListLogicalRouterPortsCall.CallCount).To(Equal(1))
			Expect(client.ListLogicalRouterPortsCall.Receives.Context).To(Equal(ctx))
			Expect(client.ListLogicalRouterPortsCall.Receives.LocalVarOptionals).To(HaveKeyWithValue("force", true))
			Expect(client.ListLogicalRouterPortsCall.Receives.LocalVarOptionals).To(HaveKeyWithValue("logicalRouterId", id))

			Expect(client.DeleteLogicalRouterPortCall.CallCount).To(Equal(1))
			Expect(client.DeleteLogicalRouterPortCall.Receives.ID).To(Equal("grape"))
			Expect(client.DeleteLogicalRouterPortCall.Receives.Context).To(Equal(ctx))
			Expect(client.DeleteLogicalRouterPortCall.Receives.LocalVarOptionals).To(HaveKeyWithValue("force", true))

			Expect(client.DeleteLogicalRouterCall.CallCount).To(Equal(1))
			Expect(client.DeleteLogicalRouterCall.Receives.ID).To(Equal(id))
			Expect(client.DeleteLogicalRouterCall.Receives.Context).To(Equal(ctx))
			Expect(client.DeleteLogicalRouterCall.Receives.LocalVarOptionals).To(HaveKeyWithValue("force", true))
		})

		Context("when the client fails to list the router ports", func() {
			BeforeEach(func() {
				client.ListLogicalRouterPortsCall.Returns.Error = errors.New("banana")
			})

			It("returns the error", func() {
				err := tier1Router.Delete()
				Expect(err).To(MatchError("List Logical Router Ports: banana"))
			})
		})

		Context("when the client fails to delete the router ports", func() {
			BeforeEach(func() {
				client.DeleteLogicalRouterPortCall.Returns.Error = errors.New("banana")
			})

			It("returns the error", func() {
				err := tier1Router.Delete()
				Expect(err).To(MatchError("Delete Logical Router Port: banana"))
			})
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
