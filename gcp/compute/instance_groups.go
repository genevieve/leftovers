package compute

import (
	"fmt"
	"strings"

	gcpcompute "google.golang.org/api/compute/v1"
)

type instanceGroupsClient interface {
	ListInstanceGroups(zone string) (*gcpcompute.InstanceGroupList, error)
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

func (s InstanceGroups) Delete(filter string) error {
	var groups []*gcpcompute.InstanceGroup
	for _, zone := range s.zones {
		l, err := s.client.ListInstanceGroups(zone)
		if err != nil {
			return fmt.Errorf("Listing instance groups for zone %s: %s", zone, err)
		}
		groups = append(groups, l.Items...)
	}

	for _, group := range groups {
		n := group.Name

		if !strings.Contains(n, filter) {
			continue
		}

		proceed := s.logger.Prompt(fmt.Sprintf("Are you sure you want to delete instance group %s?", n))
		if !proceed {
			continue
		}

		zoneName := s.zones[group.Zone]
		if err := s.client.DeleteInstanceGroup(zoneName, n); err != nil {
			s.logger.Printf("ERROR deleting instance group %s: %s\n", n, err)
		} else {
			s.logger.Printf("SUCCESS deleting instance group %s\n", n)
		}
	}

	return nil
}
