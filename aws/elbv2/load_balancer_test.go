package elbv2_test

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	awselbv2 "github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/genevieve/leftovers/aws/elbv2"
	"github.com/genevieve/leftovers/aws/elbv2/fakes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("LoadBalancer", func() {
	var (
		loadBalancer elbv2.LoadBalancer
		client       *fakes.LoadBalancersClient
		name         *string
		arn          *string
	)

	BeforeEach(func() {
		client = &fakes.LoadBalancersClient{}
		name = aws.String("the-name")
		arn = aws.String("the-arn")
	})

	Describe("NewLoadBalancer", func() {
		BeforeEach(func() {
			tags := []*awselbv2.Tag{{Key: aws.String("the-key"), Value: aws.String("the-value")}}
			client.DescribeTagsCall.Returns.DescribeTagsOutput = &awselbv2.DescribeTagsOutput{
				TagDescriptions: []*awselbv2.TagDescription{{
					ResourceArn: arn,
					Tags:        tags,
				}},
			}
		})

		It("returns the identifier", func() {
			loadBalancer = elbv2.NewLoadBalancer(client, name, arn)
			Expect(client.DescribeTagsCall.CallCount).To(Equal(1))
			Expect(client.DescribeTagsCall.Receives.DescribeTagsInput.ResourceArns[0]).To(Equal(arn))

			Expect(loadBalancer.Name()).To(Equal("the-name (the-key:the-value)"))
		})

		Context("when the describe tags call fails", func() {
			BeforeEach(func() {
				client.DescribeTagsCall.Returns.Error = errors.New("banana")
			})

			It("ignores it", func() {
				loadBalancer = elbv2.NewLoadBalancer(client, name, arn)
				Expect(loadBalancer.Name()).To(Equal("the-name"))
			})
		})
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			loadBalancer = elbv2.NewLoadBalancer(client, name, arn)
		})
		It("deletes the load balancer", func() {
			err := loadBalancer.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(client.DeleteLoadBalancerCall.CallCount).To(Equal(1))
			Expect(client.DeleteLoadBalancerCall.Receives.DeleteLoadBalancerInput.LoadBalancerArn).To(Equal(arn))
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
			loadBalancer = elbv2.NewLoadBalancer(client, name, arn)
			Expect(loadBalancer.Name()).To(Equal("the-name"))
		})
	})

	Describe("Type", func() {
		It("returns load balancer", func() {
			loadBalancer = elbv2.NewLoadBalancer(client, name, arn)
			Expect(loadBalancer.Type()).To(Equal("ELBV2 Load Balancer"))
		})
	})
})
