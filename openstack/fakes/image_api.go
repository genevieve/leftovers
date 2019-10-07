package fakes

import (
	"sync"

	"github.com/gophercloud/gophercloud/openstack/imageservice/v2/images"
	"github.com/gophercloud/gophercloud/pagination"
)

type ImageAPI struct {
	DeleteCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			ImageID string
		}
		Returns struct {
			Error error
		}
		Stub func(string) error
	}
	GetImagesPagerCall struct {
		sync.Mutex
		CallCount int
		Returns   struct {
			Pager pagination.Pager
		}
		Stub func() pagination.Pager
	}
	PageToImagesCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Page pagination.Page
		}
		Returns struct {
			ImageSlice []images.Image
			Error      error
		}
		Stub func(pagination.Page) ([]images.Image, error)
	}
	PagerToPageCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Pager pagination.Pager
		}
		Returns struct {
			Page  pagination.Page
			Error error
		}
		Stub func(pagination.Pager) (pagination.Page, error)
	}
}

func (f *ImageAPI) Delete(param1 string) error {
	f.DeleteCall.Lock()
	defer f.DeleteCall.Unlock()
	f.DeleteCall.CallCount++
	f.DeleteCall.Receives.ImageID = param1
	if f.DeleteCall.Stub != nil {
		return f.DeleteCall.Stub(param1)
	}
	return f.DeleteCall.Returns.Error
}
func (f *ImageAPI) GetImagesPager() pagination.Pager {
	f.GetImagesPagerCall.Lock()
	defer f.GetImagesPagerCall.Unlock()
	f.GetImagesPagerCall.CallCount++
	if f.GetImagesPagerCall.Stub != nil {
		return f.GetImagesPagerCall.Stub()
	}
	return f.GetImagesPagerCall.Returns.Pager
}
func (f *ImageAPI) PageToImages(param1 pagination.Page) ([]images.Image, error) {
	f.PageToImagesCall.Lock()
	defer f.PageToImagesCall.Unlock()
	f.PageToImagesCall.CallCount++
	f.PageToImagesCall.Receives.Page = param1
	if f.PageToImagesCall.Stub != nil {
		return f.PageToImagesCall.Stub(param1)
	}
	return f.PageToImagesCall.Returns.ImageSlice, f.PageToImagesCall.Returns.Error
}
func (f *ImageAPI) PagerToPage(param1 pagination.Pager) (pagination.Page, error) {
	f.PagerToPageCall.Lock()
	defer f.PagerToPageCall.Unlock()
	f.PagerToPageCall.CallCount++
	f.PagerToPageCall.Receives.Pager = param1
	if f.PagerToPageCall.Stub != nil {
		return f.PagerToPageCall.Stub(param1)
	}
	return f.PagerToPageCall.Returns.Page, f.PagerToPageCall.Returns.Error
}
