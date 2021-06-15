package compute

import (
	"fmt"

	"github.com/genevieve/leftovers/common"
	gcpcompute "google.golang.org/api/compute/v1"
)

//go:generate faux --interface urlMapsClient --output fakes/url_maps_client.go
type urlMapsClient interface {
	ListUrlMaps() (*gcpcompute.UrlMapList, error)
	DeleteUrlMap(urlMap string) error
}

type UrlMaps struct {
	client urlMapsClient
	logger logger
}

func NewUrlMaps(client urlMapsClient, logger logger) UrlMaps {
	return UrlMaps{
		client: client,
		logger: logger,
	}
}

func (u UrlMaps) List(filter string, regex bool) ([]common.Deletable, error) {
	u.logger.Debugln("Listing Url Maps...")
	urlMaps, err := u.client.ListUrlMaps()
	if err != nil {
		return nil, fmt.Errorf("List Url Maps: %s", err)
	}

	var resources []common.Deletable

	for _, urlMap := range urlMaps.Items {
		resource := NewUrlMap(u.client, urlMap.Name)

		if !common.ResourceMatches(resource.Name(), filter, regex) {
			continue
		}

		proceed := u.logger.PromptWithDetails(resource.Type(), resource.Name())
		if !proceed {
			continue
		}

		resources = append(resources, resource)
	}

	return resources, nil
}

func (u UrlMaps) Type() string {
	return "url-map"
}
