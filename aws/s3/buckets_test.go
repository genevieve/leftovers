package s3_test

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	awss3 "github.com/aws/aws-sdk-go/service/s3"
	"github.com/genevievelesperance/leftovers/aws/s3"
	"github.com/genevievelesperance/leftovers/aws/s3/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Buckets", func() {
	var (
		client  *fakes.BucketsClient
		logger  *fakes.Logger
		manager *fakes.BucketManager

		buckets s3.Buckets
	)

	BeforeEach(func() {
		client = &fakes.BucketsClient{}
		logger = &fakes.Logger{}
		manager = &fakes.BucketManager{}

		buckets = s3.NewBuckets(client, logger, manager)
	})

	Describe("Delete", func() {
		var filter string

		BeforeEach(func() {
			logger.PromptCall.Returns.Proceed = true
			client.ListBucketsCall.Returns.Output = &awss3.ListBucketsOutput{
				Buckets: []*awss3.Bucket{{
					Name: aws.String("banana"),
				}},
			}
			manager.IsInRegionCall.Returns.Output = true
			filter = "ban"
		})

		It("deletes s3 buckets", func() {
			err := buckets.Delete(filter)
			Expect(err).NotTo(HaveOccurred())

			Expect(client.ListBucketsCall.CallCount).To(Equal(1))
			Expect(manager.IsInRegionCall.CallCount).To(Equal(1))
			Expect(manager.IsInRegionCall.Receives.Bucket).To(Equal("banana"))

			Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete bucket banana?"))

			Expect(client.DeleteBucketCall.CallCount).To(Equal(1))
			Expect(client.DeleteBucketCall.Receives.Input.Bucket).To(Equal(aws.String("banana")))

			Expect(logger.PrintfCall.Messages).To(Equal([]string{"SUCCESS deleting bucket banana\n"}))
		})

		Context("when the client fails to list buckets", func() {
			BeforeEach(func() {
				client.ListBucketsCall.Returns.Error = errors.New("some error")
			})

			It("returns the error and does not try deleting them", func() {
				err := buckets.Delete(filter)
				Expect(err).To(MatchError("Listing buckets: some error"))

				Expect(client.DeleteBucketCall.CallCount).To(Equal(0))
			})
		})

		Context("when the bucket name does not contain the filter", func() {
			It("does not try deleting it", func() {
				err := buckets.Delete("kiwi")
				Expect(err).NotTo(HaveOccurred())

				Expect(manager.IsInRegionCall.CallCount).To(Equal(0))
				Expect(logger.PromptCall.CallCount).To(Equal(0))
				Expect(client.DeleteBucketCall.CallCount).To(Equal(0))
			})
		})

		Context("when the bucket isn't in the region configured", func() {
			BeforeEach(func() {
				manager.IsInRegionCall.Returns.Output = false
			})

			It("does not delete the bucket", func() {
				err := buckets.Delete(filter)
				Expect(err).NotTo(HaveOccurred())

				Expect(client.DeleteBucketCall.CallCount).To(Equal(0))
			})
		})

		Context("when the client fails to delete the bucket", func() {
			BeforeEach(func() {
				client.DeleteBucketCall.Returns.Error = errors.New("some error")
			})

			It("logs the error", func() {
				err := buckets.Delete(filter)
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PrintfCall.Messages).To(Equal([]string{"ERROR deleting bucket banana: some error\n"}))
			})
		})

		Context("when the client fails to delete the bucket and list object versions", func() {
			BeforeEach(func() {
				client.DeleteBucketCall.Returns.Error = awserr.New("BucketNotEmpty", "some error", errors.New("some error"))
				client.ListObjectVersionsCall.Returns.Error = errors.New("some other error")
			})

			It("logs the error", func() {
				err := buckets.Delete(filter)
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PrintfCall.Messages).To(Equal([]string{"ERROR deleting bucket banana: some other error\n"}))
			})
		})

		Context("when the user responds no to the prompt", func() {
			BeforeEach(func() {
				logger.PromptCall.Returns.Proceed = false
			})

			It("does not delete the bucket", func() {
				err := buckets.Delete(filter)
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete bucket banana?"))
				Expect(client.DeleteBucketCall.CallCount).To(Equal(0))
			})
		})
	})
})
