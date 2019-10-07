package fakes

import (
	"sync"

	"github.com/gophercloud/gophercloud/openstack/imageservice/v2/images"
)

type ImageServiceClient struct {
	DeleteCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Id string
		}
		Returns struct {
			Error error
		}
		Stub func(string) error
	}
	ListCall struct {
		sync.Mutex
		CallCount int
		Returns   struct {
			ImageSlice []images.Image
			Error      error
		}
		Stub func() ([]images.Image, error)
	}
}

func (f *ImageServiceClient) Delete(param1 string) error {
	f.DeleteCall.Lock()
	defer f.DeleteCall.Unlock()
	f.DeleteCall.CallCount++
	f.DeleteCall.Receives.Id = param1
	if f.DeleteCall.Stub != nil {
		return f.DeleteCall.Stub(param1)
	}
	return f.DeleteCall.Returns.Error
}
func (f *ImageServiceClient) List() ([]images.Image, error) {
	f.ListCall.Lock()
	defer f.ListCall.Unlock()
	f.ListCall.CallCount++
	if f.ListCall.Stub != nil {
		return f.ListCall.Stub()
	}
	return f.ListCall.Returns.ImageSlice, f.ListCall.Returns.Error
}
