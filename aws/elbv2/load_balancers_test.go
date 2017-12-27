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

var _ = Describe("LoadBalancers", func() {
	var (
		client *fakes.LoadBalancersClient
		logger *fakes.Logger

		loadBalancers elbv2.LoadBalancers
	)

	BeforeEach(func() {
		client = &fakes.LoadBalancersClient{}
		logger = &fakes.Logger{}

		loadBalancers = elbv2.NewLoadBalancers(client, logger)
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			logger.PromptCall.Returns.Proceed = true
			client.DescribeLoadBalancersCall.Returns.Output = &awselbv2.DescribeLoadBalancersOutput{
				LoadBalancers: []*awselbv2.LoadBalancer{{
					LoadBalancerName: aws.String("banana"),
					LoadBalancerArn:  aws.String("the-arn"),
				}},
			}
		})

		It("deletes elbv2 load balancers", func() {
			err := loadBalancers.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(client.DescribeLoadBalancersCall.CallCount).To(Equal(1))
			Expect(client.DeleteLoadBalancerCall.CallCount).To(Equal(1))
			Expect(client.DeleteLoadBalancerCall.Receives.Input.LoadBalancerArn).To(Equal(aws.String("the-arn")))
			Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete load balancer banana?"))
			Expect(logger.PrintfCall.Messages).To(Equal([]string{"SUCCESS deleting load balancer banana\n"}))
		})

		Context("when the client fails to list load balancers", func() {
			BeforeEach(func() {
				client.DescribeLoadBalancersCall.Returns.Error = errors.New("some error")
			})

			It("does not try deleting them", func() {
				err := loadBalancers.Delete()
				Expect(err).To(MatchError("Describing load balancers: some error"))

				Expect(client.DeleteLoadBalancerCall.CallCount).To(Equal(0))
			})
		})

		Context("when the client fails to delete the load balancer", func() {
			BeforeEach(func() {
				client.DeleteLoadBalancerCall.Returns.Error = errors.New("some error")
			})

			It("logs the error", func() {
				err := loadBalancers.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PrintfCall.Messages).To(Equal([]string{"ERROR deleting load balancer banana: some error\n"}))
			})
		})

		Context("when the user responds no to the prompt", func() {
			BeforeEach(func() {
				logger.PromptCall.Returns.Proceed = false
			})

			It("does not delete the load balancer", func() {
				err := loadBalancers.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete load balancer banana?"))
				Expect(client.DeleteLoadBalancerCall.CallCount).To(Equal(0))
			})
		})
	})
})
