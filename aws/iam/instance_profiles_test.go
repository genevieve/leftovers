package iam_test

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	awsiam "github.com/aws/aws-sdk-go/service/iam"
	"github.com/genevievelesperance/leftovers/aws/iam"
	"github.com/genevievelesperance/leftovers/aws/iam/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("InstanceProfiles", func() {
	var (
		iamClient *fakes.IAMClient
		logger    *fakes.Logger

		instanceProfiles iam.InstanceProfiles
	)

	BeforeEach(func() {
		iamClient = &fakes.IAMClient{}
		logger = &fakes.Logger{}

		instanceProfiles = iam.NewInstanceProfiles(iamClient, logger)
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			logger.PromptCall.Returns.Proceed = true
			iamClient.ListInstanceProfilesCall.Returns.Output = &awsiam.ListInstanceProfilesOutput{
				InstanceProfiles: []*awsiam.InstanceProfile{{
					InstanceProfileName: aws.String("banana"),
				}},
			}
		})

		It("deletes iam instance profiles", func() {
			err := instanceProfiles.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(iamClient.DeleteInstanceProfileCall.CallCount).To(Equal(1))
			Expect(iamClient.DeleteInstanceProfileCall.Receives.Input.InstanceProfileName).To(Equal(aws.String("banana")))
			Expect(logger.PrintfCall.Messages).To(Equal([]string{"SUCCESS deleting instance profile banana\n"}))
		})

		Context("when the client fails to list instance profiles", func() {
			BeforeEach(func() {
				iamClient.ListInstanceProfilesCall.Returns.Error = errors.New("listing error")
			})

			It("returns the error and does not try deleting them", func() {
				err := instanceProfiles.Delete()
				Expect(err.Error()).To(Equal("Listing instance profiles: listing error"))

				Expect(iamClient.DeleteInstanceProfileCall.CallCount).To(Equal(0))
			})
		})

		Context("when the client fails to delete the instance profile", func() {
			BeforeEach(func() {
				iamClient.DeleteInstanceProfileCall.Returns.Error = errors.New("deleting error")
			})

			It("logs the error", func() {
				err := instanceProfiles.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(iamClient.DeleteInstanceProfileCall.CallCount).To(Equal(1))
				Expect(logger.PrintfCall.Messages).To(Equal([]string{"ERROR deleting instance profile banana: deleting error\n"}))
			})
		})

		Context("when the user responds no to the prompt", func() {
			BeforeEach(func() {
				logger.PromptCall.Returns.Proceed = false
			})

			It("does not delete the instance profile", func() {
				err := instanceProfiles.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete instance profile banana?"))
				Expect(iamClient.DeleteInstanceProfileCall.CallCount).To(Equal(0))
			})
		})
	})
})
