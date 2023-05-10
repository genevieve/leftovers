package openstack_test

import (
	"errors"

	"github.com/genevieve/leftovers/openstack"
	"github.com/genevieve/leftovers/openstack/fakes"
	"github.com/gophercloud/gophercloud/openstack/imageservice/v2/images"
	"github.com/gophercloud/gophercloud/pagination"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ImagesClient", func() {
	var (
		api    *fakes.ImageAPI
		client openstack.ImagesClient
	)

	BeforeEach(func() {
		api = &fakes.ImageAPI{}
		client = openstack.NewImagesClient(api)
	})

	Describe("List", func() {
		var (
			pager pagination.Pager
			page  *fakes.Page
			imgs  []images.Image
		)

		BeforeEach(func() {
			pager = pagination.Pager{}
			page = &fakes.Page{}
			imgs = []images.Image{
				{ID: "hello there"},
				{ID: "general"},
			}
			api.GetImagesPagerCall.Returns.Pager = pager
			api.PagerToPageCall.Returns.Page = page
			api.PageToImagesCall.Returns.ImageSlice = imgs
		})

		It("lists all the images available for deletion", func() {
			list, err := client.List()
			Expect(err).NotTo(HaveOccurred())

			Expect(list).To(HaveLen(2))
			Expect(list[0].ID).To(Equal("hello there"))
			Expect(list[1].ID).To(Equal("general"))
		})

		Context("when there is an error getting a page", func() {
			BeforeEach(func() {
				pager.Err = errors.New("something went horridly wrong")
				api.GetImagesPagerCall.Returns.Pager = pager
			})

			It("returns a helpful error message", func() {
				_, err := client.List()
				Expect(err).To(MatchError("get images pager: something went horridly wrong"))
			})
		})

		Context("when a pager cannot turn into a page", func() {
			BeforeEach(func() {
				api.PagerToPageCall.Returns.Error = errors.New("banana")
			})

			It("returns a helpful error message", func() {
				_, err := client.List()
				Expect(err).To(MatchError("pager to page: banana"))
			})
		})

		Context("when the page cannot be turned into images", func() {
			BeforeEach(func() {
				api.PageToImagesCall.Returns.Error = errors.New("oh no")
			})

			It("returns a helpful error message", func() {
				_, err := client.List()
				Expect(err).To(MatchError("page to images: oh no"))
			})
		})
	})

	Describe("Delete", func() {
		It("delete a given image id", func() {
			err := client.Delete("image-id")
			Expect(err).NotTo(HaveOccurred())

			Expect(api.DeleteCall.Receives.ImageID).To(Equal("image-id"))
		})

		Context("when the api fails to delete", func() {
			BeforeEach(func() {
				api.DeleteCall.Returns.Error = errors.New("some error")
			})

			It("returns a helpful error message", func() {
				err := client.Delete("image-id")
				Expect(err).To(MatchError("some error"))
			})
		})
	})
})
