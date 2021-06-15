package route53

import (
	"fmt"
	awsroute53 "github.com/aws/aws-sdk-go/service/route53"
	"github.com/genevieve/leftovers/common"
)

type HostedZone struct {
	client     hostedZonesClient
	id         *string
	identifier string
	recordSets recordSets
	filter     string
	regex	   bool
}

func NewHostedZone(client hostedZonesClient, id, name *string, recordSets recordSets, filter string, regex bool) HostedZone {
	return HostedZone{
		client:     client,
		id:         id,
		identifier: *name,
		recordSets: recordSets,
		filter:     filter,
		regex:      regex,
	}
}

func (h HostedZone) Delete() error {
	r, err := h.recordSets.Get(h.id)
	if err != nil {
		return fmt.Errorf("Get Record Sets: %s", err)
	}

	if common.ResourceMatches(h.Name(), h.filter, h.regex) {
		err = h.recordSets.DeleteAll(h.id, h.identifier, r)
		if err != nil {
			return fmt.Errorf("Delete All Record Sets: %s", err)
		}

		_, err = h.client.DeleteHostedZone(&awsroute53.DeleteHostedZoneInput{Id: h.id})
		if err != nil {
			return fmt.Errorf("Delete: %s", err)
		}
	} else {
		err = h.recordSets.DeleteWithFilter(h.id, h.identifier, r, h.filter, h.regex)
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
