package compute

import (
	"fmt"

	"github.com/genevieve/leftovers/common"
	gcpcompute "google.golang.org/api/compute/v1"
)

//go:generate faux --interface instanceGroupsClient --output fakes/instance_groups_client.go
type instanceGroupsClient interface {
	ListInstanceGroups(zone string) ([]*gcpcompute.InstanceGroup, error)
	DeleteInstanceGroup(zone, instanceGroup string) error
}

type InstanceGroups struct {
	client instanceGroupsClient
	logger logger
	zones  map[string]string
}

func NewInstanceGroups(client instanceGroupsClient, logger logger, zones map[string]string) InstanceGroups {
	return InstanceGroups{
		client: client,
		logger: logger,
		zones:  zones,
	}
}

func (i InstanceGroups) List(filter string, regex bool) ([]common.Deletable, error) {
	groups := []*gcpcompute.InstanceGroup{}
	for _, zone := range i.zones {
		i.logger.Debugf("Listing Instance Groups for Zone %s...\n", zone)
		l, err := i.client.ListInstanceGroups(zone)
		if err != nil {
			return nil, fmt.Errorf("List Instance Groups for zone %s: %s", zone, err)
		}

		groups = append(groups, l...)
	}

	var resources []common.Deletable
	for _, group := range groups {
		resource := NewInstanceGroup(i.client, group.Name, i.zones[group.Zone])

		if !common.ResourceMatches(resource.Name(), filter, regex) {
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

func (i InstanceGroups) Type() string {
	return "instance-group"
}
