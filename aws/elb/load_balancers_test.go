package elb_test

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	awselb "github.com/aws/aws-sdk-go/service/elb"
	"github.com/genevieve/leftovers/aws/elb"
	"github.com/genevieve/leftovers/aws/elb/fakes"
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

	Describe("List", func() {
		var filter string

		BeforeEach(func() {
			logger.PromptWithDetailsCall.Returns.Proceed = true
			client.DescribeLoadBalancersCall.Returns.DescribeLoadBalancersOutput = &awselb.DescribeLoadBalancersOutput{
				LoadBalancerDescriptions: []*awselb.LoadBalancerDescription{{
					LoadBalancerName: aws.String("banana"),
				}},
			}
			filter = "ban"
		})

		It("deletes elb load balancers", func() {
			items, err := loadBalancers.List(filter, false)
			Expect(err).NotTo(HaveOccurred())

			Expect(client.DescribeLoadBalancersCall.CallCount).To(Equal(1))
			Expect(logger.PromptWithDetailsCall.Receives.ResourceType).To(Equal("ELB Load Balancer"))
			Expect(logger.PromptWithDetailsCall.Receives.ResourceName).To(Equal("banana"))

			Expect(items).To(HaveLen(1))
		})

		Context("when the client fails to list load balancers", func() {
			BeforeEach(func() {
				client.DescribeLoadBalancersCall.Returns.Error = errors.New("some error")
			})

			It("returns the error", func() {
				_, err := loadBalancers.List(filter, false)
				Expect(err).To(MatchError("Describe ELB Load Balancers: some error"))
			})
		})

		Context("when the load balancer name does not contain the filter", func() {
			It("does not return it in the list", func() {
				items, err := loadBalancers.List("kiwi", false)
				Expect(err).NotTo(HaveOccurred())

				Expect(client.DescribeLoadBalancersCall.CallCount).To(Equal(1))
				Expect(logger.PromptWithDetailsCall.CallCount).To(Equal(0))
				Expect(items).To(HaveLen(0))
			})
		})

		Context("when the user responds no to the prompt", func() {
			BeforeEach(func() {
				logger.PromptWithDetailsCall.Returns.Proceed = false
			})

			It("does not return it in the list", func() {
				items, err := loadBalancers.List(filter, false)
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PromptWithDetailsCall.CallCount).To(Equal(1))
				Expect(items).To(HaveLen(0))
			})
		})
	})
})
