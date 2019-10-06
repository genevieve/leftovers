package fakes

import "sync"

type BucketManager struct {
	IsInRegionCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Bucket string
		}
		Returns struct {
			Bool bool
		}
		Stub func(string) bool
	}
}

func (f *BucketManager) IsInRegion(param1 string) bool {
	f.IsInRegionCall.Lock()
	defer f.IsInRegionCall.Unlock()
	f.IsInRegionCall.CallCount++
	f.IsInRegionCall.Receives.Bucket = param1
	if f.IsInRegionCall.Stub != nil {
		return f.IsInRegionCall.Stub(param1)
	}
	return f.IsInRegionCall.Returns.Bool
}
