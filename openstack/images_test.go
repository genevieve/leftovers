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
			fakeImageClient *fakes.ImageServiceClient
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

			subject = openstack.NewImages(fakeImageClient, fakeLogger)
		})

		It("returns the corresponding resources", func() {
			res, err := subject.List()
			Expect(err).NotTo(HaveOccurred())

			Expect(len(res)).To(Equal(2))
			Expect(res[0].Name()).To(Equal("name 1 id 1"))
			Expect(res[1].Name()).To(Equal("name 2 id 2"))
		})

		Context("when the user wants to confirm deletions", func() {
			BeforeEach(func() {
				fakeLogger.PromptWithDetailsCall.Stub = func(string, string) bool {
					return fakeLogger.PromptWithDetailsCall.CallCount > 1
				}
			})

			It("only returns confirmed images", func() {
				res, err := subject.List()
				Expect(err).NotTo(HaveOccurred())

				Expect(len(res)).To(Equal(1))
				Expect(res[0].Name()).To(Equal("name 2 id 2"))
			})

			It("prompts the user", func() {
				subject.List()

				Expect(fakeLogger.PromptWithDetailsCall.CallCount).To(Equal(2))
				Expect(fakeLogger.PromptWithDetailsCall.Receives.ResourceType).To(Equal("Image"))
				Expect(fakeLogger.PromptWithDetailsCall.Receives.ResourceName).To(Equal("name 2 id 2"))
			})
		})

		Context("when an error occurs", func() {
			It("returns an error", func() {
				fakeImageClient.ListCall.Returns.ImageSlice = nil
				fakeImageClient.ListCall.Returns.Error = errors.New("error getting list")

				res, err := subject.List()
				Expect(err).To(MatchError("List Images: error getting list"))
				Expect(res).To(BeNil())
			})
		})
	})
})
