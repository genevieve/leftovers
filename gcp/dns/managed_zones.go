package dns

import (
	"fmt"
	"strings"

	"github.com/genevieve/leftovers/common"
	gcpdns "google.golang.org/api/dns/v1"
)

//go:generate faux --interface managedZonesClient --output fakes/managed_zones_client.go
type managedZonesClient interface {
	ListManagedZones() (*gcpdns.ManagedZonesListResponse, error)
	DeleteManagedZone(zone string) error
}

//go:generate faux --interface recordSets --output fakes/record_sets.go
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

func (m ManagedZones) List(filter string) ([]common.Deletable, error) {
	m.logger.Debugln("Listing DNS Managed Zones...")
	managedZones, err := m.client.ListManagedZones()
	if err != nil {
		return nil, fmt.Errorf("Listing DNS Managed Zones: %s", err)
	}

	var resources []common.Deletable
	for _, zone := range managedZones.ManagedZones {
		resource := NewManagedZone(m.client, m.recordSets, zone.Name)

		if !strings.Contains(resource.name, filter) {
			continue
		}

		proceed := m.logger.PromptWithDetails(resource.Type(), resource.Name())
		if !proceed {
			continue
		}

		resources = append(resources, resource)
	}

	return resources, nil
}

func (m ManagedZones) Type() string {
	return "managed-zone"
}
