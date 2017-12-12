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

var _ = Describe("Vpcs", func() {
	var (
		ec2Client *fakes.VpcClient
		logger    *fakes.Logger

		vpcs ec2.Vpcs
	)

	BeforeEach(func() {
		ec2Client = &fakes.VpcClient{}
		logger = &fakes.Logger{}
		logger.PromptCall.Returns.Proceed = true

		vpcs = ec2.NewVpcs(ec2Client, logger)
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			ec2Client.DescribeVpcsCall.Returns.Output = &awsec2.DescribeVpcsOutput{
				Vpcs: []*awsec2.Vpc{{
					Tags: []*awsec2.Tag{{
						Key:   aws.String("Name"),
						Value: aws.String("banana"),
					}},
					VpcId: aws.String("the-vpc-id"),
				}},
			}
		})

		It("deletes ec2 vpcs", func() {
			err := vpcs.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(ec2Client.DescribeVpcsCall.CallCount).To(Equal(1))
			Expect(ec2Client.DeleteVpcCall.CallCount).To(Equal(1))
			Expect(ec2Client.DeleteVpcCall.Receives.Input.VpcId).To(Equal(aws.String("the-vpc-id")))
			Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete vpc the-vpc-id/banana?"))
			Expect(logger.PrintfCall.Messages).To(Equal([]string{"SUCCESS deleting vpc the-vpc-id/banana\n"}))
		})

		Context("when there is no tag name", func() {
			BeforeEach(func() {
				ec2Client.DescribeVpcsCall.Returns.Output = &awsec2.DescribeVpcsOutput{
					Vpcs: []*awsec2.Vpc{{
						VpcId: aws.String("the-vpc-id"),
					}},
				}
			})

			It("uses just the vpc id in the prompt", func() {
				err := vpcs.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete vpc the-vpc-id?"))
				Expect(logger.PrintfCall.Messages).To(Equal([]string{"SUCCESS deleting vpc the-vpc-id\n"}))
			})
		})

		Context("when the client fails to list vpcs", func() {
			BeforeEach(func() {
				ec2Client.DescribeVpcsCall.Returns.Error = errors.New("some error")
			})

			It("does not try deleting them", func() {
				err := vpcs.Delete()
				Expect(err.Error()).To(Equal("Describing vpcs: some error"))

				Expect(ec2Client.DeleteVpcCall.CallCount).To(Equal(0))
			})
		})

		Context("when the client fails to delete the vpc", func() {
			BeforeEach(func() {
				ec2Client.DeleteVpcCall.Returns.Error = errors.New("some error")
			})

			It("logs the error", func() {
				err := vpcs.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PrintfCall.Messages).To(Equal([]string{"ERROR deleting vpc the-vpc-id/banana: some error\n"}))
			})
		})

		Context("when the user responds no to the prompt", func() {
			BeforeEach(func() {
				logger.PromptCall.Returns.Proceed = false
			})

			It("does not delete the vpc", func() {
				err := vpcs.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete vpc the-vpc-id/banana?"))
				Expect(ec2Client.DeleteVpcCall.CallCount).To(Equal(0))
			})
		})
	})
})
