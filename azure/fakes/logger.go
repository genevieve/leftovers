package fakes

import "sync"

type Logger struct {
	DebuglnCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Message string
		}
		Stub func(string)
	}
	NoConfirmCall struct {
		sync.Mutex
		CallCount int
		Stub      func()
	}
	PrintfCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Message string
			Args    []interface {
			}
		}
		Stub func(string, ...interface {
		})
	}
	PrintlnCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Message string
		}
		Stub func(string)
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

func (f *Logger) Debugln(param1 string) {
	f.DebuglnCall.Lock()
	defer f.DebuglnCall.Unlock()
	f.DebuglnCall.CallCount++
	f.DebuglnCall.Receives.Message = param1
	if f.DebuglnCall.Stub != nil {
		f.DebuglnCall.Stub(param1)
	}
}
func (f *Logger) NoConfirm() {
	f.NoConfirmCall.Lock()
	defer f.NoConfirmCall.Unlock()
	f.NoConfirmCall.CallCount++
	if f.NoConfirmCall.Stub != nil {
		f.NoConfirmCall.Stub()
	}
}
func (f *Logger) Printf(param1 string, param2 ...interface {
}) {
	f.PrintfCall.Lock()
	defer f.PrintfCall.Unlock()
	f.PrintfCall.CallCount++
	f.PrintfCall.Receives.Message = param1
	f.PrintfCall.Receives.Args = param2
	if f.PrintfCall.Stub != nil {
		f.PrintfCall.Stub(param1, param2...)
	}
}
func (f *Logger) Println(param1 string) {
	f.PrintlnCall.Lock()
	defer f.PrintlnCall.Unlock()
	f.PrintlnCall.CallCount++
	f.PrintlnCall.Receives.Message = param1
	if f.PrintlnCall.Stub != nil {
		f.PrintlnCall.Stub(param1)
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
