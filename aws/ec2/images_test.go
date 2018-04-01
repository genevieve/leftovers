package ec2_test

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
	"github.com/genevieve/leftovers/aws/ec2"
	"github.com/genevieve/leftovers/aws/ec2/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Images", func() {
	var (
		client       *fakes.ImagesClient
		logger       *fakes.Logger
		resourceTags *fakes.ResourceTags

		images ec2.Images
	)

	BeforeEach(func() {
		client = &fakes.ImagesClient{}
		logger = &fakes.Logger{}
		logger.PromptWithDetailsCall.Returns.Proceed = true
		resourceTags = &fakes.ResourceTags{}

		images = ec2.NewImages(client, logger, resourceTags)
	})

	Describe("List", func() {
		var filter string

		BeforeEach(func() {
			client.DescribeImagesCall.Returns.Output = &awsec2.DescribeImagesOutput{
				Images: []*awsec2.Image{{
					// State: &awsec2.ImageState{Name: aws.String("available")},
					// Tags: []*awsec2.Tag{{
					// 	Key:   aws.String("Name"),
					// 	Value: aws.String("banana-image"),
					// }},
					ImageId: aws.String("the-image-id"),
				}},
			}
		})

		It("returns a list of ec2 images to delete", func() {
			items, err := images.List(filter)
			Expect(err).NotTo(HaveOccurred())

			Expect(client.DescribeImagesCall.CallCount).To(Equal(1))
			Expect(logger.PromptWithDetailsCall.CallCount).To(Equal(1))
			Expect(logger.PromptWithDetailsCall.Receives.Type).To(Equal("EC2 Image"))
			Expect(logger.PromptWithDetailsCall.Receives.Name).To(Equal("the-image-id"))

			Expect(items).To(HaveLen(1))
		})

		Context("when the image name does not contain the filter", func() {
			It("does not add it to the list", func() {
				items, err := images.List("kiwi")
				Expect(err).NotTo(HaveOccurred())

				Expect(client.DescribeImagesCall.CallCount).To(Equal(1))
				Expect(logger.PromptWithDetailsCall.CallCount).To(Equal(0))

				Expect(items).To(HaveLen(0))
			})
		})

		Context("when the client fails to list images", func() {
			BeforeEach(func() {
				client.DescribeImagesCall.Returns.Error = errors.New("some error")
			})

			It("returns the error", func() {
				_, err := images.List(filter)
				Expect(err).To(MatchError("Describing EC2 Images: some error"))
			})
		})

		Context("when the user responds no to the prompt", func() {
			BeforeEach(func() {
				logger.PromptWithDetailsCall.Returns.Proceed = false
			})

			It("does not return it to the list", func() {
				items, err := images.List(filter)
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PromptWithDetailsCall.CallCount).To(Equal(1))
				Expect(items).To(HaveLen(0))
			})
		})
	})
})
