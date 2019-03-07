package fakes

type ComputeInstanceDeleter struct {
	DeleteCall struct {
		CallCount int
		Returns   struct {
			Error error
		}
		Receives struct {
			InstanceID string
		}
	}
}

func (deleter *ComputeInstanceDeleter) Delete(instanceID string) error {
	deleter.DeleteCall.CallCount++
	deleter.DeleteCall.Receives.InstanceID = instanceID

	return deleter.DeleteCall.Returns.Error
}
