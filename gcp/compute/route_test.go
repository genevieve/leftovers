package compute_test

import (
	"errors"
	"fmt"

	"github.com/genevieve/leftovers/gcp/compute"
	"github.com/genevieve/leftovers/gcp/compute/fakes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Route", func() {
	var (
		client  *fakes.RoutesClient
		name    string
		network string

		route compute.Route
	)

	BeforeEach(func() {
		client = &fakes.RoutesClient{}
		name = "banana"
		network = ".com/networks/kiwi-network"

		client.GetNetworkNameCall.Returns.Name = "kiwi-network"

		route = compute.NewRoute(client, name, network)
	})

	Describe("Delete", func() {
		It("deletes the route", func() {
			err := route.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(client.DeleteRouteCall.CallCount).To(Equal(1))
			Expect(client.DeleteRouteCall.Receives.Route).To(Equal(name))
		})

		Context("when the client fails to delete", func() {
			BeforeEach(func() {
				client.DeleteRouteCall.Returns.Error = errors.New("the-error")
			})

			It("returns the error", func() {
				err := route.Delete()
				Expect(err).To(MatchError("Delete: the-error"))
			})
		})
	})

	Describe("Name", func() {
		It("returns the name", func() {
			Expect(route.Name()).To(Equal(fmt.Sprintf("%s (kiwi-network)", name)))
		})
	})

	Describe("Type", func() {
		It("returns the type", func() {
			Expect(route.Type()).To(Equal("Route"))
		})
	})
})
