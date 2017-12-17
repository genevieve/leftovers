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

var _ = Describe("Addresses", func() {
	var (
		client *fakes.AddressesClient
		logger *fakes.Logger

		keys ec2.Addresses
	)

	BeforeEach(func() {
		client = &fakes.AddressesClient{}
		logger = &fakes.Logger{}

		keys = ec2.NewAddresses(client, logger)
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			logger.PromptCall.Returns.Proceed = true
			client.DescribeAddressesCall.Returns.Output = &awsec2.DescribeAddressesOutput{
				Addresses: []*awsec2.Address{{
					PublicIp:     aws.String("banana"),
					AllocationId: aws.String("the-allocation-id"),
					InstanceId:   aws.String(""),
				}},
			}
		})

		It("releases ec2 addresses", func() {
			err := keys.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(client.DescribeAddressesCall.CallCount).To(Equal(1))
			Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to release address banana?"))

			Expect(client.ReleaseAddressCall.CallCount).To(Equal(1))
			Expect(client.ReleaseAddressCall.Receives.Input.AllocationId).To(Equal(aws.String("the-allocation-id")))
			Expect(logger.PrintfCall.Messages).To(Equal([]string{"SUCCESS releasing address banana\n"}))
		})

		Context("when the address is in use by an instance", func() {
			BeforeEach(func() {
				logger.PromptCall.Returns.Proceed = true
				client.DescribeAddressesCall.Returns.Output = &awsec2.DescribeAddressesOutput{
					Addresses: []*awsec2.Address{{
						InstanceId: aws.String("the-instance-using-it"),
					}},
				}
			})

			It("does not try to release it", func() {
				err := keys.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(client.DescribeAddressesCall.CallCount).To(Equal(1))
				Expect(logger.PromptCall.CallCount).To(Equal(0))
				Expect(client.ReleaseAddressCall.CallCount).To(Equal(0))
			})
		})

		Context("when the client fails to describe addresses", func() {
			BeforeEach(func() {
				client.DescribeAddressesCall.Returns.Error = errors.New("some error")
			})

			It("does not try releasing them", func() {
				err := keys.Delete()
				Expect(err).To(MatchError("Describing addresses: some error"))

				Expect(client.ReleaseAddressCall.CallCount).To(Equal(0))
			})
		})

		Context("when the client fails to release the address", func() {
			BeforeEach(func() {
				client.ReleaseAddressCall.Returns.Error = errors.New("some error")
			})

			It("returns the error", func() {
				err := keys.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PrintfCall.Messages).To(Equal([]string{"ERROR releasing address banana: some error\n"}))
			})
		})

		Context("when the user responds no to the prompt", func() {
			BeforeEach(func() {
				logger.PromptCall.Returns.Proceed = false
			})

			It("does not release the address", func() {
				err := keys.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to release address banana?"))
				Expect(client.ReleaseAddressCall.CallCount).To(Equal(0))
			})
		})
	})
})
