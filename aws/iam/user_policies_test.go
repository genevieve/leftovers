package iam_test

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	awsiam "github.com/aws/aws-sdk-go/service/iam"
	"github.com/genevieve/leftovers/aws/iam"
	"github.com/genevieve/leftovers/aws/iam/fakes"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("UserPolicies", func() {
	var (
		client   *fakes.UserPoliciesClient
		logger   *fakes.Logger
		messages []string

		policies iam.UserPolicies
	)

	BeforeEach(func() {
		client = &fakes.UserPoliciesClient{}
		messages = []string{}
		logger = &fakes.Logger{}
		logger.PrintfCall.Stub = func(format string, v ...interface{}) {
			messages = append(messages, fmt.Sprintf(format, v...))
		}

		policies = iam.NewUserPolicies(client, logger)
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			client.ListAttachedUserPoliciesCall.Returns.ListAttachedUserPoliciesOutput = &awsiam.ListAttachedUserPoliciesOutput{
				AttachedPolicies: []*awsiam.AttachedPolicy{{
					PolicyName: aws.String("the-policy"),
					PolicyArn:  aws.String("the-policy-arn"),
				}},
			}
			client.ListUserPoliciesCall.Returns.ListUserPoliciesOutput = &awsiam.ListUserPoliciesOutput{
				PolicyNames: []*string{},
			}
		})

		It("detaches attached policies and deletes them", func() {
			err := policies.Delete("banana")
			Expect(err).NotTo(HaveOccurred())

			Expect(client.ListAttachedUserPoliciesCall.CallCount).To(Equal(1))
			Expect(client.ListAttachedUserPoliciesCall.Receives.ListAttachedUserPoliciesInput.UserName).To(Equal(aws.String("banana")))

			Expect(client.DetachUserPolicyCall.CallCount).To(Equal(1))
			Expect(client.DetachUserPolicyCall.Receives.DetachUserPolicyInput.UserName).To(Equal(aws.String("banana")))
			Expect(client.DetachUserPolicyCall.Receives.DetachUserPolicyInput.PolicyArn).To(Equal(aws.String("the-policy-arn")))

			Expect(client.DeleteUserPolicyCall.CallCount).To(Equal(1))
			Expect(client.DeleteUserPolicyCall.Receives.DeleteUserPolicyInput.UserName).To(Equal(aws.String("banana")))
			Expect(client.DeleteUserPolicyCall.Receives.DeleteUserPolicyInput.PolicyName).To(Equal(aws.String("the-policy")))

			Expect(messages).To(Equal([]string{
				"[IAM User: banana] Detached policy the-policy \n",
				"[IAM User: banana] Deleted policy the-policy \n",
			}))
		})

		Context("when there are unattached user policies", func() {
			BeforeEach(func() {
				client.ListAttachedUserPoliciesCall.Returns.ListAttachedUserPoliciesOutput = &awsiam.ListAttachedUserPoliciesOutput{
					AttachedPolicies: []*awsiam.AttachedPolicy{},
				}
				client.ListUserPoliciesCall.Returns.ListUserPoliciesOutput = &awsiam.ListUserPoliciesOutput{
					PolicyNames: []*string{aws.String("the-other-policy")},
				}
			})

			It("deletes them", func() {
				err := policies.Delete("banana")
				Expect(err).NotTo(HaveOccurred())

				Expect(client.ListUserPoliciesCall.CallCount).To(Equal(1))
				Expect(client.ListUserPoliciesCall.Receives.ListUserPoliciesInput.UserName).To(Equal(aws.String("banana")))

				Expect(client.DetachUserPolicyCall.CallCount).To(Equal(0))

				Expect(client.DeleteUserPolicyCall.CallCount).To(Equal(1))
				Expect(client.DeleteUserPolicyCall.Receives.DeleteUserPolicyInput.UserName).To(Equal(aws.String("banana")))
				Expect(client.DeleteUserPolicyCall.Receives.DeleteUserPolicyInput.PolicyName).To(Equal(aws.String("the-other-policy")))

				Expect(messages).To(Equal([]string{
					"[IAM User: banana] Deleted policy the-other-policy \n",
				}))
			})
		})

		Context("when the client fails to list attached user policies", func() {
			BeforeEach(func() {
				client.ListAttachedUserPoliciesCall.Returns.Error = errors.New("some error")
			})

			It("returns the error and does not try deleting them", func() {
				err := policies.Delete("banana")
				Expect(err).To(MatchError("List Attached User Policies: some error"))

				Expect(client.DetachUserPolicyCall.CallCount).To(Equal(0))
				Expect(client.DeleteUserPolicyCall.CallCount).To(Equal(0))
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
				Expect(messages).To(Equal([]string{
					"[IAM User: banana] Detach policy the-policy: some error \n",
					"[IAM User: banana] Deleted policy the-policy \n",
				}))
			})
		})

		Context("when the client fails to detach the user policy due to NoSuchEntity", func() {
			BeforeEach(func() {
				client.DetachUserPolicyCall.Returns.Error = awserr.New("NoSuchEntity", "hi", nil)
			})

			It("logs success", func() {
				err := policies.Delete("banana")
				Expect(err).NotTo(HaveOccurred())

				Expect(client.DeleteUserPolicyCall.CallCount).To(Equal(1))
				Expect(messages).To(Equal([]string{
					"[IAM User: banana] Detached policy the-policy \n",
					"[IAM User: banana] Deleted policy the-policy \n",
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

				Expect(messages).To(Equal([]string{
					"[IAM User: banana] Detached policy the-policy \n",
					"[IAM User: banana] Delete policy the-policy: some error \n",
				}))
			})
		})

		Context("when the client fails to delete the user policy due to NoSuchEntity", func() {
			BeforeEach(func() {
				client.DetachUserPolicyCall.Returns.Error = awserr.New("NoSuchEntity", "hi", nil)
				client.DeleteUserPolicyCall.Returns.Error = awserr.New("NoSuchEntity", "hi", nil)
			})

			It("logs success", func() {
				err := policies.Delete("banana")
				Expect(err).NotTo(HaveOccurred())

				Expect(client.DeleteUserPolicyCall.CallCount).To(Equal(1))
				Expect(messages).To(Equal([]string{
					"[IAM User: banana] Detached policy the-policy \n",
					"[IAM User: banana] Deleted policy the-policy \n",
				}))
			})
		})

		Context("when the client fails to list user policies", func() {
			BeforeEach(func() {
				client.ListAttachedUserPoliciesCall.Returns.ListAttachedUserPoliciesOutput = &awsiam.ListAttachedUserPoliciesOutput{
					AttachedPolicies: []*awsiam.AttachedPolicy{},
				}
				client.ListUserPoliciesCall.Returns.Error = errors.New("some error")
			})

			It("returns the error and does not try deleting them", func() {
				err := policies.Delete("banana")
				Expect(err).To(MatchError("List User Policies: some error"))

				Expect(client.DeleteUserPolicyCall.CallCount).To(Equal(0))
			})
		})
	})
})
