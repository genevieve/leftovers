package compute

import (
	"fmt"
	"strings"
	"sync"

	gcpcompute "google.golang.org/api/compute/v1"
)

type imagesClient interface {
	ListImages() (*gcpcompute.ImageList, error)
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

func (d Images) List(filter string) (map[string]string, error) {
	images, err := d.client.ListImages()
	if err != nil {
		return nil, fmt.Errorf("Listing images: %s", err)
	}

	delete := map[string]string{}
	for _, image := range images.Items {
		if !strings.Contains(image.Name, filter) {
			continue
		}

		proceed := d.logger.Prompt(fmt.Sprintf("Are you sure you want to delete image %s?", image.Name))
		if !proceed {
			continue
		}

		delete[image.Name] = ""
	}

	return delete, nil
}

func (d Images) Delete(images map[string]string) {
	var wg sync.WaitGroup

	for name, _ := range images {
		wg.Add(1)

		go func(name string) {
			err := d.client.DeleteImage(name)

			if err != nil {
				d.logger.Printf("ERROR deleting image %s: %s\n", name, err)
			} else {
				d.logger.Printf("SUCCESS deleting image %s\n", name)
			}

			wg.Done()
		}(name)
	}

	wg.Wait()
}
