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
		ec2Client *fakes.EC2Client
		logger    *fakes.Logger

		keyPairs ec2.KeyPairs
	)

	BeforeEach(func() {
		ec2Client = &fakes.EC2Client{}
		logger = &fakes.Logger{}

		keyPairs = ec2.NewKeyPairs(ec2Client, logger)
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			logger.PromptCall.Returns.Proceed = true
			ec2Client.DescribeKeyPairsCall.Returns.Output = &awsec2.DescribeKeyPairsOutput{
				KeyPairs: []*awsec2.KeyPairInfo{{
					KeyName: aws.String("banana"),
				}},
			}
		})

		It("deletes ec2 key pairs", func() {
			err := keyPairs.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(ec2Client.DescribeKeyPairsCall.CallCount).To(Equal(1))
			Expect(ec2Client.DeleteKeyPairCall.CallCount).To(Equal(1))
			Expect(ec2Client.DeleteKeyPairCall.Receives.Input.KeyName).To(Equal(aws.String("banana")))
			Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete key pair banana?"))
			Expect(logger.PrintfCall.Messages).To(Equal([]string{"SUCCESS deleting key pair banana\n"}))
		})

		Context("when the client fails to list key pairs", func() {
			BeforeEach(func() {
				ec2Client.DescribeKeyPairsCall.Returns.Error = errors.New("some error")
			})

			It("does not try deleting them", func() {
				err := keyPairs.Delete()
				Expect(err.Error()).To(Equal("Describing key pairs: some error"))

				Expect(ec2Client.DeleteKeyPairCall.CallCount).To(Equal(0))
			})
		})

		Context("when the client fails to delete the key pair", func() {
			BeforeEach(func() {
				ec2Client.DeleteKeyPairCall.Returns.Error = errors.New("some error")
			})

			It("returns the error", func() {
				err := keyPairs.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PrintfCall.Messages).To(Equal([]string{"ERROR deleting key pair banana: some error\n"}))
			})
		})

		Context("when the user responds no to the prompt", func() {
			BeforeEach(func() {
				logger.PromptCall.Returns.Proceed = false
			})

			It("does not delete the key pair", func() {
				err := keyPairs.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete key pair banana?"))
				Expect(ec2Client.DeleteKeyPairCall.CallCount).To(Equal(0))
			})
		})
	})
})
