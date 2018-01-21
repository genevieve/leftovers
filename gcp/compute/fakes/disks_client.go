package fakes

import gcpcompute "google.golang.org/api/compute/v1"

type DisksClient struct {
	ListDisksCall struct {
		CallCount int
		Receives  struct {
			Zone   string
			Filter string
		}
		Returns struct {
			Output *gcpcompute.DiskList
			Error  error
		}
	}

	DeleteDiskCall struct {
		CallCount int
		Receives  struct {
			Zone string
			Disk string
		}
		Returns struct {
			Error error
		}
	}
}

func (n *DisksClient) ListDisks(zone, filter string) (*gcpcompute.DiskList, error) {
	n.ListDisksCall.CallCount++
	n.ListDisksCall.Receives.Zone = zone
	n.ListDisksCall.Receives.Filter = filter

	return n.ListDisksCall.Returns.Output, n.ListDisksCall.Returns.Error
}

func (n *DisksClient) DeleteDisk(zone, disk string) error {
	n.DeleteDiskCall.CallCount++
	n.DeleteDiskCall.Receives.Zone = zone
	n.DeleteDiskCall.Receives.Disk = disk

	return n.DeleteDiskCall.Returns.Error
}
