package fakes

import (
	"sync"

	gcpcompute "google.golang.org/api/compute/v1"
)

type TargetPoolsClient struct {
	DeleteTargetPoolCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Region     string
			TargetPool string
		}
		Returns struct {
			Error error
		}
		Stub func(string, string) error
	}
	ListTargetPoolsCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Region string
		}
		Returns struct {
			TargetPoolList *gcpcompute.TargetPoolList
			Error          error
		}
		Stub func(string) (*gcpcompute.TargetPoolList, error)
	}
}

func (f *TargetPoolsClient) DeleteTargetPool(param1 string, param2 string) error {
	f.DeleteTargetPoolCall.Lock()
	defer f.DeleteTargetPoolCall.Unlock()
	f.DeleteTargetPoolCall.CallCount++
	f.DeleteTargetPoolCall.Receives.Region = param1
	f.DeleteTargetPoolCall.Receives.TargetPool = param2
	if f.DeleteTargetPoolCall.Stub != nil {
		return f.DeleteTargetPoolCall.Stub(param1, param2)
	}
	return f.DeleteTargetPoolCall.Returns.Error
}
func (f *TargetPoolsClient) ListTargetPools(param1 string) (*gcpcompute.TargetPoolList, error) {
	f.ListTargetPoolsCall.Lock()
	defer f.ListTargetPoolsCall.Unlock()
	f.ListTargetPoolsCall.CallCount++
	f.ListTargetPoolsCall.Receives.Region = param1
	if f.ListTargetPoolsCall.Stub != nil {
		return f.ListTargetPoolsCall.Stub(param1)
	}
	return f.ListTargetPoolsCall.Returns.TargetPoolList, f.ListTargetPoolsCall.Returns.Error
}
