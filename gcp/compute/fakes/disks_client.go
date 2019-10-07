package fakes

import (
	"sync"

	gcpcompute "google.golang.org/api/compute/v1"
)

type DisksClient struct {
	DeleteDiskCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Zone string
			Disk string
		}
		Returns struct {
			Error error
		}
		Stub func(string, string) error
	}
	ListDisksCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Zone string
		}
		Returns struct {
			DiskSlice []*gcpcompute.Disk
			Error     error
		}
		Stub func(string) ([]*gcpcompute.Disk, error)
	}
}

func (f *DisksClient) DeleteDisk(param1 string, param2 string) error {
	f.DeleteDiskCall.Lock()
	defer f.DeleteDiskCall.Unlock()
	f.DeleteDiskCall.CallCount++
	f.DeleteDiskCall.Receives.Zone = param1
	f.DeleteDiskCall.Receives.Disk = param2
	if f.DeleteDiskCall.Stub != nil {
		return f.DeleteDiskCall.Stub(param1, param2)
	}
	return f.DeleteDiskCall.Returns.Error
}
func (f *DisksClient) ListDisks(param1 string) ([]*gcpcompute.Disk, error) {
	f.ListDisksCall.Lock()
	defer f.ListDisksCall.Unlock()
	f.ListDisksCall.CallCount++
	f.ListDisksCall.Receives.Zone = param1
	if f.ListDisksCall.Stub != nil {
		return f.ListDisksCall.Stub(param1)
	}
	return f.ListDisksCall.Returns.DiskSlice, f.ListDisksCall.Returns.Error
}
