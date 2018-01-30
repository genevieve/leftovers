package fakes

import gcpdns "google.golang.org/api/dns/v1"

type ManagedZonesClient struct {
	ListManagedZonesCall struct {
		CallCount int
		Returns   struct {
			Output *gcpdns.ManagedZonesListResponse
			Error  error
		}
	}

	DeleteManagedZoneCall struct {
		CallCount int
		Receives  struct {
			ManagedZone string
		}
		Returns struct {
			Error error
		}
	}
}

func (u *ManagedZonesClient) ListManagedZones() (*gcpdns.ManagedZonesListResponse, error) {
	u.ListManagedZonesCall.CallCount++

	return u.ListManagedZonesCall.Returns.Output, u.ListManagedZonesCall.Returns.Error
}

func (u *ManagedZonesClient) DeleteManagedZone(managedZone string) error {
	u.DeleteManagedZoneCall.CallCount++
	u.DeleteManagedZoneCall.Receives.ManagedZone = managedZone

	return u.DeleteManagedZoneCall.Returns.Error
}
