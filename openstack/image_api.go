package openstack

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/imageservice/v2/images"
	"github.com/gophercloud/gophercloud/pagination"
)

type ImageAPI struct {
	serviceClient *gophercloud.ServiceClient
}

func (api ImageAPI) GetImagesPager() pagination.Pager {
	return images.List(api.serviceClient, images.ListOpts{})
}

func (api ImageAPI) PagerToPage(pager pagination.Pager) (pagination.Page, error) {
	return pager.AllPages()
}

func (api ImageAPI) PageToImages(page pagination.Page) ([]images.Image, error) {
	return images.ExtractImages(page)
}

func (api ImageAPI) Delete(imageID string) error {
	return images.Delete(api.serviceClient, imageID).ExtractErr()
}
