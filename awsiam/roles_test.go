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
		iamClient        *fakes.IAMClient
		instanceProfiles awsiam.Roles
	)

	BeforeEach(func() {
		iamClient = &fakes.IAMClient{}
		instanceProfiles = awsiam.NewRoles(iamClient)
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
			instanceProfiles.Delete()

			Expect(iamClient.DeleteRoleCall.CallCount).To(Equal(1))
			Expect(iamClient.DeleteRoleCall.Receives.Input.RoleName).To(Equal(aws.String("banana")))
		})

		Context("when the client fails to list roles", func() {
			BeforeEach(func() {
				iamClient.ListRolesCall.Returns.Error = errors.New("some error")
				iamClient.ListRolesCall.Returns.Output = &iam.ListRolesOutput{}
			})

			It("does not try deleting them", func() {
				instanceProfiles.Delete()

				Expect(iamClient.DeleteRoleCall.CallCount).To(Equal(0))
			})
		})
	})
})
