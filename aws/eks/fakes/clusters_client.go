package fakes

import (
	awseks "github.com/aws/aws-sdk-go/service/eks"
)

type ClustersClient struct {
	ListClustersCall struct {
		CallCount int
		Receives  struct {
			Input *awseks.ListClustersInput
		}
		Returns struct {
			Output *awseks.ListClustersOutput
			Error  error
		}
	}

	DeleteClusterCall struct {
		CallCount int
		Receives  struct {
			Input *awseks.DeleteClusterInput
		}
		Returns struct {
			Output *awseks.DeleteClusterOutput
			Error  error
		}
	}
}

func (c *ClustersClient) ListClusters(input *awseks.ListClustersInput) (*awseks.ListClustersOutput, error) {
	c.ListClustersCall.CallCount++
	c.ListClustersCall.Receives.Input = input

	return c.ListClustersCall.Returns.Output, c.ListClustersCall.Returns.Error
}

func (c *ClustersClient) DeleteCluster(input *awseks.DeleteClusterInput) (*awseks.DeleteClusterOutput, error) {
	c.DeleteClusterCall.CallCount++
	c.DeleteClusterCall.Receives.Input = input

	return c.DeleteClusterCall.Returns.Output, c.DeleteClusterCall.Returns.Error
}
