package groupingobjects_test

import (
	"context"
	"errors"

	"github.com/genevieve/leftovers/nsxt/groupingobjects"
	"github.com/genevieve/leftovers/nsxt/groupingobjects/fakes"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("NS Service", func() {
	var (
		client *fakes.GroupingObjectsAPI
		ctx    context.Context
		name   string
		id     string

		nsService groupingobjects.NSService
	)

	BeforeEach(func() {
		client = &fakes.GroupingObjectsAPI{}
		name = "mango"
		id = "mango-123"

		ctx = context.WithValue(context.Background(), "fruit", "mango")

		nsService = groupingobjects.NewNSService(client, ctx, name, id)
	})

	Describe("Delete", func() {
		It("deletes the ns service", func() {
			err := nsService.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(client.DeleteNSServiceCall.CallCount).To(Equal(1))
			Expect(client.DeleteNSServiceCall.Receives.String).To(Equal(id))
			Expect(client.DeleteNSServiceCall.Receives.Context).To(Equal(ctx))
		})

		Context("when the client fails to delete the ns service", func() {
			BeforeEach(func() {
				client.DeleteNSServiceCall.Returns.Error = errors.New("insufficient funds")
			})

			It("returns the error", func() {
				err := nsService.Delete()
				Expect(err).To(MatchError("Delete: insufficient funds"))
			})
		})
	})

	Describe("Name", func() {
		It("returns the name", func() {
			Expect(nsService.Name()).To(Equal(name))
		})
	})

	Describe("Type", func() {
		It("returns the type", func() {
			Expect(nsService.Type()).To(Equal("NS Service"))
		})
	})
})
