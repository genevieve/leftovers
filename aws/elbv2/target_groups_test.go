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
				TargetGroupName: aws.String("precursor"),
				TargetGroupArn:  aws.String("precursor-banana"),
			}, {
				TargetGroupName: aws.String("banana"),
				TargetGroupArn:  aws.String("the-arn"),
			}},
		}
		logger = &fakes.Logger{}
		logger.PromptCall.Returns.Proceed = true

		targetGroups = elbv2.NewTargetGroups(client, logger)
	})

	It("Deletes target groups", func() {
		err := targetGroups.Delete()
		Expect(err).NotTo(HaveOccurred())

		Expect(logger.PromptCall.CallCount).To(Equal(2))

		Expect(client.DeleteTargetGroupCall.CallCount).To(Equal(2))
		Expect(*client.DeleteTargetGroupCall.Receives.Input.TargetGroupArn).To(Equal("the-arn"))

		Expect(logger.PrintfCall.Messages).To(Equal([]string{
			"SUCCESS deleting target group precursor\n",
			"SUCCESS deleting target group banana\n",
		}))
	})

	Context("when the user doesn't want to delete", func() {
		BeforeEach(func() {
			logger.PromptCall.Returns.Proceed = false
		})

		It("does not delete the target group", func() {
			err := targetGroups.Delete()
			Expect(err).NotTo(HaveOccurred())
			Expect(client.DeleteTargetGroupCall.CallCount).To(Equal(0))
		})
	})

	Context("when we fail to describe target groups", func() {
		BeforeEach(func() {
			client.DescribeTargetGroupsCall.Returns.Error = errors.New("banana")
		})

		It("returns the error", func() {
			err := targetGroups.Delete()
			Expect(err).To(MatchError("Describing target groups: banana"))
		})
	})

	Context("when we fail to delete target groups", func() {
		BeforeEach(func() {
			client.DeleteTargetGroupCall.Returns.Error = errors.New("anana")
		})

		It("logs the error, but doesn't return", func() {
			err := targetGroups.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(logger.PrintfCall.Messages).To(Equal([]string{
				"ERROR deleting target group precursor: anana\n",
				"ERROR deleting target group banana: anana\n",
			}))
		})
	})
})
