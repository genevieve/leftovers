package fakes

import "github.com/genevieve/leftovers/openstack"

type VolumesServiceProvider struct {
	GetVolumesListerCall struct {
		CallCount int
		Returns   struct {
			VolumesLister openstack.VolumesLister
		}
	}

	GetVolumesDeleterCall struct {
		CallCount int
		Returns   struct {
			VolumesDeleter openstack.VolumesDeleter
		}
	}
}

func (v *VolumesServiceProvider) GetVolumesLister() openstack.VolumesLister {
	v.GetVolumesListerCall.CallCount++

	return v.GetVolumesListerCall.Returns.VolumesLister
}

func (v *VolumesServiceProvider) GetVolumesDeleter() openstack.VolumesDeleter {
	v.GetVolumesDeleterCall.CallCount++

	return v.GetVolumesDeleterCall.Returns.VolumesDeleter
}
