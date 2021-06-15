package rds

import (
	"fmt"
	awsrds "github.com/aws/aws-sdk-go/service/rds"

	"github.com/genevieve/leftovers/common"
)

//go:generate faux --interface dbInstancesClient --output fakes/db_instances_client.go
type dbInstancesClient interface {
	DescribeDBInstances(*awsrds.DescribeDBInstancesInput) (*awsrds.DescribeDBInstancesOutput, error)
	DeleteDBInstance(*awsrds.DeleteDBInstanceInput) (*awsrds.DeleteDBInstanceOutput, error)
	WaitUntilDBInstanceDeleted(input *awsrds.DescribeDBInstancesInput) error
}

type DBInstances struct {
	client dbInstancesClient
	logger logger
}

func NewDBInstances(client dbInstancesClient, logger logger) DBInstances {
	return DBInstances{
		client: client,
		logger: logger,
	}
}

func (d DBInstances) List(filter string, regex bool) ([]common.Deletable, error) {
	dbInstances, err := d.client.DescribeDBInstances(&awsrds.DescribeDBInstancesInput{})
	if err != nil {
		return nil, fmt.Errorf("Describing RDS DB Instances: %s", err)
	}

	var resources []common.Deletable
	for _, db := range dbInstances.DBInstances {
		if *db.DBInstanceStatus == "deleting" {
			continue
		}

		r := NewDBInstance(d.client, db.DBInstanceIdentifier)

		if !common.ResourceMatches(r.Name(),  filter, regex) {
			continue
		}

		proceed := d.logger.PromptWithDetails(r.Type(), r.Name())
		if !proceed {
			continue
		}

		resources = append(resources, r)
	}

	return resources, nil
}

func (d DBInstances) Type() string {
	return "rds-db-instance"
}
