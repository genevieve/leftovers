package rds

import awsrds "github.com/aws/aws-sdk-go/service/rds"

type dbInstancesClient interface {
	DeleteDBInstance(*awsrds.DeleteDBInstanceInput) (*awsrds.DeleteDBInstanceOutput, error)
}
