package openstack

import (
	"fmt"

	"github.com/gophercloud/gophercloud/openstack/imageservice/v2/images"
	"github.com/gophercloud/gophercloud/pagination"
)

type imageAPI interface {
	GetImagesPager() pagination.Pager
	PagerToPage(pager pagination.Pager) (pagination.Page, error)
	PageToImages(page pagination.Page) ([]images.Image, error)
	Delete(imageID string) error
}

type ImagesClient struct {
	imageAPI imageAPI
}

func NewImagesClient(imageAPI imageAPI) ImagesClient {
	return ImagesClient{
		imageAPI: imageAPI,
	}
}

func (ic ImagesClient) List() ([]images.Image, error) {
	pager := ic.imageAPI.GetImagesPager()
	if pager.Err != nil {
		return nil, fmt.Errorf("get images pager: %s", pager.Err)
	}

	page, err := ic.imageAPI.PagerToPage(pager)
	if err != nil {
		return nil, fmt.Errorf("pager to page: %s", err)
	}

	imgs, err := ic.imageAPI.PageToImages(page)
	if err != nil {
		return nil, fmt.Errorf("page to images: %s", err)
	}

	return imgs, err
}

func (ic ImagesClient) Delete(imageID string) error {
	return ic.imageAPI.Delete(imageID)
}
