package fakes

import "sync"

type RolePolicies struct {
	DeleteCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			RoleName string
		}
		Returns struct {
			Error error
		}
		Stub func(string) error
	}
}

func (f *RolePolicies) Delete(param1 string) error {
	f.DeleteCall.Lock()
	defer f.DeleteCall.Unlock()
	f.DeleteCall.CallCount++
	f.DeleteCall.Receives.RoleName = param1
	if f.DeleteCall.Stub != nil {
		return f.DeleteCall.Stub(param1)
	}
	return f.DeleteCall.Returns.Error
}
