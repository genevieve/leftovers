package openstack

import (
	"fmt"
	"strings"

	"github.com/genevieve/leftovers/common"
	"github.com/gophercloud/gophercloud/openstack/imageservice/v2/images"
)

type Images struct {
	client imageServiceClient
	logger logger
}

//go:generate faux --interface imageServiceClient --output fakes/image_service_client.go
type imageServiceClient interface {
	List() ([]images.Image, error)
	Delete(id string) error
}

func NewImages(client imageServiceClient, logger logger) Images {
	return Images{
		client: client,
		logger: logger,
	}
}

func (i Images) List(filter string) ([]common.Deletable, error) {
	i.logger.Debugln("Listing Images...")

	images, err := i.client.List()
	if err != nil {
		return nil, fmt.Errorf("List Images: %s", err)
	}

	var resources []common.Deletable
	for _, image := range images {
		r := NewImage(image.Name, image.ID, i.client)

		if !strings.Contains(image.Name, filter) {
			continue
		}

		proceed := i.logger.PromptWithDetails(r.Type(), r.Name())
		if !proceed {
			continue
		}

		resources = append(resources, r)
	}

	return resources, err
}

func (i Images) Type() string {
	return "Image"
}
