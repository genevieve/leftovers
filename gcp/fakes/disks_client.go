package fakes

import compute "google.golang.org/api/compute/v1"

type DisksClient struct {
	ListDisksCall struct {
		CallCount int
		Returns   struct {
			Output *compute.DiskList
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
			Output *compute.Operation
			Error  error
		}
	}
}

func (n *DisksClient) ListDisks() (*compute.DiskList, error) {
	n.ListDisksCall.CallCount++

	return n.ListDisksCall.Returns.Output, n.ListDisksCall.Returns.Error
}

func (n *DisksClient) DeleteDisk(zone, disk string) (*compute.Operation, error) {
	n.DeleteDiskCall.CallCount++
	n.DeleteDiskCall.Receives.Zone = zone
	n.DeleteDiskCall.Receives.Disk = disk

	return n.DeleteDiskCall.Returns.Output, n.DeleteDiskCall.Returns.Error
}
