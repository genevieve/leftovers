package ec2_test

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
	"github.com/genevieve/leftovers/aws/ec2"
	"github.com/genevieve/leftovers/aws/ec2/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Vpcs", func() {
	var (
		client *fakes.VpcsClient
		logger *fakes.Logger

		vpcs ec2.Vpcs
	)

	BeforeEach(func() {
		client = &fakes.VpcsClient{}
		logger = &fakes.Logger{}
		routes := &fakes.RouteTables{}
		subnets := &fakes.Subnets{}
		gateways := &fakes.InternetGateways{}
		resourceTags := &fakes.ResourceTags{}

		vpcs = ec2.NewVpcs(client, logger, routes, subnets, gateways, resourceTags)
	})

	Describe("List", func() {
		var filter string

		BeforeEach(func() {
			logger.PromptWithDetailsCall.Returns.Proceed = true
			client.DescribeVpcsCall.Returns.DescribeVpcsOutput = &awsec2.DescribeVpcsOutput{
				Vpcs: []*awsec2.Vpc{{
					IsDefault: aws.Bool(false),
					Tags: []*awsec2.Tag{{
						Key:   aws.String("Name"),
						Value: aws.String("banana"),
					}},
					VpcId: aws.String("the-vpc-id"),
				}},
			}
			filter = "ban"
		})

		It("returns a list of vpcs to delete", func() {
			items, err := vpcs.List(filter, false)
			Expect(err).NotTo(HaveOccurred())

			Expect(client.DescribeVpcsCall.CallCount).To(Equal(1))
			Expect(client.DescribeVpcsCall.Receives.DescribeVpcsInput.Filters[0].Name).To(Equal(aws.String("isDefault")))
			Expect(client.DescribeVpcsCall.Receives.DescribeVpcsInput.Filters[0].Values[0]).To(Equal(aws.String("false")))

			Expect(logger.PromptWithDetailsCall.CallCount).To(Equal(1))
			Expect(logger.PromptWithDetailsCall.Receives.ResourceType).To(Equal("EC2 VPC"))
			Expect(logger.PromptWithDetailsCall.Receives.ResourceName).To(Equal("the-vpc-id (Name:banana)"))

			Expect(items).To(HaveLen(1))
		})

		Context("when the vpc tags contain the filter", func() {
			It("does not return it in the list", func() {
				items, err := vpcs.List("kiwi", false)
				Expect(err).NotTo(HaveOccurred())

				Expect(items).To(HaveLen(0))
			})
		})

		Context("when there is no tag name", func() {
			BeforeEach(func() {
				client.DescribeVpcsCall.Returns.DescribeVpcsOutput = &awsec2.DescribeVpcsOutput{
					Vpcs: []*awsec2.Vpc{{
						IsDefault: aws.Bool(false),
						VpcId:     aws.String("the-vpc-id"),
					}},
				}
			})

			It("uses just the vpc id in the prompt", func() {
				items, err := vpcs.List("the-vpc", false)
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PromptWithDetailsCall.Receives.ResourceName).To(Equal("the-vpc-id"))
				Expect(items).To(HaveLen(1))
			})
		})

		Context("when the client fails to list vpcs", func() {
			BeforeEach(func() {
				client.DescribeVpcsCall.Returns.Error = errors.New("some error")
			})

			It("returns the error", func() {
				_, err := vpcs.List(filter, false)
				Expect(err).To(MatchError("Describe EC2 VPCs: some error"))
			})
		})

		Context("when the user responds no to the prompt", func() {
			BeforeEach(func() {
				logger.PromptWithDetailsCall.Returns.Proceed = false
			})

			It("does not return it in the list", func() {
				items, err := vpcs.List(filter, false)
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PromptWithDetailsCall.CallCount).To(Equal(1))
				Expect(items).To(HaveLen(0))
			})
		})
	})
})
