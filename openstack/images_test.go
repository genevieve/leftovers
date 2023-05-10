package openstack_test

import (
	"errors"

	"github.com/genevieve/leftovers/openstack"
	"github.com/genevieve/leftovers/openstack/fakes"
	"github.com/gophercloud/gophercloud/openstack/imageservice/v2/images"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Images", func() {
	Describe("Type", func() {
		It("is of type Image", func() {
			images := openstack.NewImages(nil, nil)

			Expect(images.Type()).To(Equal("Image"))
		})
	})

	Describe("List", func() {
		var (
			fakeLogger      *fakes.Logger
			fakeImageClient *fakes.ImageServiceClient
			filter          string

			subject openstack.Images
		)

		BeforeEach(func() {
			fakeImageClient = &fakes.ImageServiceClient{}
			fakeLogger = &fakes.Logger{}
			skipUserConfirmation := true
			fakeLogger.PromptWithDetailsCall.Returns.Bool = skipUserConfirmation

			fakeImageClient.ListCall.Returns.ImageSlice = []images.Image{
				images.Image{ID: "id 1", Name: "name 1"},
				images.Image{ID: "id 2", Name: "name 2"},
			}
			filter = ""

			subject = openstack.NewImages(fakeImageClient, fakeLogger)
		})

		It("returns the corresponding resources", func() {
			res, err := subject.List(filter)
			Expect(err).NotTo(HaveOccurred())

			Expect(res).To(HaveLen(2))
			Expect(res[0].Name()).To(Equal("name 1 id 1"))
			Expect(res[1].Name()).To(Equal("name 2 id 2"))
		})

		Context("when the user provides a filter", func() {
			BeforeEach(func() {
				fakeImageClient.ListCall.Returns.ImageSlice = []images.Image{
					images.Image{ID: "id", Name: "banana"},
				}
				subject = openstack.NewImages(fakeImageClient, fakeLogger)
			})

			It("returns the matching resources", func() {
				res, err := subject.List("kiwi")
				Expect(err).NotTo(HaveOccurred())

				Expect(res).To(HaveLen(0))
			})
		})

		Context("when the user wants to confirm deletions", func() {
			BeforeEach(func() {
				fakeLogger.PromptWithDetailsCall.Stub = func(string, string) bool {
					return fakeLogger.PromptWithDetailsCall.CallCount > 1
				}
			})

			It("only returns confirmed images", func() {
				res, err := subject.List(filter)
				Expect(err).NotTo(HaveOccurred())

				Expect(fakeLogger.PromptWithDetailsCall.CallCount).To(Equal(2))
				Expect(fakeLogger.PromptWithDetailsCall.Receives.ResourceType).To(Equal("Image"))
				Expect(fakeLogger.PromptWithDetailsCall.Receives.ResourceName).To(Equal("name 2 id 2"))

				Expect(res).To(HaveLen(1))
				Expect(res[0].Name()).To(Equal("name 2 id 2"))
			})
		})

		Context("when the client fails to list images", func() {
			It("returns an error", func() {
				fakeImageClient.ListCall.Returns.Error = errors.New("error getting list")

				_, err := subject.List(filter)
				Expect(err).To(MatchError("List Images: error getting list"))
			})
		})
	})
})
