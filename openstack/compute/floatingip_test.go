package compute_test

import (
	"errors"

	"github.com/genevieve/leftovers/openstack/compute"
	"github.com/genevieve/leftovers/openstack/compute/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("FloatingIP", func() {
	var (
		client *fakes.FloatingIPsClient
		name   string

		floatingIP compute.FloatingIP
	)

	BeforeEach(func() {
		client = &fakes.FloatingIPsClient{}
		name = "the-ip"

		floatingIP = compute.NewFloatingIP(client, name)
	})

	Describe("Delete", func() {
		It("deletes the floating ip", func() {
			err := floatingIP.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(client.DeleteFloatingIPCall.CallCount).To(Equal(1))
			Expect(client.DeleteFloatingIPCall.Receives.FloatingIP).To(Equal(name))
		})

		Context("when the client fails to delete", func() {
			BeforeEach(func() {
				client.DeleteFloatingIPCall.Returns.Error = errors.New("the-error")
			})

			It("returns the error", func() {
				err := floatingIP.Delete()
				Expect(err).To(MatchError("Delete: the-error"))
			})
		})
	})

	Describe("Name", func() {
		It("returns the name", func() {
			Expect(floatingIP.Name()).To(Equal(name))
		})
	})

	Describe("Type", func() {
		It("returns the type", func() {
			Expect(floatingIP.Type()).To(Equal("Floating IP"))
		})
	})
})
