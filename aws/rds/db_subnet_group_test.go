package rds_test

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/genevieve/leftovers/aws/rds"
	"github.com/genevieve/leftovers/aws/rds/fakes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("DBSubnetGroup", func() {
	var (
		dbSubnetGroup rds.DBSubnetGroup
		client        *fakes.DbSubnetGroupsClient
		name          *string
	)

	BeforeEach(func() {
		client = &fakes.DbSubnetGroupsClient{}
		name = aws.String("the-name")

		dbSubnetGroup = rds.NewDBSubnetGroup(client, name)
	})

	Describe("Delete", func() {
		It("deletes the db instance", func() {
			err := dbSubnetGroup.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(client.DeleteDBSubnetGroupCall.CallCount).To(Equal(1))
			Expect(client.DeleteDBSubnetGroupCall.Receives.DeleteDBSubnetGroupInput.DBSubnetGroupName).To(Equal(name))
		})

		Context("when the client fails", func() {
			BeforeEach(func() {
				client.DeleteDBSubnetGroupCall.Returns.Error = errors.New("banana")
			})

			It("returns the error", func() {
				err := dbSubnetGroup.Delete()
				Expect(err).To(MatchError("Delete: banana"))
			})
		})
	})

	Describe("Name", func() {
		It("returns the identifier", func() {
			Expect(dbSubnetGroup.Name()).To(Equal("the-name"))
		})
	})

	Describe("Type", func() {
		It("returns the type", func() {
			Expect(dbSubnetGroup.Type()).To(Equal("RDS DB Subnet Group"))
		})
	})
})
