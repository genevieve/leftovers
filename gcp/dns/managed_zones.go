package dns

import (
	"fmt"
	"strings"

	gcpdns "google.golang.org/api/dns/v1"
)

type managedZonesClient interface {
	ListManagedZones() (*gcpdns.ManagedZonesListResponse, error)
	DeleteManagedZone(zone string) error
}

type recordSets interface {
	Delete(managedZone string) error
}

type ManagedZones struct {
	client     managedZonesClient
	recordSets recordSets
	logger     logger
}

func NewManagedZones(client managedZonesClient, recordSets recordSets, logger logger) ManagedZones {
	return ManagedZones{
		client:     client,
		recordSets: recordSets,
		logger:     logger,
	}
}

func (m ManagedZones) List(filter string) (map[string]string, error) {
	managedZones, err := m.client.ListManagedZones()
	if err != nil {
		return nil, fmt.Errorf("Listing managed zones: %s", err)
	}

	delete := map[string]string{}
	for _, zone := range managedZones.ManagedZones {
		if !strings.Contains(zone.Name, filter) {
			continue
		}

		proceed := m.logger.Prompt(fmt.Sprintf("Are you sure you want to delete managed zone %s?", zone.Name))
		if !proceed {
			continue
		}

		delete[zone.Name] = ""
	}

	return delete, nil
}

func (m ManagedZones) Delete(zones map[string]string) {
	var resources []ManagedZone
	for name, _ := range zones {
		resources = append(resources, NewManagedZone(m.client, m.recordSets, name))
	}

	m.delete(resources)
}

func (m ManagedZones) delete(resources []ManagedZone) {
	for _, resource := range resources {
		err := resource.Delete()

		if err != nil {
			m.logger.Printf("%s\n", err)
		} else {
			m.logger.Printf("SUCCESS deleting managed zone %s\n", resource.name)
		}
	}
}
