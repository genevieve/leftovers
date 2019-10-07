package fakes

import "sync"

type ResourceTags struct {
	DeleteCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			FilterName  string
			FilterValue string
		}
		Returns struct {
			Error error
		}
		Stub func(string, string) error
	}
}

func (f *ResourceTags) Delete(param1 string, param2 string) error {
	f.DeleteCall.Lock()
	defer f.DeleteCall.Unlock()
	f.DeleteCall.CallCount++
	f.DeleteCall.Receives.FilterName = param1
	f.DeleteCall.Receives.FilterValue = param2
	if f.DeleteCall.Stub != nil {
		return f.DeleteCall.Stub(param1, param2)
	}
	return f.DeleteCall.Returns.Error
}
