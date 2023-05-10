package iam_test

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	awsiam "github.com/aws/aws-sdk-go/service/iam"
	"github.com/genevieve/leftovers/aws/iam"
	"github.com/genevieve/leftovers/aws/iam/fakes"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("AccessKeys", func() {
	var (
		client   *fakes.AccessKeysClient
		logger   *fakes.Logger
		messages []string

		accessKeys iam.AccessKeys
	)

	BeforeEach(func() {
		client = &fakes.AccessKeysClient{}

		messages = []string{}
		logger = &fakes.Logger{}
		logger.PrintfCall.Stub = func(format string, v ...interface{}) {
			messages = append(messages, fmt.Sprintf(format, v...))
		}

		accessKeys = iam.NewAccessKeys(client, logger)
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			client.ListAccessKeysCall.Returns.ListAccessKeysOutput = &awsiam.ListAccessKeysOutput{
				AccessKeyMetadata: []*awsiam.AccessKeyMetadata{{
					AccessKeyId: aws.String("banana"),
				}},
			}
		})

		It("detaches and deletes the accessKeys", func() {
			err := accessKeys.Delete("the-user")
			Expect(err).NotTo(HaveOccurred())

			Expect(client.ListAccessKeysCall.CallCount).To(Equal(1))
			Expect(client.ListAccessKeysCall.Receives.ListAccessKeysInput.UserName).To(Equal(aws.String("the-user")))

			Expect(client.DeleteAccessKeyCall.CallCount).To(Equal(1))
			Expect(client.DeleteAccessKeyCall.Receives.DeleteAccessKeyInput.UserName).To(Equal(aws.String("the-user")))
			Expect(client.DeleteAccessKeyCall.Receives.DeleteAccessKeyInput.AccessKeyId).To(Equal(aws.String("banana")))

			Expect(messages).To(Equal([]string{
				"[IAM User: the-user] Deleted access key banana \n",
			}))
		})

		Context("when the client fails to list access keys", func() {
			BeforeEach(func() {
				client.ListAccessKeysCall.Returns.Error = errors.New("some error")
			})

			It("returns the error", func() {
				err := accessKeys.Delete("the-user")
				Expect(err).To(MatchError("List IAM Access Keys: some error"))
			})
		})

		Context("when the client fails to delete the access key", func() {
			BeforeEach(func() {
				client.DeleteAccessKeyCall.Returns.Error = errors.New("some error")
			})

			It("logs the error", func() {
				err := accessKeys.Delete("the-user")
				Expect(err).NotTo(HaveOccurred())

				Expect(messages).To(Equal([]string{
					"[IAM User: the-user] Delete access key banana: some error \n",
				}))
			})
		})
	})
})
