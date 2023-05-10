package ec2_test

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
	"github.com/genevieve/leftovers/aws/ec2"
	"github.com/genevieve/leftovers/aws/ec2/fakes"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("InternetGateways", func() {
	var (
		client   *fakes.InternetGatewaysClient
		logger   *fakes.Logger
		messages []string

		gateways ec2.InternetGateways
	)

	BeforeEach(func() {
		client = &fakes.InternetGatewaysClient{}

		messages = []string{}
		logger = &fakes.Logger{}
		logger.PrintfCall.Stub = func(format string, v ...interface{}) {
			messages = append(messages, fmt.Sprintf(format, v...))
		}

		gateways = ec2.NewInternetGateways(client, logger)
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			client.DescribeInternetGatewaysCall.Returns.DescribeInternetGatewaysOutput = &awsec2.DescribeInternetGatewaysOutput{
				InternetGateways: []*awsec2.InternetGateway{{
					InternetGatewayId: aws.String("the-gateway-id"),
					Attachments: []*awsec2.InternetGatewayAttachment{{
						VpcId: aws.String("the-vpc-id"),
					}},
				}},
			}
		})

		It("detaches and deletes the internet gateways", func() {
			err := gateways.Delete("the-vpc-id")
			Expect(err).NotTo(HaveOccurred())

			Expect(client.DescribeInternetGatewaysCall.CallCount).To(Equal(1))
			Expect(client.DescribeInternetGatewaysCall.Receives.DescribeInternetGatewaysInput.Filters[0].Name).To(Equal(aws.String("attachment.vpc-id")))
			Expect(client.DescribeInternetGatewaysCall.Receives.DescribeInternetGatewaysInput.Filters[0].Values[0]).To(Equal(aws.String("the-vpc-id")))

			Expect(client.DetachInternetGatewayCall.CallCount).To(Equal(1))
			Expect(client.DetachInternetGatewayCall.Receives.DetachInternetGatewayInput.InternetGatewayId).To(Equal(aws.String("the-gateway-id")))
			Expect(client.DetachInternetGatewayCall.Receives.DetachInternetGatewayInput.VpcId).To(Equal(aws.String("the-vpc-id")))

			Expect(client.DeleteInternetGatewayCall.CallCount).To(Equal(1))
			Expect(client.DeleteInternetGatewayCall.Receives.DeleteInternetGatewayInput.InternetGatewayId).To(Equal(aws.String("the-gateway-id")))

			Expect(messages).To(Equal([]string{
				"[EC2 VPC: the-vpc-id] Detached internet gateway the-gateway-id \n",
				"[EC2 VPC: the-vpc-id] Deleted internet gateway the-gateway-id \n",
			}))
		})

		Context("when the client fails to describe attached internet gateways", func() {
			BeforeEach(func() {
				client.DescribeInternetGatewaysCall.Returns.Error = errors.New("some error")
			})

			It("returns the error and does not try deleting them", func() {
				err := gateways.Delete("banana")
				Expect(err).To(MatchError("Describe EC2 Internet Gateways: some error"))

				Expect(client.DetachInternetGatewayCall.CallCount).To(Equal(0))
				Expect(client.DeleteInternetGatewayCall.CallCount).To(Equal(0))
			})
		})

		Context("when the client fails to detach the internet gateway", func() {
			BeforeEach(func() {
				client.DetachInternetGatewayCall.Returns.Error = errors.New("some error")
			})

			It("logs the error and deletes the internet gateway", func() {
				err := gateways.Delete("banana")
				Expect(err).NotTo(HaveOccurred())

				Expect(client.DeleteInternetGatewayCall.CallCount).To(Equal(1))
				Expect(messages).To(Equal([]string{
					"[EC2 VPC: banana] Detach internet gateway the-gateway-id: some error \n",
					"[EC2 VPC: banana] Deleted internet gateway the-gateway-id \n",
				}))
			})
		})

		Context("when the client fails to delete the internet gateway", func() {
			BeforeEach(func() {
				client.DeleteInternetGatewayCall.Returns.Error = errors.New("some error")
			})

			It("returns the error", func() {
				err := gateways.Delete("banana")
				Expect(err).To(MatchError("Delete the-gateway-id: some error"))

				Expect(messages).To(Equal([]string{
					"[EC2 VPC: banana] Detached internet gateway the-gateway-id \n",
				}))
			})
		})
	})
})
