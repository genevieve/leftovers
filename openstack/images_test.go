package openstack_test

import (
	"errors"

	"github.com/genevieve/leftovers/openstack"
	"github.com/genevieve/leftovers/openstack/fakes"
	openstackimages "github.com/gophercloud/gophercloud/openstack/imageservice/v2/images"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Images", func() {
	var (
		logger *fakes.Logger
		client *fakes.ImageClient
		images openstack.Images
	)

	BeforeEach(func() {
		client = &fakes.ImageClient{}
		logger = &fakes.Logger{}

		images = openstack.NewImages(client, logger)
	})

	Describe("List", func() {
		BeforeEach(func() {
			logger.PromptWithDetailsCall.Returns.Bool = true

			client.ListCall.Returns.Images = []openstackimages.Image{
				{ID: "id 1", Name: "name 1"},
				{ID: "id 2", Name: "name 2"},
			}
		})

		It("returns the corresponding resources", func() {
			list, err := images.List()
			Expect(err).NotTo(HaveOccurred())

			Expect(logger.PromptWithDetailsCall.CallCount).To(Equal(2))
			Expect(logger.PromptWithDetailsCall.Receives.ResourceType).To(Equal("Image"))
			Expect(logger.PromptWithDetailsCall.Receives.ResourceName).To(Equal("name 2 id 2"))

			Expect(list).To(HaveLen(2))
			Expect(list[0].Name()).To(Equal("name 1 id 1"))
			Expect(list[1].Name()).To(Equal("name 2 id 2"))
		})

		Context("when the user wants to confirm deletions", func() {
			BeforeEach(func() {
				logger.PromptWithDetailsCall.ReturnsForCall = append(logger.PromptWithDetailsCall.ReturnsForCall,
					fakes.LoggerPromptWithDetailsCallReturn{Bool: false},
					fakes.LoggerPromptWithDetailsCallReturn{Bool: true},
				)
			})

			It("only returns confirmed images", func() {
				list, err := images.List()
				Expect(err).NotTo(HaveOccurred())

				Expect(list).To(HaveLen(1))
				Expect(list[0].Name()).To(Equal("name 2 id 2"))
			})
		})

		Context("when the image client returns an error", func() {
			BeforeEach(func() {
				client.ListCall.Returns.Error = errors.New("banana")
			})

			It("returns a helpful error message", func() {
				_, err := images.List()
				Expect(err).To(MatchError("List Images: banana"))
			})
		})
	})

	Describe("Type", func() {
		It("returns the resource type", func() {
			Expect(images.Type()).To(Equal("Image"))
		})
	})
})
