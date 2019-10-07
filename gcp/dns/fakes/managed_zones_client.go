package fakes

import (
	"sync"

	gcpdns "google.golang.org/api/dns/v1"
)

type ManagedZonesClient struct {
	DeleteManagedZoneCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Zone string
		}
		Returns struct {
			Error error
		}
		Stub func(string) error
	}
	ListManagedZonesCall struct {
		sync.Mutex
		CallCount int
		Returns   struct {
			ManagedZonesListResponse *gcpdns.ManagedZonesListResponse
			Error                    error
		}
		Stub func() (*gcpdns.ManagedZonesListResponse, error)
	}
}

func (f *ManagedZonesClient) DeleteManagedZone(param1 string) error {
	f.DeleteManagedZoneCall.Lock()
	defer f.DeleteManagedZoneCall.Unlock()
	f.DeleteManagedZoneCall.CallCount++
	f.DeleteManagedZoneCall.Receives.Zone = param1
	if f.DeleteManagedZoneCall.Stub != nil {
		return f.DeleteManagedZoneCall.Stub(param1)
	}
	return f.DeleteManagedZoneCall.Returns.Error
}
func (f *ManagedZonesClient) ListManagedZones() (*gcpdns.ManagedZonesListResponse, error) {
	f.ListManagedZonesCall.Lock()
	defer f.ListManagedZonesCall.Unlock()
	f.ListManagedZonesCall.CallCount++
	if f.ListManagedZonesCall.Stub != nil {
		return f.ListManagedZonesCall.Stub()
	}
	return f.ListManagedZonesCall.Returns.ManagedZonesListResponse, f.ListManagedZonesCall.Returns.Error
}
