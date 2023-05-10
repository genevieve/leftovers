package groupingobjects_test

import (
	"context"
	"errors"

	"github.com/genevieve/leftovers/nsxt/groupingobjects"
	"github.com/genevieve/leftovers/nsxt/groupingobjects/fakes"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("IP Set", func() {
	var (
		client *fakes.GroupingObjectsAPI
		ctx    context.Context
		name   string
		id     string

		ipSet groupingobjects.IPSet
	)

	BeforeEach(func() {
		client = &fakes.GroupingObjectsAPI{}
		name = "mango"
		id = "mango-123"

		ctx = context.WithValue(context.Background(), "fruit", "mango")

		ipSet = groupingobjects.NewIPSet(client, ctx, name, id)
	})

	Describe("Delete", func() {
		It("deletes the ip set", func() {
			err := ipSet.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(client.DeleteIPSetCall.CallCount).To(Equal(1))
			Expect(client.DeleteIPSetCall.Receives.String).To(Equal(id))
			Expect(client.DeleteIPSetCall.Receives.Context).To(Equal(ctx))
		})

		Context("when the client fails to delete the ip set", func() {
			BeforeEach(func() {
				client.DeleteIPSetCall.Returns.Error = errors.New("insufficient funds")
			})

			It("returns the error", func() {
				err := ipSet.Delete()
				Expect(err).To(MatchError("Delete: insufficient funds"))
			})
		})
	})

	Describe("Name", func() {
		It("returns the name", func() {
			Expect(ipSet.Name()).To(Equal(name))
		})
	})

	Describe("Type", func() {
		It("returns the type", func() {
			Expect(ipSet.Type()).To(Equal("IP Set"))
		})
	})
})
