package ec2_test

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
	"github.com/genevievelesperance/leftovers/aws/ec2"
	"github.com/genevievelesperance/leftovers/aws/ec2/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tags", func() {
	var (
		client *fakes.TagsClient
		logger *fakes.Logger

		tags ec2.Tags
	)

	BeforeEach(func() {
		client = &fakes.TagsClient{}
		logger = &fakes.Logger{}

		tags = ec2.NewTags(client, logger)
	})

	Describe("List", func() {
		var filter string

		BeforeEach(func() {
			logger.PromptCall.Returns.Proceed = true
			client.DescribeTagsCall.Returns.Output = &awsec2.DescribeTagsOutput{
				Tags: []*awsec2.TagDescription{{
					Key:        aws.String("the-key"),
					Value:      aws.String("banana-tag"),
					ResourceId: aws.String("the-resource-id"),
				}},
			}
			filter = "banana"
		})

		It("returns a list of ec2 tags to delete", func() {
			items, err := tags.List(filter)
			Expect(err).NotTo(HaveOccurred())

			Expect(client.DescribeTagsCall.CallCount).To(Equal(1))

			Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete tag banana-tag?"))

			Expect(items).To(HaveLen(1))
			Expect(items).To(HaveKeyWithValue("the-key", "the-resource-id"))
		})

		Context("when the client fails to list tags", func() {
			BeforeEach(func() {
				client.DescribeTagsCall.Returns.Error = errors.New("some error")
			})

			It("returns the error", func() {
				_, err := tags.List(filter)
				Expect(err).To(MatchError("Describing tags: some error"))
			})
		})

		Context("when the tag name does not contain the filter", func() {
			It("does not return it in the list", func() {
				items, err := tags.List("kiwi")
				Expect(err).NotTo(HaveOccurred())

				Expect(client.DescribeTagsCall.CallCount).To(Equal(1))
				Expect(logger.PromptCall.CallCount).To(Equal(0))
				Expect(items).To(HaveLen(0))
			})
		})

		Context("when the user responds no to the prompt", func() {
			BeforeEach(func() {
				logger.PromptCall.Returns.Proceed = false
			})

			It("does not return it in the list", func() {
				items, err := tags.List(filter)
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete tag banana-tag?"))
				Expect(items).To(HaveLen(0))
			})
		})
	})

	Describe("Delete", func() {
		var items map[string]string

		BeforeEach(func() {
			items = map[string]string{"the-key": "the-resource-id"}
		})

		It("deletes ec2 tags", func() {
			err := tags.Delete(items)
			Expect(err).NotTo(HaveOccurred())

			Expect(client.DeleteTagsCall.CallCount).To(Equal(1))
			Expect(client.DeleteTagsCall.Receives.Input.Tags[0].Key).To(Equal(aws.String("the-key")))
			Expect(client.DeleteTagsCall.Receives.Input.Resources[0]).To(Equal(aws.String("the-resource-id")))

			Expect(logger.PrintfCall.Messages).To(Equal([]string{"SUCCESS deleting tag the-key\n"}))
		})

		Context("when the client fails to delete the tag", func() {
			BeforeEach(func() {
				client.DeleteTagsCall.Returns.Error = errors.New("some error")
			})

			It("logs the error", func() {
				err := tags.Delete(items)
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PrintfCall.Messages).To(Equal([]string{"ERROR deleting tag the-key: some error\n"}))
			})
		})
	})
})
