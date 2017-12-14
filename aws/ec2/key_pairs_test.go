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

var _ = Describe("KeyPairs", func() {
	var (
		client *fakes.KeyPairClient
		logger *fakes.Logger

		keys ec2.KeyPairs
	)

	BeforeEach(func() {
		client = &fakes.KeyPairClient{}
		logger = &fakes.Logger{}

		keys = ec2.NewKeyPairs(client, logger)
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			logger.PromptCall.Returns.Proceed = true
			client.DescribeKeyPairsCall.Returns.Output = &awsec2.DescribeKeyPairsOutput{
				KeyPairs: []*awsec2.KeyPairInfo{{
					KeyName: aws.String("banana"),
				}},
			}
		})

		It("deletes ec2 key pairs", func() {
			err := keys.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(client.DescribeKeyPairsCall.CallCount).To(Equal(1))
			Expect(client.DeleteKeyPairCall.CallCount).To(Equal(1))
			Expect(client.DeleteKeyPairCall.Receives.Input.KeyName).To(Equal(aws.String("banana")))
			Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete key pair banana?"))
			Expect(logger.PrintfCall.Messages).To(Equal([]string{"SUCCESS deleting key pair banana\n"}))
		})

		Context("when the client fails to list key pairs", func() {
			BeforeEach(func() {
				client.DescribeKeyPairsCall.Returns.Error = errors.New("some error")
			})

			It("does not try deleting them", func() {
				err := keys.Delete()
				Expect(err).To(MatchError("Describing key pairs: some error"))

				Expect(client.DeleteKeyPairCall.CallCount).To(Equal(0))
			})
		})

		Context("when the client fails to delete the key pair", func() {
			BeforeEach(func() {
				client.DeleteKeyPairCall.Returns.Error = errors.New("some error")
			})

			It("returns the error", func() {
				err := keys.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PrintfCall.Messages).To(Equal([]string{"ERROR deleting key pair banana: some error\n"}))
			})
		})

		Context("when the user responds no to the prompt", func() {
			BeforeEach(func() {
				logger.PromptCall.Returns.Proceed = false
			})

			It("does not delete the key pair", func() {
				err := keys.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete key pair banana?"))
				Expect(client.DeleteKeyPairCall.CallCount).To(Equal(0))
			})
		})
	})
})
