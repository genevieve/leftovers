package fakes

import (
	"sync"

	awsrds "github.com/aws/aws-sdk-go/service/rds"
)

type DbSubnetGroupsClient struct {
	DeleteDBSubnetGroupCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DeleteDBSubnetGroupInput *awsrds.DeleteDBSubnetGroupInput
		}
		Returns struct {
			DeleteDBSubnetGroupOutput *awsrds.DeleteDBSubnetGroupOutput
			Error                     error
		}
		Stub func(*awsrds.DeleteDBSubnetGroupInput) (*awsrds.DeleteDBSubnetGroupOutput, error)
	}
	DescribeDBSubnetGroupsCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DescribeDBSubnetGroupsInput *awsrds.DescribeDBSubnetGroupsInput
		}
		Returns struct {
			DescribeDBSubnetGroupsOutput *awsrds.DescribeDBSubnetGroupsOutput
			Error                        error
		}
		Stub func(*awsrds.DescribeDBSubnetGroupsInput) (*awsrds.DescribeDBSubnetGroupsOutput, error)
	}
}

func (f *DbSubnetGroupsClient) DeleteDBSubnetGroup(param1 *awsrds.DeleteDBSubnetGroupInput) (*awsrds.DeleteDBSubnetGroupOutput, error) {
	f.DeleteDBSubnetGroupCall.Lock()
	defer f.DeleteDBSubnetGroupCall.Unlock()
	f.DeleteDBSubnetGroupCall.CallCount++
	f.DeleteDBSubnetGroupCall.Receives.DeleteDBSubnetGroupInput = param1
	if f.DeleteDBSubnetGroupCall.Stub != nil {
		return f.DeleteDBSubnetGroupCall.Stub(param1)
	}
	return f.DeleteDBSubnetGroupCall.Returns.DeleteDBSubnetGroupOutput, f.DeleteDBSubnetGroupCall.Returns.Error
}
func (f *DbSubnetGroupsClient) DescribeDBSubnetGroups(param1 *awsrds.DescribeDBSubnetGroupsInput) (*awsrds.DescribeDBSubnetGroupsOutput, error) {
	f.DescribeDBSubnetGroupsCall.Lock()
	defer f.DescribeDBSubnetGroupsCall.Unlock()
	f.DescribeDBSubnetGroupsCall.CallCount++
	f.DescribeDBSubnetGroupsCall.Receives.DescribeDBSubnetGroupsInput = param1
	if f.DescribeDBSubnetGroupsCall.Stub != nil {
		return f.DescribeDBSubnetGroupsCall.Stub(param1)
	}
	return f.DescribeDBSubnetGroupsCall.Returns.DescribeDBSubnetGroupsOutput, f.DescribeDBSubnetGroupsCall.Returns.Error
}
