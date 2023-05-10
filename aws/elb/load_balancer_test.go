package elb_test

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	awselb "github.com/aws/aws-sdk-go/service/elb"
	"github.com/genevieve/leftovers/aws/elb"
	"github.com/genevieve/leftovers/aws/elb/fakes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("LoadBalancer", func() {
	var (
		client *fakes.LoadBalancersClient
		name   *string

		loadBalancer elb.LoadBalancer
	)

	BeforeEach(func() {
		client = &fakes.LoadBalancersClient{}
		name = aws.String("the-name")
	})

	Describe("NewLoadBalancer", func() {
		BeforeEach(func() {
			tags := []*awselb.Tag{{Key: aws.String("the-key"), Value: aws.String("the-value")}}
			client.DescribeTagsCall.Returns.DescribeTagsOutput = &awselb.DescribeTagsOutput{
				TagDescriptions: []*awselb.TagDescription{{
					LoadBalancerName: name,
					Tags:             tags,
				}},
			}
		})

		It("returns the identifier", func() {
			loadBalancer = elb.NewLoadBalancer(client, name)
			Expect(client.DescribeTagsCall.CallCount).To(Equal(1))
			Expect(client.DescribeTagsCall.Receives.DescribeTagsInput.LoadBalancerNames[0]).To(Equal(name))

			Expect(loadBalancer.Name()).To(Equal("the-name (the-key:the-value)"))
		})

		Context("when the describe tags call fails", func() {
			BeforeEach(func() {
				client.DescribeTagsCall.Returns.Error = errors.New("banana")
			})

			It("ignores it", func() {
				loadBalancer = elb.NewLoadBalancer(client, name)
				Expect(loadBalancer.Name()).To(Equal("the-name"))
			})
		})
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			loadBalancer = elb.NewLoadBalancer(client, name)
		})

		It("deletes the load balancer", func() {
			err := loadBalancer.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(client.DeleteLoadBalancerCall.CallCount).To(Equal(1))
			Expect(client.DeleteLoadBalancerCall.Receives.DeleteLoadBalancerInput.LoadBalancerName).To(Equal(name))
		})

		Context("when the client fails", func() {
			BeforeEach(func() {
				client.DeleteLoadBalancerCall.Returns.Error = errors.New("banana")
			})

			It("returns the error", func() {
				err := loadBalancer.Delete()
				Expect(err).To(MatchError("Delete: banana"))
			})
		})
	})

	Describe("Name", func() {
		It("returns the identifier", func() {
			loadBalancer = elb.NewLoadBalancer(client, name)
			Expect(loadBalancer.Name()).To(Equal("the-name"))
		})
	})

	Describe("Type", func() {
		It("returns the type", func() {
			loadBalancer = elb.NewLoadBalancer(client, name)
			Expect(loadBalancer.Type()).To(Equal("ELB Load Balancer"))
		})
	})
})
