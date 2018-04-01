package fakes

import awsrds "github.com/aws/aws-sdk-go/service/rds"

type DBClustersClient struct {
	DescribeDBClustersCall struct {
		CallCount int
		Returns   struct {
			Output *awsrds.DescribeDBClustersOutput
			Error  error
		}
	}
	DeleteDBClusterCall struct {
		CallCount int
		Receives  struct {
			Input *awsrds.DeleteDBClusterInput
		}
		Returns struct {
			Output *awsrds.DeleteDBClusterOutput
			Error  error
		}
	}
}

func (d *DBClustersClient) DeleteDBCluster(input *awsrds.DeleteDBClusterInput) (*awsrds.DeleteDBClusterOutput, error) {
	d.DeleteDBClusterCall.CallCount++
	d.DeleteDBClusterCall.Receives.Input = input

	return d.DeleteDBClusterCall.Returns.Output, d.DeleteDBClusterCall.Returns.Error
}

func (d *DBClustersClient) DescribeDBClusters(input *awsrds.DescribeDBClustersInput) (*awsrds.DescribeDBClustersOutput, error) {
	d.DescribeDBClustersCall.CallCount++

	return d.DescribeDBClustersCall.Returns.Output, d.DescribeDBClustersCall.Returns.Error
}
