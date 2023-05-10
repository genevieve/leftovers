package openstack_test

import (
	"errors"

	"github.com/genevieve/leftovers/openstack"
	"github.com/genevieve/leftovers/openstack/fakes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Image", func() {
	Context("when an Image is created", func() {
		It("has a name and a type", func() {
			image := openstack.NewImage("some-name", "some-id", nil)

			Expect(image.Name()).To(Equal("some-name some-id"))
			Expect(image.Type()).To(Equal("Image"))
		})

		Describe("Delete", func() {
			var (
				fakeImageClient *fakes.ImageServiceClient
				image           openstack.Image
			)

			BeforeEach(func() {
				fakeImageClient = &fakes.ImageServiceClient{}
				image = openstack.NewImage("some-name", "some-id", fakeImageClient)
			})

			It("deletes the correct image", func() {
				err := image.Delete()

				Expect(err).NotTo(HaveOccurred())
				Expect(fakeImageClient.DeleteCall.Receives.Id).To(Equal("some-id"))
			})

			Context("when an error occurs", func() {
				Context("when delete fails", func() {
					It("returns an error", func() {
						fakeImageClient.DeleteCall.Returns.Error = errors.New("error description")
						err := image.Delete()

						Expect(err).To(HaveOccurred())
						Expect(err.Error()).To(Equal("error description"))
					})
				})
			})
		})
	})
})
