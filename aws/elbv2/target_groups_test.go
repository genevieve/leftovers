package elbv2_test

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	awselbv2 "github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/genevievelesperance/leftovers/aws/elbv2"
	"github.com/genevievelesperance/leftovers/aws/elbv2/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TargetGroups", func() {
	var (
		client *fakes.TargetGroupsClient
		logger *fakes.Logger

		targetGroups elbv2.TargetGroups
	)

	BeforeEach(func() {
		client = &fakes.TargetGroupsClient{}
		client.DescribeTargetGroupsCall.Returns.Output = &awselbv2.DescribeTargetGroupsOutput{
			TargetGroups: []*awselbv2.TargetGroup{{
				TargetGroupName: aws.String("precursor-banana"),
				TargetGroupArn:  aws.String("precursor-arn"),
			}, {
				TargetGroupName: aws.String("banana"),
				TargetGroupArn:  aws.String("arn"),
			}},
		}
		logger = &fakes.Logger{}
		logger.PromptCall.Returns.Proceed = true

		targetGroups = elbv2.NewTargetGroups(client, logger)
	})

	Describe("Delete", func() {
		var filter string

		BeforeEach(func() {
			filter = "banana"
		})

		It("deletes target groups", func() {
			err := targetGroups.Delete(filter)
			Expect(err).NotTo(HaveOccurred())

			Expect(logger.PromptCall.CallCount).To(Equal(2))

			Expect(client.DeleteTargetGroupCall.CallCount).To(Equal(2))
			Expect(*client.DeleteTargetGroupCall.Receives.Input.TargetGroupArn).To(Equal("arn"))

			Expect(logger.PrintfCall.Messages).To(Equal([]string{
				"SUCCESS deleting target group precursor-banana\n",
				"SUCCESS deleting target group banana\n",
			}))
		})

		Context("when the client fails to describe target groups", func() {
			BeforeEach(func() {
				client.DescribeTargetGroupsCall.Returns.Error = errors.New("banana")
			})

			It("returns the error", func() {
				err := targetGroups.Delete(filter)
				Expect(err).To(MatchError("Describing target groups: banana"))
			})
		})

		Context("when the target group name does not contain the filter", func() {
			It("does not try to delete it", func() {
				err := targetGroups.Delete("kiwi")
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PromptCall.CallCount).To(Equal(0))
				Expect(client.DeleteTargetGroupCall.CallCount).To(Equal(0))
			})
		})

		Context("when the user doesn't want to delete", func() {
			BeforeEach(func() {
				logger.PromptCall.Returns.Proceed = false
			})

			It("does not delete the target group", func() {
				err := targetGroups.Delete(filter)
				Expect(err).NotTo(HaveOccurred())
				Expect(client.DeleteTargetGroupCall.CallCount).To(Equal(0))
			})
		})

		Context("when we fail to delete target groups", func() {
			BeforeEach(func() {
				client.DeleteTargetGroupCall.Returns.Error = errors.New("anana")
			})

			It("logs the error, but doesn't return", func() {
				err := targetGroups.Delete(filter)
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PrintfCall.Messages).To(Equal([]string{
					"ERROR deleting target group precursor-banana: anana\n",
					"ERROR deleting target group banana: anana\n",
				}))
			})
		})
	})
})
