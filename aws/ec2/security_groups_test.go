package ec2_test

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
	"github.com/genevievelesperance/leftovers/aws/ec2"
	"github.com/genevievelesperance/leftovers/aws/ec2/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SecurityGroups", func() {
	var (
		ec2Client *fakes.EC2Client
		logger    *fakes.Logger

		securityGroups ec2.SecurityGroups
	)

	BeforeEach(func() {
		ec2Client = &fakes.EC2Client{}
		logger = &fakes.Logger{}
		logger.PromptCall.Returns.Proceed = true

		securityGroups = ec2.NewSecurityGroups(ec2Client, logger)
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			ec2Client.DescribeSecurityGroupsCall.Returns.Output = &awsec2.DescribeSecurityGroupsOutput{
				SecurityGroups: []*awsec2.SecurityGroup{{
					GroupName: aws.String("banana"),
					GroupId:   aws.String("the-group-id"),
				}},
			}
		})

		It("deletes ec2 security groups", func() {
			err := securityGroups.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(ec2Client.DescribeSecurityGroupsCall.CallCount).To(Equal(1))

			Expect(ec2Client.DeleteSecurityGroupCall.CallCount).To(Equal(1))
			Expect(ec2Client.DeleteSecurityGroupCall.Receives.Input.GroupId).To(Equal(aws.String("the-group-id")))

			Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete security group banana?"))
			Expect(logger.PrintfCall.Messages).To(Equal([]string{"SUCCESS deleting security group banana\n"}))
		})

		Context("when the client fails to describe security groups", func() {
			BeforeEach(func() {
				ec2Client.DescribeSecurityGroupsCall.Returns.Error = errors.New("some error")
			})

			It("does not try deleting them", func() {
				err := securityGroups.Delete()
				Expect(err.Error()).To(Equal("Describing security groups: some error"))

				Expect(ec2Client.DeleteSecurityGroupCall.CallCount).To(Equal(0))
			})
		})

		Context("when the client has ingress rules", func() {
			BeforeEach(func() {
				ec2Client.DescribeSecurityGroupsCall.Returns.Output = &awsec2.DescribeSecurityGroupsOutput{
					SecurityGroups: []*awsec2.SecurityGroup{{
						GroupName: aws.String("banana"),
						GroupId:   aws.String("the-group-id"),
						IpPermissions: []*awsec2.IpPermission{{
							IpProtocol: aws.String("tcp"),
						}},
					}},
				}
			})

			It("revokes them", func() {
				err := securityGroups.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(ec2Client.DescribeSecurityGroupsCall.CallCount).To(Equal(1))

				Expect(ec2Client.RevokeSecurityGroupIngressCall.CallCount).To(Equal(1))
				Expect(ec2Client.RevokeSecurityGroupIngressCall.Receives.Input.GroupId).To(Equal(aws.String("the-group-id")))
				Expect(ec2Client.RevokeSecurityGroupIngressCall.Receives.Input.IpPermissions[0].IpProtocol).To(Equal(aws.String("tcp")))

				Expect(ec2Client.DeleteSecurityGroupCall.CallCount).To(Equal(1))
				Expect(ec2Client.DeleteSecurityGroupCall.Receives.Input.GroupId).To(Equal(aws.String("the-group-id")))

				Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete security group banana?"))
				Expect(logger.PrintfCall.Messages).To(Equal([]string{"SUCCESS deleting security group banana\n"}))
			})

			Context("when the client fails to revoke ingress rules", func() {
				BeforeEach(func() {
					ec2Client.RevokeSecurityGroupIngressCall.Returns.Error = errors.New("some error")
				})

				It("logs the error", func() {
					err := securityGroups.Delete()
					Expect(err).NotTo(HaveOccurred())

					Expect(logger.PrintfCall.Messages).To(Equal([]string{
						"ERROR revoking security group ingress for banana: some error\n",
						"SUCCESS deleting security group banana\n",
					}))
				})
			})
		})

		Context("when the client has egress rules", func() {
			BeforeEach(func() {
				ec2Client.DescribeSecurityGroupsCall.Returns.Output = &awsec2.DescribeSecurityGroupsOutput{
					SecurityGroups: []*awsec2.SecurityGroup{{
						GroupName: aws.String("banana"),
						GroupId:   aws.String("the-group-id"),
						IpPermissionsEgress: []*awsec2.IpPermission{{
							IpProtocol: aws.String("tcp"),
						}},
					}},
				}
			})

			It("revokes them", func() {
				err := securityGroups.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(ec2Client.DescribeSecurityGroupsCall.CallCount).To(Equal(1))

				Expect(ec2Client.RevokeSecurityGroupEgressCall.CallCount).To(Equal(1))
				Expect(ec2Client.RevokeSecurityGroupEgressCall.Receives.Input.GroupId).To(Equal(aws.String("the-group-id")))
				Expect(ec2Client.RevokeSecurityGroupEgressCall.Receives.Input.IpPermissions[0].IpProtocol).To(Equal(aws.String("tcp")))

				Expect(ec2Client.DeleteSecurityGroupCall.CallCount).To(Equal(1))
				Expect(ec2Client.DeleteSecurityGroupCall.Receives.Input.GroupId).To(Equal(aws.String("the-group-id")))

				Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete security group banana?"))
				Expect(logger.PrintfCall.Messages).To(Equal([]string{"SUCCESS deleting security group banana\n"}))
			})

			Context("when the client fails to revoke egress rules", func() {
				BeforeEach(func() {
					ec2Client.RevokeSecurityGroupEgressCall.Returns.Error = errors.New("some error")
				})

				It("logs the error", func() {
					err := securityGroups.Delete()
					Expect(err).NotTo(HaveOccurred())

					Expect(logger.PrintfCall.Messages).To(Equal([]string{
						"ERROR revoking security group egress for banana: some error\n",
						"SUCCESS deleting security group banana\n",
					}))
				})
			})
		})

		Context("when the client fails to delete the security group", func() {
			BeforeEach(func() {
				ec2Client.DeleteSecurityGroupCall.Returns.Error = errors.New("some error")
			})

			It("returns the error", func() {
				err := securityGroups.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PrintfCall.Messages).To(Equal([]string{"ERROR deleting security group banana: some error\n"}))
			})
		})

		Context("when the user responds no to the prompt", func() {
			BeforeEach(func() {
				logger.PromptCall.Returns.Proceed = false
			})

			It("does not delete the security group", func() {
				err := securityGroups.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete security group banana?"))
				Expect(ec2Client.DeleteSecurityGroupCall.CallCount).To(Equal(0))
			})
		})
	})
})
