package route53

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	awsroute53 "github.com/aws/aws-sdk-go/service/route53"
)

type HostedZone struct {
	client     hostedZonesClient
	id         *string
	identifier string
}

func NewHostedZone(client hostedZonesClient, id, name *string) HostedZone {
	return HostedZone{
		client:     client,
		id:         id,
		identifier: *name,
	}
}

func (h HostedZone) Delete() error {
	var (
		resourceRecordSets []*awsroute53.ResourceRecordSet
		nextRecord         *string
	)

	for isTruncated := true; isTruncated; {
		output, err := h.client.ListResourceRecordSets(&awsroute53.ListResourceRecordSetsInput{
			HostedZoneId:    h.id,
			StartRecordName: nextRecord,
		})
		if err != nil {
			return fmt.Errorf("List Resource Record Sets: %s", err)
		}

		resourceRecordSets = append(resourceRecordSets, output.ResourceRecordSets...)

		isTruncated = *output.IsTruncated
		nextRecord = output.NextRecordName
	}

	var changes []*awsroute53.Change
	for _, record := range resourceRecordSets {
		if *record.Type == "NS" || *record.Type == "SOA" {
			continue
		}
		changes = append(changes, &awsroute53.Change{
			Action:            aws.String("DELETE"),
			ResourceRecordSet: record,
		})
	}

	if len(changes) > 0 {
		_, err := h.client.ChangeResourceRecordSets(&awsroute53.ChangeResourceRecordSetsInput{
			HostedZoneId: h.id,
			ChangeBatch:  &awsroute53.ChangeBatch{Changes: changes},
		})
		if err != nil {
			return fmt.Errorf("Delete Resource Record Sets: %s", err)
		}
	}

	_, err := h.client.DeleteHostedZone(&awsroute53.DeleteHostedZoneInput{Id: h.id})
	if err != nil {
		return fmt.Errorf("Delete: %s", err)
	}

	return nil
}

func (h HostedZone) Name() string {
	return h.identifier
}

func (h HostedZone) Type() string {
	return "Route53 Hosted Zone"
}
