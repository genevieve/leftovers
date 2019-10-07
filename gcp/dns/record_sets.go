package dns

import (
	"fmt"

	gcpdns "google.golang.org/api/dns/v1"
)

//go:generate faux --interface recordSetsClient --output fakes/record_sets_client.go
type recordSetsClient interface {
	ListRecordSets(managedZone string) (*gcpdns.ResourceRecordSetsListResponse, error)
	DeleteRecordSets(managedZone string, change *gcpdns.Change) error
}

type RecordSets struct {
	client recordSetsClient
	logger logger
}

func NewRecordSets(client recordSetsClient, logger logger) RecordSets {
	return RecordSets{
		client: client,
		logger: logger,
	}
}

func (r RecordSets) Delete(managedZone string) error {
	r.logger.Debugln("Listing DNS Record Sets...")
	recordSets, err := r.client.ListRecordSets(managedZone)
	if err != nil {
		return fmt.Errorf("Listing DNS Record Sets: %s", err)
	}

	deletions := []*gcpdns.ResourceRecordSet{}
	for _, record := range recordSets.Rrsets {
		if record.Type == "NS" || record.Type == "SOA" {
			continue
		}

		deletions = append(deletions, record)
	}

	if len(deletions) > 0 {
		err = r.client.DeleteRecordSets(managedZone, &gcpdns.Change{
			Deletions: deletions,
		})
		if err != nil {
			return fmt.Errorf("Delete record sets: %s", err)
		}
	}
	return nil
}
