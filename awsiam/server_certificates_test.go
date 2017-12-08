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
		iamClient          *fakes.IAMClient
		serverCertificates awsiam.ServerCertificates
	)

	BeforeEach(func() {
		iamClient = &fakes.IAMClient{}
		serverCertificates = awsiam.NewServerCertificates(iamClient)
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			iamClient.ListServerCertificatesCall.Returns.Output = &iam.ListServerCertificatesOutput{
				ServerCertificateMetadataList: []*iam.ServerCertificateMetadata{{
					ServerCertificateName: aws.String("banana"),
				}},
			}
		})

		It("deletes iam server certificates", func() {
			serverCertificates.Delete()

			Expect(iamClient.DeleteServerCertificateCall.CallCount).To(Equal(1))
			Expect(iamClient.DeleteServerCertificateCall.Receives.Input.ServerCertificateName).To(Equal(aws.String("banana")))
		})

		Context("when the client fails to list server certificates", func() {
			BeforeEach(func() {
				iamClient.ListServerCertificatesCall.Returns.Error = errors.New("some error")
				iamClient.ListServerCertificatesCall.Returns.Output = &iam.ListServerCertificatesOutput{}
			})

			It("does not try deleting them", func() {
				serverCertificates.Delete()

				Expect(iamClient.DeleteServerCertificateCall.CallCount).To(Equal(0))
			})
		})
	})
})
