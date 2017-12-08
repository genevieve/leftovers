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

var _ = Describe("InstanceProfiles", func() {
	var (
		iamClient        *fakes.IAMClient
		instanceProfiles awsiam.InstanceProfiles
	)

	BeforeEach(func() {
		iamClient = &fakes.IAMClient{}
		instanceProfiles = awsiam.NewInstanceProfiles(iamClient)
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			iamClient.ListInstanceProfilesCall.Returns.Output = &iam.ListInstanceProfilesOutput{
				InstanceProfiles: []*iam.InstanceProfile{{
					InstanceProfileName: aws.String("banana"),
				}},
			}
		})

		It("deletes iam instance profiles", func() {
			instanceProfiles.Delete()

			Expect(iamClient.DeleteInstanceProfileCall.CallCount).To(Equal(1))
			Expect(iamClient.DeleteInstanceProfileCall.Receives.Input.InstanceProfileName).To(Equal(aws.String("banana")))
		})

		Context("when the client fails to list instance profiles", func() {
			BeforeEach(func() {
				iamClient.ListInstanceProfilesCall.Returns.Error = errors.New("some error")
			})

			It("does not try deleting them", func() {
				instanceProfiles.Delete()

				Expect(iamClient.DeleteInstanceProfileCall.CallCount).To(Equal(0))
			})
		})
	})
})
