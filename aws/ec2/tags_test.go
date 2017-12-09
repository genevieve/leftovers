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
		ec2Client *fakes.EC2Client
		logger    *fakes.Logger

		tags ec2.Tags
	)

	BeforeEach(func() {
		ec2Client = &fakes.EC2Client{}
		logger = &fakes.Logger{}

		tags = ec2.NewTags(ec2Client, logger)
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			logger.PromptCall.Returns.Proceed = true
			ec2Client.DescribeTagsCall.Returns.Output = &awsec2.DescribeTagsOutput{
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

			Expect(ec2Client.DescribeTagsCall.CallCount).To(Equal(1))
			Expect(ec2Client.DeleteTagsCall.CallCount).To(Equal(1))
			Expect(ec2Client.DeleteTagsCall.Receives.Input.Tags[0].Key).To(Equal(aws.String("the-key")))
			Expect(ec2Client.DeleteTagsCall.Receives.Input.Resources[0]).To(Equal(aws.String("the-resource-id")))
			Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete tag banana?"))
			Expect(logger.PrintfCall.Messages).To(Equal([]string{"SUCCESS deleting tag banana\n"}))
		})

		Context("when the client fails to list tags", func() {
			BeforeEach(func() {
				ec2Client.DescribeTagsCall.Returns.Error = errors.New("some error")
			})

			It("does not try deleting them", func() {
				err := tags.Delete()
				Expect(err.Error()).To(Equal("Describing tags: some error"))

				Expect(ec2Client.DeleteTagsCall.CallCount).To(Equal(0))
			})
		})

		Context("when the client fails to delete the tag", func() {
			BeforeEach(func() {
				ec2Client.DeleteTagsCall.Returns.Error = errors.New("some error")
			})

			It("returns the error", func() {
				err := tags.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PrintfCall.Messages).To(Equal([]string{"ERROR deleting tag banana: some error\n"}))
			})
		})

		Context("when the user responds no to the prompt", func() {
			BeforeEach(func() {
				logger.PromptCall.Returns.Proceed = false
			})

			It("returns the error", func() {
				err := tags.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete tag banana?"))
				Expect(ec2Client.DeleteTagsCall.CallCount).To(Equal(0))
			})
		})
	})
})
