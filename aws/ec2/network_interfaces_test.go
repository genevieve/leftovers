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

var _ = Describe("NetworkInterfaces", func() {
	var (
		client *fakes.NetworkInterfaceClient
		logger *fakes.Logger

		networkInterfaces ec2.NetworkInterfaces
	)

	BeforeEach(func() {
		client = &fakes.NetworkInterfaceClient{}
		logger = &fakes.Logger{}

		networkInterfaces = ec2.NewNetworkInterfaces(client, logger)
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			logger.PromptCall.Returns.Proceed = true
			client.DescribeNetworkInterfacesCall.Returns.Output = &awsec2.DescribeNetworkInterfacesOutput{
				NetworkInterfaces: []*awsec2.NetworkInterface{{
					NetworkInterfaceId: aws.String("banana"),
				}},
			}
		})

		It("deletes network interfaces", func() {
			err := networkInterfaces.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(client.DescribeNetworkInterfacesCall.CallCount).To(Equal(1))
			Expect(client.DeleteNetworkInterfaceCall.CallCount).To(Equal(1))
			Expect(client.DeleteNetworkInterfaceCall.Receives.Input.NetworkInterfaceId).To(Equal(aws.String("banana")))

			Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete network interface banana?"))
			Expect(logger.PrintfCall.Messages).To(Equal([]string{"SUCCESS deleting network interface banana\n"}))
		})

		Context("when the client fails to list network interfaces", func() {
			BeforeEach(func() {
				client.DescribeNetworkInterfacesCall.Returns.Error = errors.New("some error")
			})

			It("does not try deleting them", func() {
				err := networkInterfaces.Delete()
				Expect(err).To(MatchError("Describing network interfaces: some error"))

				Expect(client.DeleteNetworkInterfaceCall.CallCount).To(Equal(0))
			})
		})

		Context("when the client fails to delete the network interface", func() {
			BeforeEach(func() {
				client.DeleteNetworkInterfaceCall.Returns.Error = errors.New("some error")
			})

			It("returns the error", func() {
				err := networkInterfaces.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PrintfCall.Messages).To(Equal([]string{"ERROR deleting network interface banana: some error\n"}))
			})
		})

		Context("when the user responds no to the prompt", func() {
			BeforeEach(func() {
				logger.PromptCall.Returns.Proceed = false
			})

			It("does not delete the network interface", func() {
				err := networkInterfaces.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete network interface banana?"))
				Expect(client.DeleteNetworkInterfaceCall.CallCount).To(Equal(0))
			})
		})
	})
})
