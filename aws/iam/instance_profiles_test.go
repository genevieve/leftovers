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
		client *fakes.InstanceProfilesClient
		logger *fakes.Logger

		instanceProfiles iam.InstanceProfiles
	)

	BeforeEach(func() {
		client = &fakes.InstanceProfilesClient{}
		logger = &fakes.Logger{}

		instanceProfiles = iam.NewInstanceProfiles(client, logger)
	})

	Describe("Delete", func() {
		var filter string

		BeforeEach(func() {
			logger.PromptCall.Returns.Proceed = true
			client.ListInstanceProfilesCall.Returns.Output = &awsiam.ListInstanceProfilesOutput{
				InstanceProfiles: []*awsiam.InstanceProfile{{
					InstanceProfileName: aws.String("banana-profile"),
				}},
			}
			filter = "banana"
		})

		It("deletes iam instance profiles and detaches roles", func() {
			err := instanceProfiles.Delete(filter)
			Expect(err).NotTo(HaveOccurred())

			Expect(logger.PromptCall.CallCount).To(Equal(1))

			Expect(client.DeleteInstanceProfileCall.CallCount).To(Equal(1))
			Expect(client.DeleteInstanceProfileCall.Receives.Input.InstanceProfileName).To(Equal(aws.String("banana-profile")))

			Expect(logger.PrintfCall.Messages).To(Equal([]string{
				"SUCCESS deleting instance profile banana-profile\n",
			}))
		})

		Context("when the instance profile name does not contain the filter", func() {
			It("does not delete it", func() {
				err := instanceProfiles.Delete("kiwi")
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PromptCall.CallCount).To(Equal(0))
				Expect(client.DeleteInstanceProfileCall.CallCount).To(Equal(0))
			})
		})

		Context("when the client fails to list instance profiles", func() {
			BeforeEach(func() {
				client.ListInstanceProfilesCall.Returns.Error = errors.New("listing error")
			})

			It("returns the error and does not try deleting them", func() {
				err := instanceProfiles.Delete(filter)
				Expect(err).To(MatchError("Listing instance profiles: listing error"))

				Expect(client.DeleteInstanceProfileCall.CallCount).To(Equal(0))
			})
		})

		Context("when there are roles", func() {
			BeforeEach(func() {
				client.ListInstanceProfilesCall.Returns.Output = &awsiam.ListInstanceProfilesOutput{
					InstanceProfiles: []*awsiam.InstanceProfile{{
						InstanceProfileName: aws.String("banana-profile"),
						Roles:               []*awsiam.Role{{RoleName: aws.String("the-role")}},
					}},
				}
			})

			It("removes the roles and uses them in the name", func() {
				err := instanceProfiles.Delete(filter)
				Expect(err).NotTo(HaveOccurred())

				Expect(client.RemoveRoleFromInstanceProfileCall.CallCount).To(Equal(1))
				Expect(client.RemoveRoleFromInstanceProfileCall.Receives.Input.InstanceProfileName).To(Equal(aws.String("banana-profile")))
				Expect(client.RemoveRoleFromInstanceProfileCall.Receives.Input.RoleName).To(Equal(aws.String("the-role")))

				Expect(client.DeleteInstanceProfileCall.CallCount).To(Equal(1))
				Expect(client.DeleteInstanceProfileCall.Receives.Input.InstanceProfileName).To(Equal(aws.String("banana-profile")))

				Expect(logger.PrintfCall.Messages).To(Equal([]string{
					"SUCCESS removing role the-role from instance profile banana-profile (Role:the-role)\n",
					"SUCCESS deleting instance profile banana-profile (Role:the-role)\n",
				}))
			})
		})

		Context("when the client fails to remove the role from the instance profile", func() {
			BeforeEach(func() {
				client.ListInstanceProfilesCall.Returns.Output = &awsiam.ListInstanceProfilesOutput{
					InstanceProfiles: []*awsiam.InstanceProfile{{
						InstanceProfileName: aws.String("banana-profile"),
						Roles:               []*awsiam.Role{{RoleName: aws.String("the-role")}},
					}},
				}
				client.RemoveRoleFromInstanceProfileCall.Returns.Error = errors.New("some error")
			})

			It("logs the error and continues", func() {
				err := instanceProfiles.Delete(filter)
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PrintfCall.Messages).To(Equal([]string{
					"ERROR removing role the-role from instance profile banana-profile (Role:the-role): some error\n",
					"SUCCESS deleting instance profile banana-profile (Role:the-role)\n",
				}))
				Expect(client.DeleteInstanceProfileCall.CallCount).To(Equal(1))
			})
		})

		Context("when the client fails to delete the instance profile", func() {
			BeforeEach(func() {
				client.DeleteInstanceProfileCall.Returns.Error = errors.New("deleting error")
			})

			It("logs the error", func() {
				err := instanceProfiles.Delete(filter)
				Expect(err).NotTo(HaveOccurred())

				Expect(client.DeleteInstanceProfileCall.CallCount).To(Equal(1))
				Expect(logger.PrintfCall.Messages).To(Equal([]string{
					"ERROR deleting instance profile banana-profile: deleting error\n",
				}))
			})
		})

		Context("when the user responds no to the prompt", func() {
			BeforeEach(func() {
				logger.PromptCall.Returns.Proceed = false
			})

			It("does not delete the instance profile", func() {
				err := instanceProfiles.Delete(filter)
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete instance profile banana-profile?"))
				Expect(client.DeleteInstanceProfileCall.CallCount).To(Equal(0))
			})
		})
	})
})
