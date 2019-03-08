package openstack_test

import (
	"errors"

	"github.com/genevieve/leftovers/openstack"
	"github.com/genevieve/leftovers/openstack/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Volume", func() {
	Describe("NewVolume", func() {
		It("has a name and type", func() {
			volume := openstack.NewVolume("some-name", "some-id", nil)

			Expect(volume.Name()).To(Equal("some-name some-id"))
			Expect(volume.Type()).To(Equal("Volume"))
		})
	})

	Describe("Delete", func() {
		var (
			fakeVolumesClient *fakes.VolumesClient
			volume            openstack.Volume
		)

		BeforeEach(func() {
			fakeVolumesClient = &fakes.VolumesClient{}
			volume = openstack.NewVolume("some-name", "some-id", fakeVolumesClient)
		})

		It("deletes the correct volume", func() {
			err := volume.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(fakeVolumesClient.DeleteCall.Receives.VolumeID).To(Equal("some-id"))
		})

		Context("when an error occurs", func() {
			It("returns an error", func() {
				fakeVolumesClient.DeleteCall.Returns.Error = errors.New("error description")

				err := volume.Delete()
				Expect(err).To(HaveOccurred())

				Expect(err).To(MatchError("error description"))
			})
		})
	})
})
