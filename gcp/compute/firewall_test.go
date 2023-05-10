package compute_test

import (
	"errors"
	"fmt"

	"github.com/genevieve/leftovers/gcp/compute"
	"github.com/genevieve/leftovers/gcp/compute/fakes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Firewall", func() {
	var (
		client  *fakes.FirewallsClient
		name    string
		network string

		firewall compute.Firewall
	)

	BeforeEach(func() {
		client = &fakes.FirewallsClient{}
		name = "banana"
		network = "global/networks/kiwi-network"

		client.GetNetworkNameCall.Returns.Name = "kiwi-network"

		firewall = compute.NewFirewall(client, name, network)
	})

	Describe("Delete", func() {
		It("deletes the firewall", func() {
			err := firewall.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(client.DeleteFirewallCall.CallCount).To(Equal(1))
			Expect(client.DeleteFirewallCall.Receives.Firewall).To(Equal(name))
		})

		Context("when the client fails to delete", func() {
			BeforeEach(func() {
				client.DeleteFirewallCall.Returns.Error = errors.New("the-error")
			})

			It("returns the error", func() {
				err := firewall.Delete()
				Expect(err).To(MatchError("Delete: the-error"))
			})
		})
	})

	Describe("Name", func() {
		It("returns the name", func() {
			Expect(firewall.Name()).To(Equal(fmt.Sprintf("%s (kiwi-network)", name)))
		})
	})

	Describe("Type", func() {
		It("returns the type", func() {
			Expect(firewall.Type()).To(Equal("Firewall"))
		})
	})
})
