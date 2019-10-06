package fakes

import (
	"sync"

	gcpcompute "google.golang.org/api/compute/v1"
)

type ImagesClient struct {
	DeleteImageCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Image string
		}
		Returns struct {
			Error error
		}
		Stub func(string) error
	}
	ListImagesCall struct {
		sync.Mutex
		CallCount int
		Returns   struct {
			ImageSlice []*gcpcompute.Image
			Error      error
		}
		Stub func() ([]*gcpcompute.Image, error)
	}
}

func (f *ImagesClient) DeleteImage(param1 string) error {
	f.DeleteImageCall.Lock()
	defer f.DeleteImageCall.Unlock()
	f.DeleteImageCall.CallCount++
	f.DeleteImageCall.Receives.Image = param1
	if f.DeleteImageCall.Stub != nil {
		return f.DeleteImageCall.Stub(param1)
	}
	return f.DeleteImageCall.Returns.Error
}
func (f *ImagesClient) ListImages() ([]*gcpcompute.Image, error) {
	f.ListImagesCall.Lock()
	defer f.ListImagesCall.Unlock()
	f.ListImagesCall.CallCount++
	if f.ListImagesCall.Stub != nil {
		return f.ListImagesCall.Stub()
	}
	return f.ListImagesCall.Returns.ImageSlice, f.ListImagesCall.Returns.Error
}
