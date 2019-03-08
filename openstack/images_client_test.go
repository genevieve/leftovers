package openstack_test

import (
	"errors"

	"github.com/genevieve/leftovers/openstack"
	"github.com/genevieve/leftovers/openstack/fakes"
	"github.com/gophercloud/gophercloud/openstack/imageservice/v2/images"
	"github.com/gophercloud/gophercloud/pagination"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ImagesClient", func() {
	var (
		fakeImageAPI *fakes.ImageAPI
		subject      openstack.ImageServiceClient
	)

	BeforeEach(func() {
		fakeImageAPI = &fakes.ImageAPI{}
		subject = openstack.NewImagesClient(fakeImageAPI)
	})

	Describe("List", func() {
		It("lists all the images available for deletion", func() {
			pager := pagination.Pager{}
			page := fakes.Page{}
			imgs := []images.Image{
				images.Image{ID: "hello there"},
				images.Image{ID: "general"},
			}
			fakeImageAPI.GetImagePagerCall.Returns.Pager = pager
			fakeImageAPI.PagerToPageCall.Returns.Page = page
			fakeImageAPI.PageToImagesCall.Returns.Images = imgs

			result, err := subject.List()
			Expect(err).NotTo(HaveOccurred())

			Expect(len(result)).To(Equal(2))
			Expect(result[0].ID).To(Equal("hello there"))
			Expect(result[1].ID).To(Equal("general"))
			Expect(fakeImageAPI.PagerToPageCall.Receives.Pager).To(Equal(pager))
			Expect(fakeImageAPI.PageToImagesCall.Receives.Page).To(Equal(page))
		})

		Context("when an error occurs", func() {
			Context("when there is an error getting a page", func() {
				It("propogates the error", func() {
					pager := pagination.Pager{}
					pager.Err = errors.New("something went horridly wrong")

					fakeImageAPI.GetImagePagerCall.Returns.Pager = pager

					result, err := subject.List()
					Expect(err).To(HaveOccurred())

					Expect(err).To(MatchError("something went horridly wrong"))
					Expect(result).To(BeNil())
				})
			})

			Context("when a pager cannot turn into a page", func() {
				It("propogates the error", func() {
					pager := pagination.Pager{}
					page := fakes.Page{}
					fakeImageAPI.GetImagePagerCall.Returns.Pager = pager
					fakeImageAPI.PagerToPageCall.Returns.Page = page
					fakeImageAPI.PagerToPageCall.Returns.Error = errors.New("oh heck")

					result, err := subject.List()
					Expect(err).To(HaveOccurred())

					Expect(err).To(MatchError("oh heck"))
					Expect(result).To(BeNil())
				})
			})

			Context("when the page cannot be turned into images", func() {
				It("propogates the error", func() {
					pager := pagination.Pager{}
					page := fakes.Page{}
					fakeImageAPI.GetImagePagerCall.Returns.Pager = pager
					fakeImageAPI.PagerToPageCall.Returns.Page = page
					fakeImageAPI.PageToImagesCall.Returns.Error = errors.New("oh no")

					result, err := subject.List()
					Expect(err).To(HaveOccurred())

					Expect(err).To(MatchError("oh no"))
					Expect(result).To(BeNil())
				})
			})
		})
	})

	Describe("Delete", func() {
		It("delete a given image id", func() {
			err := subject.Delete("image-id")
			Expect(err).NotTo(HaveOccurred())

			Expect(fakeImageAPI.DeleteCall.Receives.ImageID).To(Equal("image-id"))
		})

		Context("when an error occurs", func() {
			Context("when deleting fails", func() {
				It("propogates the error", func() {
					fakeImageAPI.DeleteCall.Returns.Error = errors.New("some error")
					err := subject.Delete("image-id")
					Expect(err).To(HaveOccurred())

					Expect(err).To(MatchError("some error"))
				})
			})
		})
	})
})
