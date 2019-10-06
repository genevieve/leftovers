package fakes

import "sync"

type Page struct {
	GetBodyCall struct {
		sync.Mutex
		CallCount int
		Returns   struct {
			Interface interface {
			}
		}
		Stub func() interface {
		}
	}
	IsEmptyCall struct {
		sync.Mutex
		CallCount int
		Returns   struct {
			Bool  bool
			Error error
		}
		Stub func() (bool, error)
	}
	NextPageURLCall struct {
		sync.Mutex
		CallCount int
		Returns   struct {
			String string
			Error  error
		}
		Stub func() (string, error)
	}
}

func (f *Page) GetBody() interface {
} {
	f.GetBodyCall.Lock()
	defer f.GetBodyCall.Unlock()
	f.GetBodyCall.CallCount++
	if f.GetBodyCall.Stub != nil {
		return f.GetBodyCall.Stub()
	}
	return f.GetBodyCall.Returns.Interface
}
func (f *Page) IsEmpty() (bool, error) {
	f.IsEmptyCall.Lock()
	defer f.IsEmptyCall.Unlock()
	f.IsEmptyCall.CallCount++
	if f.IsEmptyCall.Stub != nil {
		return f.IsEmptyCall.Stub()
	}
	return f.IsEmptyCall.Returns.Bool, f.IsEmptyCall.Returns.Error
}
func (f *Page) NextPageURL() (string, error) {
	f.NextPageURLCall.Lock()
	defer f.NextPageURLCall.Unlock()
	f.NextPageURLCall.CallCount++
	if f.NextPageURLCall.Stub != nil {
		return f.NextPageURLCall.Stub()
	}
	return f.NextPageURLCall.Returns.String, f.NextPageURLCall.Returns.Error
}
