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
		client   *fakes.VpcClient
		logger   *fakes.Logger
		subnets  *fakes.Subnets
		gateways *fakes.InternetGateways

		vpcs ec2.Vpcs
	)

	BeforeEach(func() {
		client = &fakes.VpcClient{}
		logger = &fakes.Logger{}
		subnets = &fakes.Subnets{}
		gateways = &fakes.InternetGateways{}

		vpcs = ec2.NewVpcs(client, logger, subnets, gateways)
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			logger.PromptCall.Returns.Proceed = true
			client.DescribeVpcsCall.Returns.Output = &awsec2.DescribeVpcsOutput{
				Vpcs: []*awsec2.Vpc{{
					IsDefault: aws.Bool(false),
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

			Expect(client.DescribeVpcsCall.CallCount).To(Equal(1))

			Expect(subnets.DeleteCall.CallCount).To(Equal(1))
			Expect(subnets.DeleteCall.Receives.VpcId).To(Equal("the-vpc-id"))

			Expect(gateways.DeleteCall.CallCount).To(Equal(1))
			Expect(gateways.DeleteCall.Receives.VpcId).To(Equal("the-vpc-id"))

			Expect(client.DeleteVpcCall.CallCount).To(Equal(1))
			Expect(client.DeleteVpcCall.Receives.Input.VpcId).To(Equal(aws.String("the-vpc-id")))

			Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete vpc the-vpc-id/banana?"))
			Expect(logger.PrintfCall.Messages).To(Equal([]string{"SUCCESS deleting vpc the-vpc-id/banana\n"}))
		})

		Context("when the vpc is a default", func() {
			BeforeEach(func() {
				client.DescribeVpcsCall.Returns.Output = &awsec2.DescribeVpcsOutput{
					Vpcs: []*awsec2.Vpc{{
						IsDefault: aws.Bool(true),
						VpcId:     aws.String("the-vpc-id"),
					}},
				}
			})

			It("does not try deleting it", func() {
				err := vpcs.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(gateways.DeleteCall.CallCount).To(Equal(0))
				Expect(client.DeleteVpcCall.CallCount).To(Equal(0))
			})
		})

		Context("when there is no tag name", func() {
			BeforeEach(func() {
				client.DescribeVpcsCall.Returns.Output = &awsec2.DescribeVpcsOutput{
					Vpcs: []*awsec2.Vpc{{
						IsDefault: aws.Bool(false),
						VpcId:     aws.String("the-vpc-id"),
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
				client.DescribeVpcsCall.Returns.Error = errors.New("some error")
			})

			It("does not try deleting them", func() {
				err := vpcs.Delete()
				Expect(err).To(MatchError("Describing vpcs: some error"))

				Expect(client.DeleteVpcCall.CallCount).To(Equal(0))
			})
		})

		Context("when subnets fail to delete", func() {
			BeforeEach(func() {
				subnets.DeleteCall.Returns.Error = errors.New("some error")
			})

			It("returns the error", func() {
				err := vpcs.Delete()
				Expect(err).To(MatchError("Deleting subnets for the-vpc-id: some error"))

				Expect(client.DeleteVpcCall.CallCount).To(Equal(0))
			})
		})

		Context("when gateways fail to delete", func() {
			BeforeEach(func() {
				gateways.DeleteCall.Returns.Error = errors.New("some error")
			})

			It("returns the error", func() {
				err := vpcs.Delete()
				Expect(err).To(MatchError("Deleting internet gateways for the-vpc-id: some error"))

				Expect(client.DeleteVpcCall.CallCount).To(Equal(0))
			})
		})

		Context("when the client fails to delete the vpc", func() {
			BeforeEach(func() {
				client.DeleteVpcCall.Returns.Error = errors.New("some error")
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
				Expect(client.DeleteVpcCall.CallCount).To(Equal(0))
			})
		})
	})
})
