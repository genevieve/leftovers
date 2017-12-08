package awsiam_test

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/genevievelesperance/leftovers/awsiam"
	"github.com/genevievelesperance/leftovers/awsiam/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ServerCertificates", func() {
	var (
		iamClient *fakes.IAMClient
		logger    *fakes.Logger

		serverCertificates awsiam.ServerCertificates
	)

	BeforeEach(func() {
		iamClient = &fakes.IAMClient{}
		logger = &fakes.Logger{}

		serverCertificates = awsiam.NewServerCertificates(iamClient, logger)
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			logger.PromptCall.Returns.Proceed = true
			iamClient.ListServerCertificatesCall.Returns.Output = &iam.ListServerCertificatesOutput{
				ServerCertificateMetadataList: []*iam.ServerCertificateMetadata{{
					ServerCertificateName: aws.String("banana"),
				}},
			}
		})

		It("deletes iam server certificates", func() {
			err := serverCertificates.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(iamClient.DeleteServerCertificateCall.CallCount).To(Equal(1))
			Expect(iamClient.DeleteServerCertificateCall.Receives.Input.ServerCertificateName).To(Equal(aws.String("banana")))
			Expect(logger.PrintfCall.Messages).To(Equal([]string{"SUCCESS deleting server certificate banana\n"}))
		})

		Context("when the client fails to list server certificates", func() {
			BeforeEach(func() {
				iamClient.ListServerCertificatesCall.Returns.Error = errors.New("some error")
			})

			It("does not try deleting them", func() {
				err := serverCertificates.Delete()
				Expect(err.Error()).To(Equal("Listing server certificates: some error"))

				Expect(iamClient.DeleteServerCertificateCall.CallCount).To(Equal(0))
			})
		})

		Context("when the client fails to delete the server certificate", func() {
			BeforeEach(func() {
				iamClient.DeleteServerCertificateCall.Returns.Error = errors.New("some error")
			})

			It("does not try deleting them", func() {
				err := serverCertificates.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PrintfCall.Messages).To(Equal([]string{"ERROR deleting server certificate banana: some error\n"}))
			})
		})

		Context("when the user responds no to the prompt", func() {
			BeforeEach(func() {
				logger.PromptCall.Returns.Proceed = false
			})

			It("returns the error", func() {
				err := serverCertificates.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete server certificate banana?"))
				Expect(iamClient.DeleteServerCertificateCall.CallCount).To(Equal(0))
			})
		})
	})
})
