package fakes

import (
	awsroute53 "github.com/aws/aws-sdk-go/service/route53"
)

type RecordSets struct {
	GetCall struct {
		CallCount int
		Receives  struct {
			HostedZoneId *string
		}
		Returns struct {
			Records []*awsroute53.ResourceRecordSet
			Error   error
		}
	}

	DeleteCall struct {
		CallCount int
		Receives  struct {
			HostedZoneId   *string
			HostedZoneName string
			Records        []*awsroute53.ResourceRecordSet
		}
		Returns struct {
			Error error
		}
	}
}

func (r *RecordSets) Get(hostedZoneId *string) ([]*awsroute53.ResourceRecordSet, error) {
	r.GetCall.CallCount++
	r.GetCall.Receives.HostedZoneId = hostedZoneId

	return r.GetCall.Returns.Records, r.GetCall.Returns.Error
}

func (r *RecordSets) Delete(hostedZoneId *string, hostedZoneName string, records []*awsroute53.ResourceRecordSet) error {
	r.DeleteCall.CallCount++
	r.DeleteCall.Receives.HostedZoneId = hostedZoneId
	r.DeleteCall.Receives.HostedZoneName = hostedZoneName
	r.DeleteCall.Receives.Records = records

	return r.DeleteCall.Returns.Error
}
