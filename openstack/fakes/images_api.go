package fakes

import (
	"github.com/gophercloud/gophercloud/openstack/imageservice/v2/images"
	"github.com/gophercloud/gophercloud/pagination"
)

type ImageAPI struct {
	GetImagePagerCall struct {
		CallCount int
		Returns   struct {
			Pager pagination.Pager
		}
	}
	PagerToPageCall struct {
		CallCount int
		Returns   struct {
			Page  pagination.Page
			Error error
		}
		Receives struct {
			Pager pagination.Pager
		}
	}
	PageToImagesCall struct {
		CallCount int
		Returns   struct {
			Images []images.Image
			Error  error
		}
		Receives struct {
			Page pagination.Page
		}
	}
	DeleteCall struct {
		CallCount int
		Returns   struct {
			Error error
		}
		Receives struct {
			ImageID string
		}
	}
}

func (api *ImageAPI) GetImagesPager() pagination.Pager {
	api.GetImagePagerCall.CallCount++

	return api.GetImagePagerCall.Returns.Pager
}

func (api *ImageAPI) PagerToPage(pager pagination.Pager) (pagination.Page, error) {
	api.PagerToPageCall.CallCount++
	api.PagerToPageCall.Receives.Pager = pager

	return api.PagerToPageCall.Returns.Page, api.PagerToPageCall.Returns.Error
}

func (api *ImageAPI) PageToImages(page pagination.Page) ([]images.Image, error) {
	api.PageToImagesCall.CallCount++

	api.PageToImagesCall.Receives.Page = page
	return api.PageToImagesCall.Returns.Images, api.PageToImagesCall.Returns.Error
}

func (api *ImageAPI) Delete(imageID string) error {
	api.DeleteCall.CallCount++

	api.DeleteCall.Receives.ImageID = imageID
	return api.DeleteCall.Returns.Error
}
