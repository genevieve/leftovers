package rds_test

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	awsrds "github.com/aws/aws-sdk-go/service/rds"
	"github.com/genevieve/leftovers/aws/rds"
	"github.com/genevieve/leftovers/aws/rds/fakes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("DBInstances", func() {
	var (
		client *fakes.DbInstancesClient
		logger *fakes.Logger

		dbInstances rds.DBInstances
	)

	BeforeEach(func() {
		client = &fakes.DbInstancesClient{}
		logger = &fakes.Logger{}

		dbInstances = rds.NewDBInstances(client, logger)
	})

	Describe("List", func() {
		var filter string

		BeforeEach(func() {
			logger.PromptWithDetailsCall.Returns.Proceed = true
			client.DescribeDBInstancesCall.Returns.DescribeDBInstancesOutput = &awsrds.DescribeDBInstancesOutput{
				DBInstances: []*awsrds.DBInstance{{
					DBInstanceIdentifier: aws.String("banana"),
					DBInstanceStatus:     aws.String("status"),
				}},
			}
			filter = "ban"
		})

		It("deletes db instances", func() {
			items, err := dbInstances.List(filter)
			Expect(err).NotTo(HaveOccurred())

			Expect(client.DescribeDBInstancesCall.CallCount).To(Equal(1))
			Expect(logger.PromptWithDetailsCall.Receives.ResourceType).To(Equal("RDS DB Instance"))
			Expect(logger.PromptWithDetailsCall.Receives.ResourceName).To(Equal("banana"))

			Expect(items).To(HaveLen(1))
		})

		Context("when the client fails to list db instances", func() {
			BeforeEach(func() {
				client.DescribeDBInstancesCall.Returns.Error = errors.New("some error")
			})

			It("returns the error", func() {
				_, err := dbInstances.List(filter)
				Expect(err).To(MatchError("Describing RDS DB Instances: some error"))
			})
		})

		Context("when the db instance name does not contain the filter", func() {
			It("does not return it in the list", func() {
				items, err := dbInstances.List("kiwi")
				Expect(err).NotTo(HaveOccurred())

				Expect(client.DescribeDBInstancesCall.CallCount).To(Equal(1))
				Expect(logger.PromptWithDetailsCall.CallCount).To(Equal(0))
				Expect(items).To(HaveLen(0))
			})
		})

		Context("when the db instance is being deleted", func() {
			BeforeEach(func() {
				client.DescribeDBInstancesCall.Returns.DescribeDBInstancesOutput = &awsrds.DescribeDBInstancesOutput{
					DBInstances: []*awsrds.DBInstance{{
						DBInstanceIdentifier: aws.String("banana"),
						DBInstanceStatus:     aws.String("deleting"),
					}},
				}
			})

			It("does not return it in the list", func() {
				items, err := dbInstances.List(filter)
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PromptWithDetailsCall.CallCount).To(Equal(0))
				Expect(items).To(HaveLen(0))
			})
		})

		Context("when the user responds no to the prompt", func() {
			BeforeEach(func() {
				logger.PromptWithDetailsCall.Returns.Proceed = false
			})

			It("does not return it in the list", func() {
				items, err := dbInstances.List(filter)
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PromptWithDetailsCall.CallCount).To(Equal(1))
				Expect(items).To(HaveLen(0))
			})
		})
	})
})
