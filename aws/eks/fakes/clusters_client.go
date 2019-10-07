package fakes

import (
	"sync"

	awseks "github.com/aws/aws-sdk-go/service/eks"
)

type ClustersClient struct {
	DeleteClusterCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DeleteClusterInput *awseks.DeleteClusterInput
		}
		Returns struct {
			DeleteClusterOutput *awseks.DeleteClusterOutput
			Error               error
		}
		Stub func(*awseks.DeleteClusterInput) (*awseks.DeleteClusterOutput, error)
	}
	ListClustersCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			ListClustersInput *awseks.ListClustersInput
		}
		Returns struct {
			ListClustersOutput *awseks.ListClustersOutput
			Error              error
		}
		Stub func(*awseks.ListClustersInput) (*awseks.ListClustersOutput, error)
	}
}

func (f *ClustersClient) DeleteCluster(param1 *awseks.DeleteClusterInput) (*awseks.DeleteClusterOutput, error) {
	f.DeleteClusterCall.Lock()
	defer f.DeleteClusterCall.Unlock()
	f.DeleteClusterCall.CallCount++
	f.DeleteClusterCall.Receives.DeleteClusterInput = param1
	if f.DeleteClusterCall.Stub != nil {
		return f.DeleteClusterCall.Stub(param1)
	}
	return f.DeleteClusterCall.Returns.DeleteClusterOutput, f.DeleteClusterCall.Returns.Error
}
func (f *ClustersClient) ListClusters(param1 *awseks.ListClustersInput) (*awseks.ListClustersOutput, error) {
	f.ListClustersCall.Lock()
	defer f.ListClustersCall.Unlock()
	f.ListClustersCall.CallCount++
	f.ListClustersCall.Receives.ListClustersInput = param1
	if f.ListClustersCall.Stub != nil {
		return f.ListClustersCall.Stub(param1)
	}
	return f.ListClustersCall.Returns.ListClustersOutput, f.ListClustersCall.Returns.Error
}
