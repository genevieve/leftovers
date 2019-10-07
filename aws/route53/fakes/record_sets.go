package fakes

import (
	"sync"

	awsroute53 "github.com/aws/aws-sdk-go/service/route53"
)

type RecordSets struct {
	DeleteCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			HostedZoneId   *string
			HostedZoneName string
			RecordSets     []*awsroute53.ResourceRecordSet
		}
		Returns struct {
			Error error
		}
		Stub func(*string, string, []*awsroute53.ResourceRecordSet) error
	}
	GetCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			HostedZoneId *string
		}
		Returns struct {
			ResourceRecordSetSlice []*awsroute53.ResourceRecordSet
			Error                  error
		}
		Stub func(*string) ([]*awsroute53.ResourceRecordSet, error)
	}
}

func (f *RecordSets) Delete(param1 *string, param2 string, param3 []*awsroute53.ResourceRecordSet) error {
	f.DeleteCall.Lock()
	defer f.DeleteCall.Unlock()
	f.DeleteCall.CallCount++
	f.DeleteCall.Receives.HostedZoneId = param1
	f.DeleteCall.Receives.HostedZoneName = param2
	f.DeleteCall.Receives.RecordSets = param3
	if f.DeleteCall.Stub != nil {
		return f.DeleteCall.Stub(param1, param2, param3)
	}
	return f.DeleteCall.Returns.Error
}
func (f *RecordSets) Get(param1 *string) ([]*awsroute53.ResourceRecordSet, error) {
	f.GetCall.Lock()
	defer f.GetCall.Unlock()
	f.GetCall.CallCount++
	f.GetCall.Receives.HostedZoneId = param1
	if f.GetCall.Stub != nil {
		return f.GetCall.Stub(param1)
	}
	return f.GetCall.Returns.ResourceRecordSetSlice, f.GetCall.Returns.Error
}
