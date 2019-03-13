package openstack

import (
	"github.com/genevieve/leftovers/common"
	"github.com/gophercloud/gophercloud/openstack/imageservice/v2/images"
)

type Images struct {
	imageServiceClient ImageServiceClient
	logger             logger
}

type ImageServiceClient interface {
	List() ([]images.Image, error)
	Delete(imageID string) error
}

func NewImages(imageServiceClient ImageServiceClient, logger logger) Images {
	return Images{
		imageServiceClient: imageServiceClient,
		logger:             logger,
	}
}

func (images Images) List() ([]common.Deletable, error) {
	res, err := images.imageServiceClient.List()
	if err != nil {
		return nil, err
	}
	var deletables []common.Deletable
	for _, resource := range res {
		deletable := NewImage(resource.Name, resource.ID, images.imageServiceClient)
		confirm := images.logger.PromptWithDetails(deletable.Type(), deletable.Name())
		if confirm {
			deletables = append(deletables, deletable)
		}
	}
	return deletables, err
}

func (images Images) Type() string {
	return "Image"
}
