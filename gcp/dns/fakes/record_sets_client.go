package fakes

import (
	"sync"

	gcpdns "google.golang.org/api/dns/v1"
)

type RecordSetsClient struct {
	DeleteRecordSetsCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			ManagedZone string
			Change      *gcpdns.Change
		}
		Returns struct {
			Error error
		}
		Stub func(string, *gcpdns.Change) error
	}
	ListRecordSetsCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			ManagedZone string
		}
		Returns struct {
			ResourceRecordSetsListResponse *gcpdns.ResourceRecordSetsListResponse
			Error                          error
		}
		Stub func(string) (*gcpdns.ResourceRecordSetsListResponse, error)
	}
}

func (f *RecordSetsClient) DeleteRecordSets(param1 string, param2 *gcpdns.Change) error {
	f.DeleteRecordSetsCall.Lock()
	defer f.DeleteRecordSetsCall.Unlock()
	f.DeleteRecordSetsCall.CallCount++
	f.DeleteRecordSetsCall.Receives.ManagedZone = param1
	f.DeleteRecordSetsCall.Receives.Change = param2
	if f.DeleteRecordSetsCall.Stub != nil {
		return f.DeleteRecordSetsCall.Stub(param1, param2)
	}
	return f.DeleteRecordSetsCall.Returns.Error
}
func (f *RecordSetsClient) ListRecordSets(param1 string) (*gcpdns.ResourceRecordSetsListResponse, error) {
	f.ListRecordSetsCall.Lock()
	defer f.ListRecordSetsCall.Unlock()
	f.ListRecordSetsCall.CallCount++
	f.ListRecordSetsCall.Receives.ManagedZone = param1
	if f.ListRecordSetsCall.Stub != nil {
		return f.ListRecordSetsCall.Stub(param1)
	}
	return f.ListRecordSetsCall.Returns.ResourceRecordSetsListResponse, f.ListRecordSetsCall.Returns.Error
}
