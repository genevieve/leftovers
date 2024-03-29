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

var _ = Describe("InstanceProfile", func() {
	var (
		instanceProfile iam.InstanceProfile
		client          *fakes.InstanceProfilesClient
		name            *string
		logger          *fakes.Logger
		messages        []string
	)

	BeforeEach(func() {
		client = &fakes.InstanceProfilesClient{}
		name = aws.String("the-name")
		roles := []*awsiam.Role{}

		messages = []string{}
		logger = &fakes.Logger{}
		logger.PrintfCall.Stub = func(format string, v ...interface{}) {
			messages = append(messages, fmt.Sprintf(format, v...))
		}

		instanceProfile = iam.NewInstanceProfile(client, name, roles, logger)
	})

	Describe("Delete", func() {
		It("deletes the instance profile", func() {
			err := instanceProfile.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(client.DeleteInstanceProfileCall.CallCount).To(Equal(1))
			Expect(client.DeleteInstanceProfileCall.Receives.DeleteInstanceProfileInput.InstanceProfileName).To(Equal(name))
		})

		Context("when the client fails", func() {
			BeforeEach(func() {
				client.DeleteInstanceProfileCall.Returns.Error = errors.New("banana")
			})

			It("returns the error", func() {
				err := instanceProfile.Delete()
				Expect(err).To(MatchError("Delete: banana"))
			})
		})

		Context("when there are roles", func() {
			BeforeEach(func() {
				roles := []*awsiam.Role{{RoleName: aws.String("the-role")}}
				instanceProfile = iam.NewInstanceProfile(client, name, roles, logger)
			})

			It("removes the roles and uses them in the name", func() {
				err := instanceProfile.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(client.RemoveRoleFromInstanceProfileCall.CallCount).To(Equal(1))
				Expect(client.RemoveRoleFromInstanceProfileCall.Receives.RemoveRoleFromInstanceProfileInput.InstanceProfileName).To(Equal(aws.String("the-name")))
				Expect(client.RemoveRoleFromInstanceProfileCall.Receives.RemoveRoleFromInstanceProfileInput.RoleName).To(Equal(aws.String("the-role")))

				Expect(messages).To(Equal([]string{
					"[IAM Instance Profile: the-name (Role:the-role)] Removed role the-role \n",
				}))
			})

			Context("when the client fails to remove the role from the instance profile", func() {
				BeforeEach(func() {
					client.RemoveRoleFromInstanceProfileCall.Returns.Error = errors.New("some error")
				})

				It("logs the error", func() {
					err := instanceProfile.Delete()
					Expect(err).NotTo(HaveOccurred())

					Expect(messages).To(Equal([]string{
						"[IAM Instance Profile: the-name (Role:the-role)] Remove role the-role: some error \n",
					}))
				})
			})
		})
	})

	Describe("Name", func() {
		It("returns the identifier", func() {
			Expect(instanceProfile.Name()).To(Equal("the-name"))
		})
	})

	Describe("Type", func() {
		It("returns the type", func() {
			Expect(instanceProfile.Type()).To(Equal("IAM Instance Profile"))
		})
	})
})
