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

	Describe("Delete", func() {
		BeforeEach(func() {
			logger.PromptCall.Returns.Proceed = true
			client.DescribeTagsCall.Returns.Output = &awsec2.DescribeTagsOutput{
				Tags: []*awsec2.TagDescription{{
					Key:        aws.String("the-key"),
					Value:      aws.String("banana"),
					ResourceId: aws.String("the-resource-id"),
				}},
			}
		})

		It("deletes ec2 tags", func() {
			err := tags.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(client.DescribeTagsCall.CallCount).To(Equal(1))
			Expect(client.DeleteTagsCall.CallCount).To(Equal(1))
			Expect(client.DeleteTagsCall.Receives.Input.Tags[0].Key).To(Equal(aws.String("the-key")))
			Expect(client.DeleteTagsCall.Receives.Input.Resources[0]).To(Equal(aws.String("the-resource-id")))
			Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete tag banana?"))
			Expect(logger.PrintfCall.Messages).To(Equal([]string{"SUCCESS deleting tag banana\n"}))
		})

		Context("when the client fails to list tags", func() {
			BeforeEach(func() {
				client.DescribeTagsCall.Returns.Error = errors.New("some error")
			})

			It("does not try deleting them", func() {
				err := tags.Delete()
				Expect(err.Error()).To(Equal("Describing tags: some error"))

				Expect(client.DeleteTagsCall.CallCount).To(Equal(0))
			})
		})

		Context("when the client fails to delete the tag", func() {
			BeforeEach(func() {
				client.DeleteTagsCall.Returns.Error = errors.New("some error")
			})

			It("logs the error", func() {
				err := tags.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PrintfCall.Messages).To(Equal([]string{"ERROR deleting tag banana: some error\n"}))
			})
		})

		Context("when the user responds no to the prompt", func() {
			BeforeEach(func() {
				logger.PromptCall.Returns.Proceed = false
			})

			It("does not delete the tag", func() {
				err := tags.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete tag banana?"))
				Expect(client.DeleteTagsCall.CallCount).To(Equal(0))
			})
		})
	})
})
