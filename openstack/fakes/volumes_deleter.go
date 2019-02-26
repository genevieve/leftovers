package fakes

type VolumesDeleter struct {
	DeleteCall struct {
		CallCount int
		Receives  struct {
			VolumeID string
		}
		Returns struct {
			Error error
		}
	}
}

func (v *VolumesDeleter) Delete(volumeID string) error {
	v.DeleteCall.CallCount++
	v.DeleteCall.Receives.VolumeID = volumeID

	return v.DeleteCall.Returns.Error
}
