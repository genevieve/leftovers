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

var _ = Describe("Roles", func() {
	var (
		client   *fakes.RolesClient
		logger   *fakes.Logger
		policies *fakes.RolePolicies

		roles iam.Roles
	)

	BeforeEach(func() {
		client = &fakes.RolesClient{}
		logger = &fakes.Logger{}
		policies = &fakes.RolePolicies{}

		roles = iam.NewRoles(client, logger, policies)
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			logger.PromptCall.Returns.Proceed = true
			client.ListRolesCall.Returns.Output = &awsiam.ListRolesOutput{
				Roles: []*awsiam.Role{{
					RoleName: aws.String("banana"),
				}},
			}
		})

		It("deletes iam roles and associated policies", func() {
			err := roles.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(client.ListRolesCall.CallCount).To(Equal(1))
			Expect(policies.DeleteCall.CallCount).To(Equal(1))
			Expect(policies.DeleteCall.Receives.RoleName).To(Equal("banana"))
			Expect(client.DeleteRoleCall.CallCount).To(Equal(1))
			Expect(client.DeleteRoleCall.Receives.Input.RoleName).To(Equal(aws.String("banana")))
			Expect(logger.PrintfCall.Messages).To(Equal([]string{"SUCCESS deleting role banana\n"}))
		})

		Context("when the client fails to list roles", func() {
			BeforeEach(func() {
				client.ListRolesCall.Returns.Error = errors.New("some error")
			})

			It("returns the error and does not try deleting them", func() {
				err := roles.Delete()
				Expect(err).To(MatchError("Listing roles: some error"))

				Expect(client.DeleteRoleCall.CallCount).To(Equal(0))
			})
		})

		Context("when policies fails to delete", func() {
			BeforeEach(func() {
				policies.DeleteCall.Returns.Error = errors.New("some error")
			})

			It("returns the error", func() {
				err := roles.Delete()
				Expect(err).To(MatchError("Deleting policies for banana: some error"))

				Expect(policies.DeleteCall.CallCount).To(Equal(1))
			})
		})

		Context("when the client fails to delete the role", func() {
			BeforeEach(func() {
				client.DeleteRoleCall.Returns.Error = errors.New("some error")
			})

			It("logs the error", func() {
				err := roles.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PrintfCall.Messages).To(Equal([]string{"ERROR deleting role banana: some error\n"}))
			})
		})

		Context("when the user responds no to the prompt", func() {
			BeforeEach(func() {
				logger.PromptCall.Returns.Proceed = false
			})

			It("does not delete the role", func() {
				err := roles.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete role banana?"))
				Expect(client.DeleteRoleCall.CallCount).To(Equal(0))
			})
		})
	})
})
