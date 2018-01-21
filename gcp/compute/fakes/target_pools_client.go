package fakes

import gcpcompute "google.golang.org/api/compute/v1"

type TargetPoolsClient struct {
	ListTargetPoolsCall struct {
		CallCount int
		Receives  struct {
			Region string
			Filter string
		}
		Returns struct {
			Output *gcpcompute.TargetPoolList
			Error  error
		}
	}

	DeleteTargetPoolCall struct {
		CallCount int
		Receives  struct {
			Region     string
			TargetPool string
		}
		Returns struct {
			Error error
		}
	}
}

func (n *TargetPoolsClient) ListTargetPools(region, filter string) (*gcpcompute.TargetPoolList, error) {
	n.ListTargetPoolsCall.CallCount++
	n.ListTargetPoolsCall.Receives.Region = region
	n.ListTargetPoolsCall.Receives.Filter = filter

	return n.ListTargetPoolsCall.Returns.Output, n.ListTargetPoolsCall.Returns.Error
}

func (n *TargetPoolsClient) DeleteTargetPool(region, targetPool string) error {
	n.DeleteTargetPoolCall.CallCount++
	n.DeleteTargetPoolCall.Receives.Region = region
	n.DeleteTargetPoolCall.Receives.TargetPool = targetPool

	return n.DeleteTargetPoolCall.Returns.Error
}
