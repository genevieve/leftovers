package compute

import (
	"fmt"
	"strings"

	gcpcompute "google.golang.org/api/compute/v1"
)

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

func (u UrlMaps) List(filter string) (map[string]string, error) {
	urlMaps, err := u.client.ListUrlMaps()
	if err != nil {
		return nil, fmt.Errorf("Listing url maps: %s", err)
	}

	delete := map[string]string{}
	for _, urlMap := range urlMaps.Items {
		if !strings.Contains(urlMap.Name, filter) {
			continue
		}

		proceed := u.logger.Prompt(fmt.Sprintf("Are you sure you want to delete url map %s?", urlMap.Name))
		if !proceed {
			continue
		}

		delete[urlMap.Name] = ""
	}

	return delete, nil
}

func (u UrlMaps) Delete(urlMaps map[string]string) {
	var resources []UrlMap
	for name, _ := range urlMaps {
		resources = append(resources, NewUrlMap(u.client, name))
	}

	u.delete(resources)
}

func (u UrlMaps) delete(resources []UrlMap) {
	for _, resource := range resources {
		err := resource.Delete()

		if err != nil {
			u.logger.Printf("%s\n", err)
		} else {
			u.logger.Printf("SUCCESS deleting url map %s\n", resource.name)
		}
	}
}
