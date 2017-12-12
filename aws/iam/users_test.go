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

var _ = Describe("Users", func() {
	var (
		client   *fakes.UsersClient
		logger   *fakes.Logger
		policies *fakes.UserPolicies

		users iam.Users
	)

	BeforeEach(func() {
		client = &fakes.UsersClient{}
		logger = &fakes.Logger{}
		policies = &fakes.UserPolicies{}

		users = iam.NewUsers(client, logger, policies)
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			logger.PromptCall.Returns.Proceed = true
			client.ListUsersCall.Returns.Output = &awsiam.ListUsersOutput{
				Users: []*awsiam.User{{
					UserName: aws.String("banana"),
				}},
			}
		})

		It("deletes iam users and associated policies", func() {
			err := users.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(client.ListUsersCall.CallCount).To(Equal(1))
			Expect(policies.DeleteCall.CallCount).To(Equal(1))
			Expect(policies.DeleteCall.Receives.UserName).To(Equal("banana"))
			Expect(client.DeleteUserCall.CallCount).To(Equal(1))
			Expect(client.DeleteUserCall.Receives.Input.UserName).To(Equal(aws.String("banana")))
			Expect(logger.PrintfCall.Messages).To(Equal([]string{"SUCCESS deleting user banana\n"}))
		})

		Context("when the client fails to list users", func() {
			BeforeEach(func() {
				client.ListUsersCall.Returns.Error = errors.New("some error")
			})

			It("returns the error and does not try deleting them", func() {
				err := users.Delete()
				Expect(err.Error()).To(Equal("Listing users: some error"))

				Expect(client.DeleteUserCall.CallCount).To(Equal(0))
			})
		})

		Context("when policies fails to delete", func() {
			BeforeEach(func() {
				policies.DeleteCall.Returns.Error = errors.New("some error")
			})

			It("returns the error", func() {
				err := users.Delete()
				Expect(err.Error()).To(Equal("Deleting policies for banana: some error"))

				Expect(policies.DeleteCall.CallCount).To(Equal(1))
			})
		})

		Context("when the client fails to delete the user", func() {
			BeforeEach(func() {
				client.DeleteUserCall.Returns.Error = errors.New("some error")
			})

			It("logs the error", func() {
				err := users.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PrintfCall.Messages).To(Equal([]string{"ERROR deleting user banana: some error\n"}))
			})
		})

		Context("when the user responds no to the prompt", func() {
			BeforeEach(func() {
				logger.PromptCall.Returns.Proceed = false
			})

			It("does not delete the user", func() {
				err := users.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete user banana?"))
				Expect(client.DeleteUserCall.CallCount).To(Equal(0))
			})
		})
	})
})
