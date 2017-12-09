package elb_test

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	awselb "github.com/aws/aws-sdk-go/service/elb"
	"github.com/genevievelesperance/leftovers/aws/elb"
	"github.com/genevievelesperance/leftovers/aws/elb/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("LoadBalancers", func() {
	var (
		elbClient *fakes.ELBClient
		logger    *fakes.Logger

		loadBalancers elb.LoadBalancers
	)

	BeforeEach(func() {
		elbClient = &fakes.ELBClient{}
		logger = &fakes.Logger{}

		loadBalancers = elb.NewLoadBalancers(elbClient, logger)
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			logger.PromptCall.Returns.Proceed = true
			elbClient.DescribeLoadBalancersCall.Returns.Output = &awselb.DescribeLoadBalancersOutput{
				LoadBalancerDescriptions: []*awselb.LoadBalancerDescription{{
					LoadBalancerName: aws.String("banana"),
				}},
			}
		})

		It("deletes elb load balancers", func() {
			err := loadBalancers.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(elbClient.DescribeLoadBalancersCall.CallCount).To(Equal(1))
			Expect(elbClient.DeleteLoadBalancerCall.CallCount).To(Equal(1))
			Expect(elbClient.DeleteLoadBalancerCall.Receives.Input.LoadBalancerName).To(Equal(aws.String("banana")))
			Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete load balancer banana?"))
			Expect(logger.PrintfCall.Messages).To(Equal([]string{"SUCCESS deleting load balancer banana\n"}))
		})

		Context("when the client fails to list load balancers", func() {
			BeforeEach(func() {
				elbClient.DescribeLoadBalancersCall.Returns.Error = errors.New("some error")
			})

			It("does not try deleting them", func() {
				err := loadBalancers.Delete()
				Expect(err.Error()).To(Equal("Describing load balancers: some error"))

				Expect(elbClient.DeleteLoadBalancerCall.CallCount).To(Equal(0))
			})
		})

		Context("when the client fails to delete the loadBalancer", func() {
			BeforeEach(func() {
				elbClient.DeleteLoadBalancerCall.Returns.Error = errors.New("some error")
			})

			It("returns the error", func() {
				err := loadBalancers.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PrintfCall.Messages).To(Equal([]string{"ERROR deleting load balancer banana: some error\n"}))
			})
		})

		Context("when the user responds no to the prompt", func() {
			BeforeEach(func() {
				logger.PromptCall.Returns.Proceed = false
			})

			It("returns the error", func() {
				err := loadBalancers.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete load balancer banana?"))
				Expect(elbClient.DeleteLoadBalancerCall.CallCount).To(Equal(0))
			})
		})
	})
})
