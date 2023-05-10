package iam_test

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	awsiam "github.com/aws/aws-sdk-go/service/iam"
	"github.com/genevieve/leftovers/aws/iam"
	"github.com/genevieve/leftovers/aws/iam/fakes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Policy", func() {
	var (
		policy   iam.Policy
		client   *fakes.PoliciesClient
		logger   *fakes.Logger
		messages []string
		name     *string
		arn      *string
	)

	BeforeEach(func() {
		client = &fakes.PoliciesClient{}
		name = aws.String("banana")
		arn = aws.String("the-arn")

		messages = []string{}
		logger = &fakes.Logger{}
		logger.PrintfCall.Stub = func(format string, v ...interface{}) {
			messages = append(messages, fmt.Sprintf(format, v...))
		}

		policy = iam.NewPolicy(client, logger, name, arn)

		client.ListPolicyVersionsCall.Returns.ListPolicyVersionsOutput = &awsiam.ListPolicyVersionsOutput{
			Versions: []*awsiam.PolicyVersion{},
		}
	})

	Describe("Delete", func() {
		It("deletes the policy", func() {
			err := policy.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(client.DeletePolicyCall.CallCount).To(Equal(1))
			Expect(client.DeletePolicyCall.Receives.DeletePolicyInput.PolicyArn).To(Equal(arn))
		})

		Context("when the policy has non-default versions", func() {
			BeforeEach(func() {
				client.ListPolicyVersionsCall.Returns.ListPolicyVersionsOutput = &awsiam.ListPolicyVersionsOutput{
					Versions: []*awsiam.PolicyVersion{
						{IsDefaultVersion: aws.Bool(true), VersionId: aws.String("v2")},
						{IsDefaultVersion: aws.Bool(false), VersionId: aws.String("v1")},
					},
				}
			})

			It("deletes all non-default versions", func() {
				err := policy.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(client.ListPolicyVersionsCall.CallCount).To(Equal(1))

				Expect(client.DeletePolicyVersionCall.CallCount).To(Equal(1))
				Expect(client.DeletePolicyVersionCall.Receives.DeletePolicyVersionInput.PolicyArn).To(Equal(arn))
				Expect(client.DeletePolicyVersionCall.Receives.DeletePolicyVersionInput.VersionId).To(Equal(aws.String("v1")))
			})

			Context("when the client fails to delete policy versions", func() {
				BeforeEach(func() {
					client.DeletePolicyVersionCall.Returns.Error = errors.New("some error")
				})

				It("logs the error", func() {
					err := policy.Delete()
					Expect(err).NotTo(HaveOccurred())

					Expect(messages).To(Equal([]string{
						"[IAM Policy: banana] Delete policy version v1: some error \n",
					}))
				})
			})
		})

		Context("when the client fails to delete the policy", func() {
			BeforeEach(func() {
				client.DeletePolicyCall.Returns.Error = errors.New("some error")
			})

			It("returns the error", func() {
				err := policy.Delete()
				Expect(err).To(MatchError("Delete: some error"))
			})
		})

		Context("when the client fails to list policy versions", func() {
			BeforeEach(func() {
				client.ListPolicyVersionsCall.Returns.Error = errors.New("some error")
			})

			It("returns the error", func() {
				err := policy.Delete()
				Expect(err).To(MatchError("List IAM Policy Versions: some error"))
			})
		})
	})

	Describe("Name", func() {
		It("returns the identifier", func() {
			Expect(policy.Name()).To(Equal("banana"))
		})
	})

	Describe("Type", func() {
		It("returns \"policy\"", func() {
			Expect(policy.Type()).To(Equal("IAM Policy"))
		})
	})
})
