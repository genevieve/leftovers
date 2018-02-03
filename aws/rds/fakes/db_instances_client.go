package fakes

import awsrds "github.com/aws/aws-sdk-go/service/rds"

type DBInstancesClient struct {
	DeleteDBInstanceCall struct {
		CallCount int
		Receives  struct {
			Input *awsrds.DeleteDBInstanceInput
		}
		Returns struct {
			Output *awsrds.DeleteDBInstanceOutput
			Error  error
		}
	}
}

func (d *DBInstancesClient) DeleteDBInstance(input *awsrds.DeleteDBInstanceInput) (*awsrds.DeleteDBInstanceOutput, error) {
	d.DeleteDBInstanceCall.CallCount++
	d.DeleteDBInstanceCall.Receives.Input = input

	return d.DeleteDBInstanceCall.Returns.Output, d.DeleteDBInstanceCall.Returns.Error
}
