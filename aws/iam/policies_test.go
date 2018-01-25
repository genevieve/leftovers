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

var _ = Describe("Policies", func() {
	var (
		client *fakes.PoliciesClient
		logger *fakes.Logger

		policies iam.Policies
	)

	BeforeEach(func() {
		client = &fakes.PoliciesClient{}
		logger = &fakes.Logger{}

		policies = iam.NewPolicies(client, logger)
	})

	Describe("Delete", func() {
		var filter string

		BeforeEach(func() {
			logger.PromptCall.Returns.Proceed = true
			client.ListPoliciesCall.Returns.Output = &awsiam.ListPoliciesOutput{
				Policies: []*awsiam.Policy{{
					Arn:        aws.String("the-policy-arn"),
					PolicyName: aws.String("banana-policy"),
				}},
			}
			filter = "banana"
		})

		It("deletes iam policies and associated policies", func() {
			err := policies.Delete(filter)
			Expect(err).NotTo(HaveOccurred())

			Expect(client.ListPoliciesCall.CallCount).To(Equal(1))

			Expect(client.DeletePolicyCall.CallCount).To(Equal(1))
			Expect(client.DeletePolicyCall.Receives.Input.PolicyArn).To(Equal(aws.String("the-policy-arn")))

			Expect(logger.PrintfCall.Messages).To(Equal([]string{"SUCCESS deleting policy banana-policy\n"}))
		})

		Context("when the client fails to list policies", func() {
			BeforeEach(func() {
				client.ListPoliciesCall.Returns.Error = errors.New("some error")
			})

			It("returns the error and does not try deleting them", func() {
				err := policies.Delete(filter)
				Expect(err).To(MatchError("Listing policies: some error"))

				Expect(client.DeletePolicyCall.CallCount).To(Equal(0))
			})
		})

		Context("when the policy name does not contain the filter", func() {
			It("does not try to delete it", func() {
				err := policies.Delete("kiwi")
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PromptCall.CallCount).To(Equal(0))
				Expect(client.DeletePolicyCall.CallCount).To(Equal(0))
			})
		})

		Context("when the client fails to delete the policy", func() {
			BeforeEach(func() {
				client.DeletePolicyCall.Returns.Error = errors.New("some error")
			})

			It("logs the error", func() {
				err := policies.Delete(filter)
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PrintfCall.Messages).To(Equal([]string{"ERROR deleting policy banana-policy: some error\n"}))
			})
		})

		Context("when the user responds no to the prompt", func() {
			BeforeEach(func() {
				logger.PromptCall.Returns.Proceed = false
			})

			It("does not delete the policy", func() {
				err := policies.Delete(filter)
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete policy banana-policy?"))
				Expect(client.DeletePolicyCall.CallCount).To(Equal(0))
			})
		})
	})
})
