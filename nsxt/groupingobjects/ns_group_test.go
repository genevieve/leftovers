package groupingobjects_test

import (
	"context"
	"errors"

	"github.com/genevieve/leftovers/nsxt/groupingobjects"
	"github.com/genevieve/leftovers/nsxt/groupingobjects/fakes"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("NS Group", func() {
	var (
		client *fakes.GroupingObjectsAPI
		ctx    context.Context
		name   string
		id     string

		nsGroup groupingobjects.NSGroup
	)

	BeforeEach(func() {
		client = &fakes.GroupingObjectsAPI{}
		name = "mango"
		id = "mango-123"

		ctx = context.WithValue(context.Background(), "fruit", "mango")

		nsGroup = groupingobjects.NewNSGroup(client, ctx, name, id)
	})

	Describe("Delete", func() {
		It("deletes the ns group", func() {
			err := nsGroup.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(client.DeleteNSGroupCall.CallCount).To(Equal(1))
			Expect(client.DeleteNSGroupCall.Receives.String).To(Equal(id))
			Expect(client.DeleteNSGroupCall.Receives.Context).To(Equal(ctx))
		})

		Context("when the client fails to delete the ns group", func() {
			BeforeEach(func() {
				client.DeleteNSGroupCall.Returns.Error = errors.New("insufficient funds")
			})

			It("returns the error", func() {
				err := nsGroup.Delete()
				Expect(err).To(MatchError("Delete: insufficient funds"))
			})
		})
	})

	Describe("Name", func() {
		It("returns the name", func() {
			Expect(nsGroup.Name()).To(Equal(name))
		})
	})

	Describe("Type", func() {
		It("returns the type", func() {
			Expect(nsGroup.Type()).To(Equal("NS Group"))
		})
	})
})
