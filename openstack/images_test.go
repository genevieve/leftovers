package openstack_test

import (
	"errors"

	"github.com/genevieve/leftovers/openstack"
	"github.com/genevieve/leftovers/openstack/fakes"
	"github.com/gophercloud/gophercloud/openstack/imageservice/v2/images"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Images", func() {
	Describe("Type", func() {
		It("is of type Image", func() {
			images := openstack.NewImages(nil, nil)
			result := images.Type()

			Expect(result).To(Equal("Image"))
		})
	})

	Describe("List", func() {
		var (
			subject         openstack.Images
			fakeLogger      *fakes.Logger
			fakeImageClient *fakes.ImageClient
		)

		BeforeEach(func() {
			fakeImageClient = &fakes.ImageClient{}
			fakeLogger = &fakes.Logger{}
			skipUserConfirmation := true
			fakeLogger.PromptWithDetailsCall.Returns.Bool = skipUserConfirmation

			subject = openstack.NewImages(fakeImageClient, fakeLogger)
		})

		It("should return the corresponding resources", func() {
			fakeImageClient.ListCall.Returns.Images = []images.Image{
				images.Image{ID: "id 1", Name: "name 1"},
				images.Image{ID: "id 2", Name: "name 2"},
			}

			res, err := subject.List()

			Expect(err).NotTo(HaveOccurred())
			Expect(res[0].Name()).To(Equal("name 1 id 1"))

			Expect(res[1].Name()).To(Equal("name 2 id 2"))
			Expect(len(res)).To(Equal(2))
		})

		Context("when an error occurs", func() {
			It("returns an error", func() {
				fakeImageClient.ListCall.Returns.Images = nil
				fakeImageClient.ListCall.Returns.Error = errors.New("error getting list")

				res, err := subject.List()

				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("error getting list"))
				Expect(res).To(BeNil())
			})
		})
	})
})
