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
		keys     *fakes.AccessKeys

		users iam.Users
	)

	BeforeEach(func() {
		client = &fakes.UsersClient{}
		logger = &fakes.Logger{}
		policies = &fakes.UserPolicies{}
		keys = &fakes.AccessKeys{}

		users = iam.NewUsers(client, logger, policies, keys)
	})

	Describe("Delete", func() {
		var filter string

		BeforeEach(func() {
			logger.PromptCall.Returns.Proceed = true
			client.ListUsersCall.Returns.Output = &awsiam.ListUsersOutput{
				Users: []*awsiam.User{{
					UserName: aws.String("banana-user"),
				}},
			}
			filter = "banana"
		})

		It("deletes iam users and associated policies", func() {
			err := users.Delete(filter)
			Expect(err).NotTo(HaveOccurred())

			Expect(client.ListUsersCall.CallCount).To(Equal(1))

			Expect(logger.PromptCall.CallCount).To(Equal(1))

			Expect(keys.DeleteCall.CallCount).To(Equal(1))
			Expect(keys.DeleteCall.Receives.UserName).To(Equal("banana-user"))

			Expect(policies.DeleteCall.CallCount).To(Equal(1))
			Expect(policies.DeleteCall.Receives.UserName).To(Equal("banana-user"))

			Expect(client.DeleteUserCall.CallCount).To(Equal(1))
			Expect(client.DeleteUserCall.Receives.Input.UserName).To(Equal(aws.String("banana-user")))

			Expect(logger.PrintfCall.Messages).To(Equal([]string{"SUCCESS deleting user banana-user\n"}))
		})

		Context("when the client fails to list users", func() {
			BeforeEach(func() {
				client.ListUsersCall.Returns.Error = errors.New("some error")
			})

			It("returns the error and does not try deleting them", func() {
				err := users.Delete(filter)
				Expect(err).To(MatchError("Listing users: some error"))

				Expect(client.DeleteUserCall.CallCount).To(Equal(0))
			})
		})

		Context("when the user name does not contain the filter", func() {
			It("returns the error", func() {
				err := users.Delete("kiwi")
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PromptCall.CallCount).To(Equal(0))
				Expect(keys.DeleteCall.CallCount).To(Equal(0))
				Expect(client.DeleteUserCall.CallCount).To(Equal(0))
			})
		})

		Context("when access keys fails to delete", func() {
			BeforeEach(func() {
				keys.DeleteCall.Returns.Error = errors.New("some error")
			})

			It("returns the error", func() {
				err := users.Delete(filter)
				Expect(err).To(MatchError("Deleting access keys for banana-user: some error"))

				Expect(keys.DeleteCall.CallCount).To(Equal(1))
			})
		})

		Context("when policies fails to delete", func() {
			BeforeEach(func() {
				policies.DeleteCall.Returns.Error = errors.New("some error")
			})

			It("returns the error", func() {
				err := users.Delete(filter)
				Expect(err).To(MatchError("Deleting policies for banana-user: some error"))

				Expect(policies.DeleteCall.CallCount).To(Equal(1))
			})
		})

		Context("when the client fails to delete the user", func() {
			BeforeEach(func() {
				client.DeleteUserCall.Returns.Error = errors.New("some error")
			})

			It("logs the error", func() {
				err := users.Delete(filter)
				Expect(err).NotTo(HaveOccurred())

				Expect(client.DeleteUserCall.CallCount).To(Equal(1))
				Expect(logger.PrintfCall.Messages).To(Equal([]string{"ERROR deleting user banana-user: some error\n"}))
			})
		})

		Context("when the user responds no to the prompt", func() {
			BeforeEach(func() {
				logger.PromptCall.Returns.Proceed = false
			})

			It("does not delete the user", func() {
				err := users.Delete(filter)
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete user banana-user?"))
				Expect(client.DeleteUserCall.CallCount).To(Equal(0))
			})
		})
	})
})
