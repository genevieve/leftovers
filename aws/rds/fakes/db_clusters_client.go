package fakes

import (
	"sync"

	awsrds "github.com/aws/aws-sdk-go/service/rds"
)

type DbClustersClient struct {
	DeleteDBClusterCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DeleteDBClusterInput *awsrds.DeleteDBClusterInput
		}
		Returns struct {
			DeleteDBClusterOutput *awsrds.DeleteDBClusterOutput
			Error                 error
		}
		Stub func(*awsrds.DeleteDBClusterInput) (*awsrds.DeleteDBClusterOutput, error)
	}
	DescribeDBClustersCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DescribeDBClustersInput *awsrds.DescribeDBClustersInput
		}
		Returns struct {
			DescribeDBClustersOutput *awsrds.DescribeDBClustersOutput
			Error                    error
		}
		Stub func(*awsrds.DescribeDBClustersInput) (*awsrds.DescribeDBClustersOutput, error)
	}
}

func (f *DbClustersClient) DeleteDBCluster(param1 *awsrds.DeleteDBClusterInput) (*awsrds.DeleteDBClusterOutput, error) {
	f.DeleteDBClusterCall.Lock()
	defer f.DeleteDBClusterCall.Unlock()
	f.DeleteDBClusterCall.CallCount++
	f.DeleteDBClusterCall.Receives.DeleteDBClusterInput = param1
	if f.DeleteDBClusterCall.Stub != nil {
		return f.DeleteDBClusterCall.Stub(param1)
	}
	return f.DeleteDBClusterCall.Returns.DeleteDBClusterOutput, f.DeleteDBClusterCall.Returns.Error
}
func (f *DbClustersClient) DescribeDBClusters(param1 *awsrds.DescribeDBClustersInput) (*awsrds.DescribeDBClustersOutput, error) {
	f.DescribeDBClustersCall.Lock()
	defer f.DescribeDBClustersCall.Unlock()
	f.DescribeDBClustersCall.CallCount++
	f.DescribeDBClustersCall.Receives.DescribeDBClustersInput = param1
	if f.DescribeDBClustersCall.Stub != nil {
		return f.DescribeDBClustersCall.Stub(param1)
	}
	return f.DescribeDBClustersCall.Returns.DescribeDBClustersOutput, f.DescribeDBClustersCall.Returns.Error
}
