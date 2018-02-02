package fakes

type RecordSets struct {
	DeleteCall struct {
		CallCount int
		Receives  struct {
			ManagedZone string
		}
		Returns struct {
			Error error
		}
	}
}

func (r *RecordSets) Delete(managedZone string) error {
	r.DeleteCall.CallCount++
	r.DeleteCall.Receives.ManagedZone = managedZone

	return r.DeleteCall.Returns.Error
}
