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

var _ = Describe("UserPolicies", func() {
	var (
		client *fakes.UserPoliciesClient
		logger *fakes.Logger

		policies iam.UserPolicies
	)

	BeforeEach(func() {
		client = &fakes.UserPoliciesClient{}
		logger = &fakes.Logger{}

		policies = iam.NewUserPolicies(client, logger)
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			logger.PromptCall.Returns.Proceed = true
			client.ListUserPoliciesCall.Returns.Output = &awsiam.ListUserPoliciesOutput{
				PolicyNames: []*string{aws.String("the-policy")},
			}
			client.ListPoliciesCall.Returns.Output = &awsiam.ListPoliciesOutput{
				Policies: []*awsiam.Policy{{
					Arn:        aws.String("the-policy-arn"),
					PolicyName: aws.String("the-policy"),
				}},
			}
		})

		It("detaches and deletes the policies", func() {
			err := policies.Delete("the-user")
			Expect(err).NotTo(HaveOccurred())

			Expect(client.ListUserPoliciesCall.CallCount).To(Equal(1))
			Expect(client.ListUserPoliciesCall.Receives.Input.UserName).To(Equal(aws.String("the-user")))

			Expect(client.ListPoliciesCall.CallCount).To(Equal(1))
			Expect(client.ListPoliciesCall.Receives.Input.Scope).To(Equal(aws.String("Local")))

			Expect(client.DetachUserPolicyCall.CallCount).To(Equal(1))
			Expect(client.DetachUserPolicyCall.Receives.Input.UserName).To(Equal(aws.String("the-user")))
			Expect(client.DetachUserPolicyCall.Receives.Input.PolicyArn).To(Equal(aws.String("the-policy-arn")))

			Expect(client.DeleteUserPolicyCall.CallCount).To(Equal(1))
			Expect(client.DeleteUserPolicyCall.Receives.Input.UserName).To(Equal(aws.String("the-user")))
			Expect(client.DeleteUserPolicyCall.Receives.Input.PolicyName).To(Equal(aws.String("the-policy")))

			Expect(logger.PrintfCall.Messages).To(Equal([]string{
				"SUCCESS detaching user policy the-policy\n",
				"SUCCESS deleting user policy the-policy\n",
			}))
		})

		Context("when the client fails to list user policies", func() {
			BeforeEach(func() {
				client.ListUserPoliciesCall.Returns.Error = errors.New("some error")
			})

			It("returns the error and does not try deleting them", func() {
				err := policies.Delete("banana")
				Expect(err.Error()).To(Equal("Listing user policies: some error"))

				Expect(client.DeleteUserPolicyCall.CallCount).To(Equal(0))
			})
		})

		Context("when the client fails to get the user policy", func() {
			BeforeEach(func() {
				client.ListPoliciesCall.Returns.Error = errors.New("some error")
			})

			It("logs the error and does not try to detach the user policy", func() {
				err := policies.Delete("banana")
				Expect(err).NotTo(HaveOccurred())

				Expect(client.DetachUserPolicyCall.CallCount).To(Equal(0))
				Expect(client.DeleteUserPolicyCall.CallCount).To(Equal(1))
				Expect(logger.PrintfCall.Messages).To(Equal([]string{
					"ERROR getting user policy the-policy: some error\n",
					"SUCCESS deleting user policy the-policy\n",
				}))
			})
		})

		Context("when the client fails to detach the user policy", func() {
			BeforeEach(func() {
				client.DetachUserPolicyCall.Returns.Error = errors.New("some error")
			})

			It("logs the error and deletes the user policy", func() {
				err := policies.Delete("banana")
				Expect(err).NotTo(HaveOccurred())

				Expect(client.DeleteUserPolicyCall.CallCount).To(Equal(1))
				Expect(logger.PrintfCall.Messages).To(Equal([]string{
					"ERROR detaching user policy the-policy: some error\n",
					"SUCCESS deleting user policy the-policy\n",
				}))
			})
		})

		Context("when the client fails to delete the user policy", func() {
			BeforeEach(func() {
				client.DeleteUserPolicyCall.Returns.Error = errors.New("some error")
			})

			It("logs the error", func() {
				err := policies.Delete("banana")
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PrintfCall.Messages).To(Equal([]string{
					"SUCCESS detaching user policy the-policy\n",
					"ERROR deleting user policy the-policy: some error\n",
				}))
			})
		})

		Context("when the user responds no to the prompt", func() {
			BeforeEach(func() {
				logger.PromptCall.Returns.Proceed = false
			})

			It("does not delete the user policy", func() {
				err := policies.Delete("banana")
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete user policy the-policy?"))
				Expect(client.DeleteUserPolicyCall.CallCount).To(Equal(0))
			})
		})
	})
})
