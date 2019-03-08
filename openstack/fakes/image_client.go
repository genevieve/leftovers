package fakes

import "github.com/gophercloud/gophercloud/openstack/imageservice/v2/images"

type ImageClient struct {
	DeleteCall struct {
		CallCount int
		Returns   struct {
			Error error
		}
		Receives struct {
			ImageID string
		}
	}
	ListCall struct {
		CallCount int
		Returns   struct {
			Images []images.Image
			Error  error
		}
	}
}

func (client *ImageClient) Delete(imageID string) error {
	client.DeleteCall.CallCount++
	client.DeleteCall.Receives.ImageID = imageID

	return client.DeleteCall.Returns.Error
}

func (client *ImageClient) List() ([]images.Image, error) {
	client.ListCall.CallCount++

	return client.ListCall.Returns.Images, client.ListCall.Returns.Error
}
