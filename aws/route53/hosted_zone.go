package route53

import (
	"fmt"
	"strings"

	awsroute53 "github.com/aws/aws-sdk-go/service/route53"
)

type HostedZone struct {
	client     hostedZonesClient
	id         *string
	identifier string
	recordSets recordSets
	filter     string
}

func NewHostedZone(client hostedZonesClient, id, name *string, recordSets recordSets, filter string) HostedZone {
	return HostedZone{
		client:     client,
		id:         id,
		identifier: *name,
		recordSets: recordSets,
		filter:     filter,
	}
}

func (h HostedZone) Delete() error {
	r, err := h.recordSets.Get(h.id)
	if err != nil {
		return fmt.Errorf("Get Record Sets: %s", err)
	}

	if strings.Contains(h.Name(), h.filter) {
		err = h.recordSets.DeleteAll(h.id, h.identifier, r)
		if err != nil {
			return fmt.Errorf("Delete All Record Sets: %s", err)
		}

		_, err = h.client.DeleteHostedZone(&awsroute53.DeleteHostedZoneInput{Id: h.id})
		if err != nil {
			return fmt.Errorf("Delete: %s", err)
		}
	} else {
		err = h.recordSets.DeleteWithFilter(h.id, h.identifier, r, h.filter)
		if err != nil {
			return fmt.Errorf("Delete Record Sets With Filter: %s", err)
		}
	}

	return nil
}

func (h HostedZone) Name() string {
	return h.identifier
}

func (h HostedZone) Type() string {
	return "Route53 Hosted Zone"
}
