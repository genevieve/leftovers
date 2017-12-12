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
			logger.PromptCall.Returns.Proceed = true
			client.ListAttachedRolePoliciesCall.Returns.Output = &awsiam.ListAttachedRolePoliciesOutput{
				AttachedPolicies: []*awsiam.AttachedPolicy{{
					PolicyName: aws.String("the-policy"),
					PolicyArn:  aws.String("the-policy-arn"),
				}},
			}
		})

		It("detaches and deletes the policies", func() {
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
				"SUCCESS detaching role policy the-policy\n",
				"SUCCESS deleting role policy the-policy\n",
			}))
		})

		Context("when the client fails to list attached role policies", func() {
			BeforeEach(func() {
				client.ListAttachedRolePoliciesCall.Returns.Error = errors.New("some error")
			})

			It("returns the error and does not try deleting them", func() {
				err := policies.Delete("banana")
				Expect(err).To(MatchError("Listing role policies: some error"))

				Expect(client.DetachRolePolicyCall.CallCount).To(Equal(0))
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
					"ERROR detaching role policy the-policy: some error\n",
					"SUCCESS deleting role policy the-policy\n",
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
					"SUCCESS detaching role policy the-policy\n",
					"ERROR deleting role policy the-policy: some error\n",
				}))
			})
		})

		Context("when the user responds no to the prompt", func() {
			BeforeEach(func() {
				logger.PromptCall.Returns.Proceed = false
			})

			It("does not delete the role policy", func() {
				err := policies.Delete("banana")
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete role policy the-policy?"))
				Expect(client.DeleteRolePolicyCall.CallCount).To(Equal(0))
			})
		})
	})
})
