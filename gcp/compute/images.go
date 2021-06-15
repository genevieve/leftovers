package compute

import (
	"fmt"
	"github.com/genevieve/leftovers/common"
	gcpcompute "google.golang.org/api/compute/v1"
)

//go:generate faux --interface imagesClient --output fakes/images_client.go
type imagesClient interface {
	ListImages() ([]*gcpcompute.Image, error)
	DeleteImage(image string) error
}

type Images struct {
	client imagesClient
	logger logger
}

func NewImages(client imagesClient, logger logger) Images {
	return Images{
		client: client,
		logger: logger,
	}
}

func (i Images) List(filter string, regex bool) ([]common.Deletable, error) {
	i.logger.Debugln("Listing Images...")

	images, err := i.client.ListImages()
	if err != nil {
		return nil, fmt.Errorf("List Images: %s", err)
	}

	var resources []common.Deletable
	for _, image := range images {
		resource := NewImage(i.client, image.Name)

		if !common.ResourceMatches(image.Name, filter, regex) {
			continue
		}

		proceed := i.logger.PromptWithDetails(resource.Type(), resource.Name())
		if !proceed {
			continue
		}

		resources = append(resources, resource)
	}

	return resources, nil
}

func (i Images) Type() string {
	return "image"
}
