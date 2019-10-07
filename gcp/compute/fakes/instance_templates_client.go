package fakes

import (
	"sync"

	gcp "google.golang.org/api/compute/v1"
)

type InstanceTemplatesClient struct {
	DeleteInstanceTemplateCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Template string
		}
		Returns struct {
			Error error
		}
		Stub func(string) error
	}
	ListInstanceTemplatesCall struct {
		sync.Mutex
		CallCount int
		Returns   struct {
			InstanceTemplateSlice []*gcp.InstanceTemplate
			Error                 error
		}
		Stub func() ([]*gcp.InstanceTemplate, error)
	}
}

func (f *InstanceTemplatesClient) DeleteInstanceTemplate(param1 string) error {
	f.DeleteInstanceTemplateCall.Lock()
	defer f.DeleteInstanceTemplateCall.Unlock()
	f.DeleteInstanceTemplateCall.CallCount++
	f.DeleteInstanceTemplateCall.Receives.Template = param1
	if f.DeleteInstanceTemplateCall.Stub != nil {
		return f.DeleteInstanceTemplateCall.Stub(param1)
	}
	return f.DeleteInstanceTemplateCall.Returns.Error
}
func (f *InstanceTemplatesClient) ListInstanceTemplates() ([]*gcp.InstanceTemplate, error) {
	f.ListInstanceTemplatesCall.Lock()
	defer f.ListInstanceTemplatesCall.Unlock()
	f.ListInstanceTemplatesCall.CallCount++
	if f.ListInstanceTemplatesCall.Stub != nil {
		return f.ListInstanceTemplatesCall.Stub()
	}
	return f.ListInstanceTemplatesCall.Returns.InstanceTemplateSlice, f.ListInstanceTemplatesCall.Returns.Error
}
