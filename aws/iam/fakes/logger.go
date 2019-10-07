package fakes

import "sync"

type Logger struct {
	PrintfCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			M string
			A []interface {
			}
		}
		Stub func(string, ...interface {
		})
	}
	PromptWithDetailsCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			ResourceType string
			ResourceName string
		}
		Returns struct {
			Proceed bool
		}
		Stub func(string, string) bool
	}
}

func (f *Logger) Printf(param1 string, param2 ...interface {
}) {
	f.PrintfCall.Lock()
	defer f.PrintfCall.Unlock()
	f.PrintfCall.CallCount++
	f.PrintfCall.Receives.M = param1
	f.PrintfCall.Receives.A = param2
	if f.PrintfCall.Stub != nil {
		f.PrintfCall.Stub(param1, param2...)
	}
}
func (f *Logger) PromptWithDetails(param1 string, param2 string) bool {
	f.PromptWithDetailsCall.Lock()
	defer f.PromptWithDetailsCall.Unlock()
	f.PromptWithDetailsCall.CallCount++
	f.PromptWithDetailsCall.Receives.ResourceType = param1
	f.PromptWithDetailsCall.Receives.ResourceName = param2
	if f.PromptWithDetailsCall.Stub != nil {
		return f.PromptWithDetailsCall.Stub(param1, param2)
	}
	return f.PromptWithDetailsCall.Returns.Proceed
}
