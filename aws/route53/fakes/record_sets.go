package fakes

import (
	"sync"

	awsroute53 "github.com/aws/aws-sdk-go/service/route53"
)

type RecordSets struct {
	DeleteAllCall struct {
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
	DeleteWithFilterCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			HostedZoneId   *string
			HostedZoneName string
			RecordSets     []*awsroute53.ResourceRecordSet
			Filter         string
		}
		Returns struct {
			Error error
		}
		Stub func(*string, string, []*awsroute53.ResourceRecordSet, string) error
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

func (f *RecordSets) DeleteAll(param1 *string, param2 string, param3 []*awsroute53.ResourceRecordSet) error {
	f.DeleteAllCall.Lock()
	defer f.DeleteAllCall.Unlock()
	f.DeleteAllCall.CallCount++
	f.DeleteAllCall.Receives.HostedZoneId = param1
	f.DeleteAllCall.Receives.HostedZoneName = param2
	f.DeleteAllCall.Receives.RecordSets = param3
	if f.DeleteAllCall.Stub != nil {
		return f.DeleteAllCall.Stub(param1, param2, param3)
	}
	return f.DeleteAllCall.Returns.Error
}
func (f *RecordSets) DeleteWithFilter(param1 *string, param2 string, param3 []*awsroute53.ResourceRecordSet, param4 string) error {
	f.DeleteWithFilterCall.Lock()
	defer f.DeleteWithFilterCall.Unlock()
	f.DeleteWithFilterCall.CallCount++
	f.DeleteWithFilterCall.Receives.HostedZoneId = param1
	f.DeleteWithFilterCall.Receives.HostedZoneName = param2
	f.DeleteWithFilterCall.Receives.RecordSets = param3
	f.DeleteWithFilterCall.Receives.Filter = param4
	if f.DeleteWithFilterCall.Stub != nil {
		return f.DeleteWithFilterCall.Stub(param1, param2, param3, param4)
	}
	return f.DeleteWithFilterCall.Returns.Error
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
