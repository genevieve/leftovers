package openstack_test

import (
	"errors"

	"github.com/genevieve/leftovers/openstack"
	"github.com/genevieve/leftovers/openstack/openstackfakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Volume", func() {
	Context("when a volume is created", func() {

		It("should return volume", func() {
			volume := openstack.NewVolume("some-name", "some-id", nil)
			Expect(volume.Name()).To(Equal("some-name some-id"))
		})

		It("should create a volume with the correct type name", func() {
			volume := openstack.NewVolume("some-name", "some-id", nil)
			Expect(volume.Type()).To(Equal("Volume"))
		})

		Context("when Delete is called", func() {
			var mockVolumesDeleter *openstackfakes.FakeVolumesDeleter
			var volume openstack.Volume
			BeforeEach(func() {
				mockVolumesDeleter = &openstackfakes.FakeVolumesDeleter{}
				volume = openstack.NewVolume("some-name", "some-id", mockVolumesDeleter)
			})
			Context("when there is an error", func() {
				It("should delete the volume", func() {
					mockVolumesDeleter.DeleteReturns(errors.New("error description"))
					err := volume.Delete()

					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(Equal("error description"))
				})
			})

			It("should delete the correct volume", func() {
				mockVolumesDeleter.DeleteReturns(nil)

				err := volume.Delete()

				Expect(err).NotTo(HaveOccurred())
				Expect(mockVolumesDeleter.DeleteArgsForCall(0)).To(Equal("some-id"))
			})

		})
	})

})
