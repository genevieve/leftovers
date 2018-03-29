package iam_test

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	awsiam "github.com/aws/aws-sdk-go/service/iam"
	"github.com/genevieve/leftovers/aws/iam"
	"github.com/genevieve/leftovers/aws/iam/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RolePolicies", func() {
	var (
		client *fakes.RolePoliciesClient
		logger *fakes.Logger

		policies iam.RolePolicies
	)

	BeforeEach(func() {
		client = &fakes.RolePoliciesClient{}
		logger = &fakes.Logger{}

		policies = iam.NewRolePolicies(client, logger)
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			client.ListAttachedRolePoliciesCall.Returns.Output = &awsiam.ListAttachedRolePoliciesOutput{
				AttachedPolicies: []*awsiam.AttachedPolicy{{
					PolicyName: aws.String("the-policy"),
					PolicyArn:  aws.String("the-policy-arn"),
				}},
			}
			client.ListRolePoliciesCall.Returns.Output = &awsiam.ListRolePoliciesOutput{}
		})

		It("detaches and deletes the attached policies", func() {
			err := policies.Delete("banana")
			Expect(err).NotTo(HaveOccurred())

			Expect(client.ListAttachedRolePoliciesCall.CallCount).To(Equal(1))
			Expect(client.ListAttachedRolePoliciesCall.Receives.Input.RoleName).To(Equal(aws.String("banana")))

			Expect(client.DetachRolePolicyCall.CallCount).To(Equal(1))
			Expect(client.DetachRolePolicyCall.Receives.Input.RoleName).To(Equal(aws.String("banana")))
			Expect(client.DetachRolePolicyCall.Receives.Input.PolicyArn).To(Equal(aws.String("the-policy-arn")))

			Expect(client.DeleteRolePolicyCall.CallCount).To(Equal(1))
			Expect(client.DeleteRolePolicyCall.Receives.Input.RoleName).To(Equal(aws.String("banana")))
			Expect(client.DeleteRolePolicyCall.Receives.Input.PolicyName).To(Equal(aws.String("the-policy")))

			Expect(logger.PrintfCall.Messages).To(Equal([]string{
				"[INFO] Detached IAM Role Policy the-policy for IAM Role banana\n",
				"[INFO] Deleted IAM Role Policy the-policy for IAM Role banana\n",
			}))
		})

		Context("when the policies are not attached", func() {
			BeforeEach(func() {
				client.ListAttachedRolePoliciesCall.Returns.Output = &awsiam.ListAttachedRolePoliciesOutput{}
				client.ListRolePoliciesCall.Returns.Output = &awsiam.ListRolePoliciesOutput{
					PolicyNames: []*string{aws.String("the-not-attached-policy")},
				}
			})

			It("deletes the policies", func() {
				err := policies.Delete("banana")
				Expect(err).NotTo(HaveOccurred())

				Expect(client.ListRolePoliciesCall.CallCount).To(Equal(1))
				Expect(client.ListRolePoliciesCall.Receives.Input.RoleName).To(Equal(aws.String("banana")))

				Expect(client.DetachRolePolicyCall.CallCount).To(Equal(0))

				Expect(client.DeleteRolePolicyCall.CallCount).To(Equal(1))
				Expect(client.DeleteRolePolicyCall.Receives.Input.RoleName).To(Equal(aws.String("banana")))
				Expect(client.DeleteRolePolicyCall.Receives.Input.PolicyName).To(Equal(aws.String("the-not-attached-policy")))

				Expect(logger.PrintfCall.Messages).To(Equal([]string{
					"[INFO] Deleted IAM Role Policy the-not-attached-policy for IAM Role banana\n",
				}))
			})
		})

		Context("when the client fails to list attached role policies", func() {
			BeforeEach(func() {
				client.ListAttachedRolePoliciesCall.Returns.Error = errors.New("some error")
				client.ListRolePoliciesCall.Returns.Output = &awsiam.ListRolePoliciesOutput{}
			})

			It("returns the error and does not try deleting them", func() {
				err := policies.Delete("banana")
				Expect(err).To(MatchError("List IAM Attached Role Policies: some error"))

				Expect(client.DetachRolePolicyCall.CallCount).To(Equal(0))
				Expect(client.DeleteRolePolicyCall.CallCount).To(Equal(0))
			})
		})

		Context("when the client fails to list role policies", func() {
			BeforeEach(func() {
				client.ListAttachedRolePoliciesCall.Returns.Output = &awsiam.ListAttachedRolePoliciesOutput{}
				client.ListRolePoliciesCall.Returns.Error = errors.New("some error")
			})

			It("returns the error and does not try deleting them", func() {
				err := policies.Delete("banana")
				Expect(err).To(MatchError("List IAM Role Policies: some error"))

				Expect(client.DeleteRolePolicyCall.CallCount).To(Equal(0))
			})
		})

		Context("when the client fails to detach the role policy", func() {
			BeforeEach(func() {
				client.DetachRolePolicyCall.Returns.Error = errors.New("some error")
			})

			It("logs the error and deletes the role policy", func() {
				err := policies.Delete("banana")
				Expect(err).NotTo(HaveOccurred())

				Expect(client.DeleteRolePolicyCall.CallCount).To(Equal(1))
				Expect(logger.PrintfCall.Messages).To(Equal([]string{
					"[WARNING] Detach IAM Role Policy the-policy for IAM Role banana: some error\n",
					"[INFO] Deleted IAM Role Policy the-policy for IAM Role banana\n",
				}))
			})
		})

		Context("when the client fails to delete the role policy", func() {
			BeforeEach(func() {
				client.DeleteRolePolicyCall.Returns.Error = errors.New("some error")
			})

			It("logs the error", func() {
				err := policies.Delete("banana")
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PrintfCall.Messages).To(Equal([]string{
					"[INFO] Detached IAM Role Policy the-policy for IAM Role banana\n",
					"[WARNING] Delete IAM Role Policy the-policy for IAM Role banana: some error\n",
				}))
			})
		})
	})
})
