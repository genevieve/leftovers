package fakes

import (
	"sync"

	awsrds "github.com/aws/aws-sdk-go/service/rds"
)

type DbInstancesClient struct {
	DeleteDBInstanceCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DeleteDBInstanceInput *awsrds.DeleteDBInstanceInput
		}
		Returns struct {
			DeleteDBInstanceOutput *awsrds.DeleteDBInstanceOutput
			Error                  error
		}
		Stub func(*awsrds.DeleteDBInstanceInput) (*awsrds.DeleteDBInstanceOutput, error)
	}
	DescribeDBInstancesCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DescribeDBInstancesInput *awsrds.DescribeDBInstancesInput
		}
		Returns struct {
			DescribeDBInstancesOutput *awsrds.DescribeDBInstancesOutput
			Error                     error
		}
		Stub func(*awsrds.DescribeDBInstancesInput) (*awsrds.DescribeDBInstancesOutput, error)
	}
}

func (f *DbInstancesClient) DeleteDBInstance(param1 *awsrds.DeleteDBInstanceInput) (*awsrds.DeleteDBInstanceOutput, error) {
	f.DeleteDBInstanceCall.Lock()
	defer f.DeleteDBInstanceCall.Unlock()
	f.DeleteDBInstanceCall.CallCount++
	f.DeleteDBInstanceCall.Receives.DeleteDBInstanceInput = param1
	if f.DeleteDBInstanceCall.Stub != nil {
		return f.DeleteDBInstanceCall.Stub(param1)
	}
	return f.DeleteDBInstanceCall.Returns.DeleteDBInstanceOutput, f.DeleteDBInstanceCall.Returns.Error
}
func (f *DbInstancesClient) DescribeDBInstances(param1 *awsrds.DescribeDBInstancesInput) (*awsrds.DescribeDBInstancesOutput, error) {
	f.DescribeDBInstancesCall.Lock()
	defer f.DescribeDBInstancesCall.Unlock()
	f.DescribeDBInstancesCall.CallCount++
	f.DescribeDBInstancesCall.Receives.DescribeDBInstancesInput = param1
	if f.DescribeDBInstancesCall.Stub != nil {
		return f.DescribeDBInstancesCall.Stub(param1)
	}
	return f.DescribeDBInstancesCall.Returns.DescribeDBInstancesOutput, f.DescribeDBInstancesCall.Returns.Error
}
