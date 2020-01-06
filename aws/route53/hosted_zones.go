package route53

import (
	"fmt"
	"strings"

	awsroute53 "github.com/aws/aws-sdk-go/service/route53"
	"github.com/genevieve/leftovers/common"
)

//go:generate faux --interface hostedZonesClient --output fakes/hosted_zones_client.go
type hostedZonesClient interface {
	ListHostedZones(*awsroute53.ListHostedZonesInput) (*awsroute53.ListHostedZonesOutput, error)
	DeleteHostedZone(*awsroute53.DeleteHostedZoneInput) (*awsroute53.DeleteHostedZoneOutput, error)
}

type HostedZones struct {
	client     hostedZonesClient
	logger     logger
	recordSets recordSets
}

//go:generate faux --interface recordSets --output fakes/record_sets.go
type recordSets interface {
	Get(hostedZoneId *string) ([]*awsroute53.ResourceRecordSet, error)
	DeleteAll(hostedZoneId *string, hostedZoneName string, recordSets []*awsroute53.ResourceRecordSet) error
	DeleteWithFilter(hostedZoneId *string, hostedZoneName string, recordSets []*awsroute53.ResourceRecordSet, filter string) error
}

func NewHostedZones(client hostedZonesClient, logger logger, recordSets recordSets) HostedZones {
	return HostedZones{
		client:     client,
		logger:     logger,
		recordSets: recordSets,
	}
}

func (z HostedZones) List(filter string) ([]common.Deletable, error) {
	zones, err := z.client.ListHostedZones(&awsroute53.ListHostedZonesInput{})
	if err != nil {
		return nil, fmt.Errorf("List Route53 Hosted Zones: %s", err)
	}

	var resources []common.Deletable
	for _, zone := range zones.HostedZones {
		r := NewHostedZone(z.client, zone.Id, zone.Name, z.recordSets, filter)

		if !strings.Contains(r.Name(), filter) && !z.recordSetsContainFilter(zone.Id, filter) {
			continue
		}

		proceed := z.logger.PromptWithDetails(r.Type(), r.Name())
		if !proceed {
			continue
		}

		resources = append(resources, r)
	}

	return resources, nil
}

func (z HostedZones) Type() string {
	return "route53-hosted-zone"
}

// Check if any record sets in the hosted zone reference the filter.
func (z HostedZones) recordSetsContainFilter(hostedZoneId *string, filter string) bool {
	records, err := z.recordSets.Get(hostedZoneId)
	if err != nil {
		return false
	}

	for _, record := range records {
		if strings.Contains(*record.Name, filter) {
			return true
		}
	}

	return false
}
