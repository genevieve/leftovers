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
		client *fakes.IAMClient
		logger *fakes.Logger

		policies iam.RolePolicies
	)

	BeforeEach(func() {
		client = &fakes.IAMClient{}
		logger = &fakes.Logger{}

		policies = iam.NewRolePolicies(client, logger)
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			logger.PromptCall.Returns.Proceed = true
			client.ListRolePoliciesCall.Returns.Output = &awsiam.ListRolePoliciesOutput{
				PolicyNames: []*string{aws.String("the-policy")},
			}
		})

		It("deletes the policies", func() {
			err := policies.Delete("the-role")
			Expect(err).NotTo(HaveOccurred())

			Expect(client.ListRolePoliciesCall.CallCount).To(Equal(1))
			Expect(client.ListRolePoliciesCall.Receives.Input.RoleName).To(Equal(aws.String("the-role")))
			Expect(client.DeleteRolePolicyCall.CallCount).To(Equal(1))
			Expect(client.DeleteRolePolicyCall.Receives.Input.RoleName).To(Equal(aws.String("the-role")))
			Expect(client.DeleteRolePolicyCall.Receives.Input.PolicyName).To(Equal(aws.String("the-policy")))
			Expect(logger.PrintfCall.Messages).To(Equal([]string{"SUCCESS deleting role policy the-policy\n"}))
		})

		Context("when the client fails to list role policies", func() {
			BeforeEach(func() {
				client.ListRolePoliciesCall.Returns.Error = errors.New("some error")
			})

			It("returns the error and does not try deleting them", func() {
				err := policies.Delete("banana")
				Expect(err.Error()).To(Equal("Listing role policies: some error"))

				Expect(client.DeleteRolePolicyCall.CallCount).To(Equal(0))
			})
		})

		Context("when the client fails to delete the role policy", func() {
			BeforeEach(func() {
				client.DeleteRolePolicyCall.Returns.Error = errors.New("some error")
			})

			It("logs the error", func() {
				err := policies.Delete("banana")
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PrintfCall.Messages).To(Equal([]string{"ERROR deleting role policy the-policy: some error\n"}))
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
