package gcp

import compute "google.golang.org/api/compute/v1"

type zonesClient interface {
	ListZones() (*compute.ZoneList, error)
}
