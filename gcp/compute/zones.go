package compute

import gcpcompute "google.golang.org/api/compute/v1"

type zonesClient interface {
	ListZones() (*gcpcompute.ZoneList, error)
}
