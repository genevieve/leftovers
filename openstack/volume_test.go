package openstack_test

import (
	"errors"

	"github.com/genevieve/leftovers/openstack"
	"github.com/genevieve/leftovers/openstack/fakes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Volume", func() {
	var (
		client *fakes.VolumesClient
		volume openstack.Volume
	)

	BeforeEach(func() {
		client = &fakes.VolumesClient{}
		volume = openstack.NewVolume("some-name", "some-id", client)
	})

	Describe("Delete", func() {
		It("deletes the correct volume", func() {
			err := volume.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(client.DeleteCall.Receives.VolumeID).To(Equal("some-id"))
		})

		Context("when the client fails to delete", func() {
			BeforeEach(func() {
				client.DeleteCall.Returns.Error = errors.New("error description")
			})

			It("returns the error", func() {
				err := volume.Delete()
				Expect(err).To(MatchError("error description"))
			})
		})
	})

	Describe("Type", func() {
		It("returns the type", func() {
			Expect(volume.Type()).To(Equal("Volume"))
		})
	})

	Describe("Name", func() {
		It("returns the name", func() {
			Expect(volume.Name()).To(Equal("some-name some-id"))
		})
	})
})
