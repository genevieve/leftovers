package fakes

import gcpdns "google.golang.org/api/dns/v1"

type RecordSetsClient struct {
	ListRecordSetsCall struct {
		CallCount int
		Receives  struct {
			ManagedZone string
		}
		Returns struct {
			Output *gcpdns.ResourceRecordSetsListResponse
			Error  error
		}
	}

	DeleteRecordSetsCall struct {
		CallCount int
		Receives  struct {
			ManagedZone string
			Change      *gcpdns.Change
		}
		Returns struct {
			Error error
		}
	}
}

func (r *RecordSetsClient) ListRecordSets(managedZone string) (*gcpdns.ResourceRecordSetsListResponse, error) {
	r.ListRecordSetsCall.CallCount++
	r.ListRecordSetsCall.Receives.ManagedZone = managedZone

	return r.ListRecordSetsCall.Returns.Output, r.ListRecordSetsCall.Returns.Error
}

func (r *RecordSetsClient) DeleteRecordSets(managedZone string, change *gcpdns.Change) error {
	r.DeleteRecordSetsCall.CallCount++
	r.DeleteRecordSetsCall.Receives.ManagedZone = managedZone
	r.DeleteRecordSetsCall.Receives.Change = change

	return r.DeleteRecordSetsCall.Returns.Error
}
