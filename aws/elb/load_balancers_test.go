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
		client *fakes.LoadBalancersClient
		logger *fakes.Logger

		loadBalancers elb.LoadBalancers
	)

	BeforeEach(func() {
		client = &fakes.LoadBalancersClient{}
		logger = &fakes.Logger{}

		loadBalancers = elb.NewLoadBalancers(client, logger)
	})

	Describe("Delete", func() {
		var filter string

		BeforeEach(func() {
			logger.PromptCall.Returns.Proceed = true
			client.DescribeLoadBalancersCall.Returns.Output = &awselb.DescribeLoadBalancersOutput{
				LoadBalancerDescriptions: []*awselb.LoadBalancerDescription{{
					LoadBalancerName: aws.String("banana"),
				}},
			}
			filter = "ban"
		})

		It("deletes elb load balancers", func() {
			err := loadBalancers.Delete(filter)
			Expect(err).NotTo(HaveOccurred())

			Expect(client.DescribeLoadBalancersCall.CallCount).To(Equal(1))
			Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete load balancer banana?"))

			Expect(client.DeleteLoadBalancerCall.CallCount).To(Equal(1))
			Expect(client.DeleteLoadBalancerCall.Receives.Input.LoadBalancerName).To(Equal(aws.String("banana")))
			Expect(logger.PrintfCall.Messages).To(Equal([]string{"SUCCESS deleting load balancer banana\n"}))
		})

		Context("when the client fails to list load balancers", func() {
			BeforeEach(func() {
				client.DescribeLoadBalancersCall.Returns.Error = errors.New("some error")
			})

			It("does not try deleting them", func() {
				err := loadBalancers.Delete(filter)
				Expect(err).To(MatchError("Describing load balancers: some error"))

				Expect(client.DeleteLoadBalancerCall.CallCount).To(Equal(0))
			})
		})

		Context("when the load balancer name does not contian the filter", func() {
			It("does not try to delete it", func() {
				err := loadBalancers.Delete("kiwi")
				Expect(err).NotTo(HaveOccurred())

				Expect(client.DescribeLoadBalancersCall.CallCount).To(Equal(1))
				Expect(logger.PromptCall.CallCount).To(Equal(0))
				Expect(client.DeleteLoadBalancerCall.CallCount).To(Equal(0))
			})
		})

		Context("when the client fails to delete the load balancer", func() {
			BeforeEach(func() {
				client.DeleteLoadBalancerCall.Returns.Error = errors.New("some error")
			})

			It("logs the error", func() {
				err := loadBalancers.Delete(filter)
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PrintfCall.Messages).To(Equal([]string{"ERROR deleting load balancer banana: some error\n"}))
			})
		})

		Context("when the user responds no to the prompt", func() {
			BeforeEach(func() {
				logger.PromptCall.Returns.Proceed = false
			})

			It("does not delete the load balancer", func() {
				err := loadBalancers.Delete(filter)
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete load balancer banana?"))
				Expect(client.DeleteLoadBalancerCall.CallCount).To(Equal(0))
			})
		})
	})
})
