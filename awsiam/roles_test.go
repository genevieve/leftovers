package awsiam_test

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/genevievelesperance/leftovers/awsiam"
	"github.com/genevievelesperance/leftovers/awsiam/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Roles", func() {
	var (
		iamClient *fakes.IAMClient
		logger    *fakes.Logger

		instanceProfiles awsiam.Roles
	)

	BeforeEach(func() {
		iamClient = &fakes.IAMClient{}
		logger = &fakes.Logger{}

		instanceProfiles = awsiam.NewRoles(iamClient, logger)
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			iamClient.ListRolesCall.Returns.Output = &iam.ListRolesOutput{
				Roles: []*iam.Role{{
					RoleName: aws.String("banana"),
				}},
			}
		})

		It("deletes iam roles", func() {
			err := instanceProfiles.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(iamClient.DeleteRoleCall.CallCount).To(Equal(1))
			Expect(iamClient.DeleteRoleCall.Receives.Input.RoleName).To(Equal(aws.String("banana")))
			Expect(logger.PrintfCall.Messages).To(Equal([]string{"SUCCESS deleting role banana\n"}))
		})

		Context("when the client fails to list roles", func() {
			BeforeEach(func() {
				iamClient.ListRolesCall.Returns.Error = errors.New("some error")
			})

			It("does not try deleting them", func() {
				err := instanceProfiles.Delete()
				Expect(err.Error()).To(Equal("Listing roles: some error"))

				Expect(iamClient.DeleteRoleCall.CallCount).To(Equal(0))
			})
		})

		Context("when the client fails to delete the role", func() {
			BeforeEach(func() {
				iamClient.DeleteRoleCall.Returns.Error = errors.New("some error")
			})

			It("returns the error", func() {
				err := instanceProfiles.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PrintfCall.Messages).To(Equal([]string{"ERROR deleting role banana: some error\n"}))
			})
		})
	})
})
