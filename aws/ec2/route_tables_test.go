package ec2_test

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
	"github.com/genevieve/leftovers/aws/ec2"
	"github.com/genevieve/leftovers/aws/ec2/fakes"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("RouteTables", func() {
	var (
		client       *fakes.RouteTablesClient
		logger       *fakes.Logger
		resourceTags *fakes.ResourceTags
		messages     []string

		routeTables ec2.RouteTables
	)

	BeforeEach(func() {
		client = &fakes.RouteTablesClient{}
		resourceTags = &fakes.ResourceTags{}

		messages = []string{}
		logger = &fakes.Logger{}
		logger.PrintfCall.Stub = func(format string, v ...interface{}) {
			messages = append(messages, fmt.Sprintf(format, v...))
		}

		routeTables = ec2.NewRouteTables(client, logger, resourceTags)
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			client.DescribeRouteTablesCall.Returns.DescribeRouteTablesOutput = &awsec2.DescribeRouteTablesOutput{
				RouteTables: []*awsec2.RouteTable{{
					RouteTableId: aws.String("the-route-table-id"),
					VpcId:        aws.String("the-vpc-id"),
				}},
			}
		})

		It("detaches and deletes the route tables", func() {
			err := routeTables.Delete("the-vpc-id")
			Expect(err).NotTo(HaveOccurred())

			Expect(client.DescribeRouteTablesCall.CallCount).To(Equal(1))
			Expect(client.DescribeRouteTablesCall.Receives.DescribeRouteTablesInput.Filters[0].Name).To(Equal(aws.String("vpc-id")))
			Expect(client.DescribeRouteTablesCall.Receives.DescribeRouteTablesInput.Filters[0].Values[0]).To(Equal(aws.String("the-vpc-id")))
			Expect(client.DescribeRouteTablesCall.Receives.DescribeRouteTablesInput.Filters[1].Name).To(Equal(aws.String("association.main")))
			Expect(client.DescribeRouteTablesCall.Receives.DescribeRouteTablesInput.Filters[1].Values[0]).To(Equal(aws.String("false")))

			Expect(client.DeleteRouteTableCall.CallCount).To(Equal(1))
			Expect(client.DeleteRouteTableCall.Receives.DeleteRouteTableInput.RouteTableId).To(Equal(aws.String("the-route-table-id")))

			Expect(resourceTags.DeleteCall.CallCount).To(Equal(1))
			Expect(resourceTags.DeleteCall.Receives.FilterName).To(Equal("route-table"))
			Expect(resourceTags.DeleteCall.Receives.FilterValue).To(Equal("the-route-table-id"))

			Expect(messages).To(Equal([]string{
				"[EC2 VPC: the-vpc-id] Deleted route table the-route-table-id \n",
				"[EC2 VPC: the-vpc-id] Deleted route table the-route-table-id tags \n",
			}))
		})

		Context("when the client fails to describe route tables", func() {
			BeforeEach(func() {
				client.DescribeRouteTablesCall.Returns.Error = errors.New("some error")
			})

			It("returns the error and does not try deleting them", func() {
				err := routeTables.Delete("banana")
				Expect(err).To(MatchError("Describe EC2 Route Tables: some error"))

				Expect(client.DeleteRouteTableCall.CallCount).To(Equal(0))
			})
		})

		Context("when the route table has an association id", func() {
			BeforeEach(func() {
				client.DescribeRouteTablesCall.Returns.DescribeRouteTablesOutput = &awsec2.DescribeRouteTablesOutput{
					RouteTables: []*awsec2.RouteTable{{
						RouteTableId: aws.String("the-route-table-id"),
						VpcId:        aws.String("the-vpc-id"),
						Associations: []*awsec2.RouteTableAssociation{{
							Main:                    aws.Bool(false),
							RouteTableAssociationId: aws.String("the-association-id"),
							RouteTableId:            aws.String("the-route-table-id"),
							SubnetId:                aws.String("the-subnet-id"),
						}},
					}},
				}
			})

			It("disassociates it from the subnet before trying to delete it", func() {
				err := routeTables.Delete("the-vpc-id")
				Expect(err).NotTo(HaveOccurred())

				Expect(client.DescribeRouteTablesCall.CallCount).To(Equal(1))
				Expect(client.DisassociateRouteTableCall.CallCount).To(Equal(1))
				Expect(client.DisassociateRouteTableCall.Receives.DisassociateRouteTableInput.AssociationId).To(Equal(aws.String("the-association-id")))
				Expect(client.DeleteRouteTableCall.CallCount).To(Equal(1))

				Expect(messages).To(Equal([]string{
					"[EC2 VPC: the-vpc-id] Disassociated route table the-route-table-id \n",
					"[EC2 VPC: the-vpc-id] Deleted route table the-route-table-id \n",
					"[EC2 VPC: the-vpc-id] Deleted route table the-route-table-id tags \n",
				}))
			})

			Context("when the client fails to disassociate the route table", func() {
				BeforeEach(func() {
					client.DisassociateRouteTableCall.Returns.Error = errors.New("some error")
				})

				It("logs the error", func() {
					err := routeTables.Delete("the-vpc-id")
					Expect(err).NotTo(HaveOccurred())

					Expect(client.DisassociateRouteTableCall.CallCount).To(Equal(1))
					Expect(messages).To(Equal([]string{
						"[EC2 VPC: the-vpc-id] Disassociate route table the-route-table-id: some error \n",
						"[EC2 VPC: the-vpc-id] Deleted route table the-route-table-id \n",
						"[EC2 VPC: the-vpc-id] Deleted route table the-route-table-id tags \n",
					}))
				})
			})
		})

		Context("when the client fails to delete the route table", func() {
			BeforeEach(func() {
				client.DeleteRouteTableCall.Returns.Error = errors.New("banana")
			})

			It("returns the error", func() {
				err := routeTables.Delete("the-vpc-id")
				Expect(err).To(MatchError("Delete the-route-table-id: banana"))
			})
		})
	})
})
