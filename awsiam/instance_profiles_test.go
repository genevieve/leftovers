package awsiam_test

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/genevievelesperance/leftovers/awsiam"
	"github.com/genevievelesperance/leftovers/awsiam/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("InstanceProfiles", func() {
	var (
		iamClient *fakes.IAMClient
		logger    *fakes.Logger

		instanceProfiles awsiam.InstanceProfiles
	)

	BeforeEach(func() {
		iamClient = &fakes.IAMClient{}
		logger = &fakes.Logger{}

		instanceProfiles = awsiam.NewInstanceProfiles(iamClient, logger)
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			iamClient.ListInstanceProfilesCall.Returns.Output = &iam.ListInstanceProfilesOutput{
				InstanceProfiles: []*iam.InstanceProfile{{
					InstanceProfileName: aws.String("banana"),
				}},
			}
		})

		It("deletes iam instance profiles", func() {
			instanceProfiles.Delete()

			Expect(iamClient.DeleteInstanceProfileCall.CallCount).To(Equal(1))
			Expect(iamClient.DeleteInstanceProfileCall.Receives.Input.InstanceProfileName).To(Equal(aws.String("banana")))
			Expect(logger.PrintfCall.Messages).To(Equal([]string{"SUCCESS deleting instance profile banana\n"}))
		})

		Context("when the client fails to list instance profiles", func() {
			BeforeEach(func() {
				iamClient.ListInstanceProfilesCall.Returns.Error = errors.New("some error")
				iamClient.ListInstanceProfilesCall.Returns.Output = &iam.ListInstanceProfilesOutput{}
			})

			It("does not try deleting them", func() {
				instanceProfiles.Delete()

				Expect(iamClient.DeleteInstanceProfileCall.CallCount).To(Equal(0))
				Expect(logger.PrintfCall.Messages).To(Equal([]string{"ERROR listing instance profiles: some error\n"}))
			})
		})

		Context("when the client fails to delete the instance profile", func() {
			BeforeEach(func() {
				iamClient.DeleteInstanceProfileCall.Returns.Error = errors.New("some error")
			})

			It("returns the error", func() {
				instanceProfiles.Delete()

				Expect(iamClient.DeleteInstanceProfileCall.CallCount).To(Equal(1))
				Expect(logger.PrintfCall.Messages).To(Equal([]string{"ERROR deleting instance profile banana: some error\n"}))
			})
		})
	})
})
