package iam_test

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/genevieve/leftovers/aws/iam"
	"github.com/genevieve/leftovers/aws/iam/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PolicyVersion", func() {
	var (
		policyVersion iam.PolicyVersion
		client        *fakes.PoliciesClient
		name          *string
		arn           *string
		version       *string
	)

	BeforeEach(func() {
		client = &fakes.PoliciesClient{}
		name = aws.String("the-name")
		arn = aws.String("the-arn")
		version = aws.String("the-version")

		policyVersion = iam.NewPolicyVersion(client, name, arn, version)
	})

	Describe("Delete", func() {
		It("deletes the policy version", func() {
			err := policyVersion.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(client.DeletePolicyVersionCall.CallCount).To(Equal(1))
			Expect(client.DeletePolicyVersionCall.Receives.Input.PolicyArn).To(Equal(arn))
			Expect(client.DeletePolicyVersionCall.Receives.Input.VersionId).To(Equal(version))
		})

		Context("when the client fails", func() {
			BeforeEach(func() {
				client.DeletePolicyVersionCall.Returns.Error = errors.New("banana")
			})

			It("returns the error", func() {
				err := policyVersion.Delete()
				Expect(err).To(MatchError("FAILED deleting policy version the-name-the-version: banana"))
			})
		})
	})

	Describe("Name", func() {
		It("returns the identifier", func() {
			Expect(policyVersion.Name()).To(Equal("the-name-the-version"))
		})
	})
})
