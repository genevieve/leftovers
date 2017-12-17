package ec2_test

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
	"github.com/genevievelesperance/leftovers/aws/ec2"
	"github.com/genevievelesperance/leftovers/aws/ec2/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RouteTables", func() {
	var (
		client *fakes.RouteTablesClient
		logger *fakes.Logger

		routeTables ec2.RouteTables
	)

	BeforeEach(func() {
		client = &fakes.RouteTablesClient{}
		logger = &fakes.Logger{}

		routeTables = ec2.NewRouteTables(client, logger)
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			client.DescribeRouteTablesCall.Returns.Output = &awsec2.DescribeRouteTablesOutput{
				RouteTables: []*awsec2.RouteTable{{
					RouteTableId: aws.String("the-route-table-id"),
					VpcId:        aws.String("the-vpc-id"),
				}},
			}
		})

		It("detaches and deletes the routeTables", func() {
			err := routeTables.Delete("the-vpc-id")
			Expect(err).NotTo(HaveOccurred())

			Expect(client.DescribeRouteTablesCall.CallCount).To(Equal(1))
			Expect(client.DescribeRouteTablesCall.Receives.Input.Filters[0].Name).To(Equal(aws.String("vpc-id")))
			Expect(client.DescribeRouteTablesCall.Receives.Input.Filters[0].Values[0]).To(Equal(aws.String("the-vpc-id")))

			Expect(client.DeleteRouteTableCall.CallCount).To(Equal(1))
			Expect(client.DeleteRouteTableCall.Receives.Input.RouteTableId).To(Equal(aws.String("the-route-table-id")))

			Expect(logger.PrintfCall.Messages).To(Equal([]string{
				"SUCCESS deleting route table the-route-table-id\n",
			}))
		})

		Context("when the client fails to describe route tables", func() {
			BeforeEach(func() {
				client.DescribeRouteTablesCall.Returns.Error = errors.New("some error")
			})

			It("returns the error and does not try deleting them", func() {
				err := routeTables.Delete("banana")
				Expect(err).To(MatchError("Describing route tables: some error"))

				Expect(client.DeleteRouteTableCall.CallCount).To(Equal(0))
			})
		})

		Context("when the client fails to delete the route-table", func() {
			BeforeEach(func() {
				client.DeleteRouteTableCall.Returns.Error = errors.New("some error")
			})

			It("logs the error", func() {
				err := routeTables.Delete("banana")
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PrintfCall.Messages).To(Equal([]string{
					"ERROR deleting route table the-route-table-id: some error\n",
				}))
			})
		})
	})
})
