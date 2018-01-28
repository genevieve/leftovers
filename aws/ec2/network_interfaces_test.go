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

	Describe("List", func() {
		var filter string

		BeforeEach(func() {
			logger.PromptCall.Returns.Proceed = true
			client.DescribeNetworkInterfacesCall.Returns.Output = &awsec2.DescribeNetworkInterfacesOutput{
				NetworkInterfaces: []*awsec2.NetworkInterface{{
					NetworkInterfaceId: aws.String("banana"),
				}},
			}
			filter = "ban"
		})

		It("returns a list of network interfaces to delete", func() {
			items, err := networkInterfaces.List(filter)
			Expect(err).NotTo(HaveOccurred())

			Expect(client.DescribeNetworkInterfacesCall.CallCount).To(Equal(1))

			Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete network interface banana?"))

			Expect(items).To(HaveLen(1))
			Expect(items).To(HaveKeyWithValue("banana", "banana"))
		})

		Context("when the client fails to list network interfaces", func() {
			BeforeEach(func() {
				client.DescribeNetworkInterfacesCall.Returns.Error = errors.New("some error")
			})

			It("returns the error", func() {
				_, err := networkInterfaces.List(filter)
				Expect(err).To(MatchError("Describing network interfaces: some error"))
			})
		})

		Context("when the network interface name does not contain the filter", func() {
			It("does not return it in the list", func() {
				items, err := networkInterfaces.List("kiwi")
				Expect(err).NotTo(HaveOccurred())

				Expect(client.DescribeNetworkInterfacesCall.CallCount).To(Equal(1))
				Expect(logger.PromptCall.CallCount).To(Equal(0))
				Expect(items).To(HaveLen(0))
			})
		})

		Context("when the network interface has tags", func() {
			BeforeEach(func() {
				client.DescribeNetworkInterfacesCall.Returns.Output = &awsec2.DescribeNetworkInterfacesOutput{
					NetworkInterfaces: []*awsec2.NetworkInterface{{
						NetworkInterfaceId: aws.String("banana"),
						TagSet: []*awsec2.Tag{{
							Key:   aws.String("the-key"),
							Value: aws.String("the-value"),
						}},
					}},
				}
			})

			It("uses them in the prompt", func() {
				items, err := networkInterfaces.List(filter)
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete network interface banana (the-key:the-value)?"))
				Expect(items).To(HaveKeyWithValue("banana (the-key:the-value)", "banana"))
			})
		})

		Context("when the user responds no to the prompt", func() {
			BeforeEach(func() {
				logger.PromptCall.Returns.Proceed = false
			})

			It("does not return it in the list", func() {
				items, err := networkInterfaces.List(filter)
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete network interface banana?"))
				Expect(items).To(HaveLen(0))
			})
		})
	})

	Describe("Delete", func() {
		var items map[string]string

		BeforeEach(func() {
			items = map[string]string{"banana": "the-id"}
		})

		It("deletes network interfaces", func() {
			err := networkInterfaces.Delete(items)
			Expect(err).NotTo(HaveOccurred())

			Expect(client.DeleteNetworkInterfaceCall.CallCount).To(Equal(1))
			Expect(client.DeleteNetworkInterfaceCall.Receives.Input.NetworkInterfaceId).To(Equal(aws.String("the-id")))

			Expect(logger.PrintfCall.Messages).To(Equal([]string{"SUCCESS deleting network interface the-id\n"}))
		})

		Context("when the client fails to delete the network interface", func() {
			BeforeEach(func() {
				client.DeleteNetworkInterfaceCall.Returns.Error = errors.New("some error")
			})

			It("logs the error", func() {
				err := networkInterfaces.Delete(items)
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PrintfCall.Messages).To(Equal([]string{"ERROR deleting network interface the-id: some error\n"}))
			})
		})
	})
})
