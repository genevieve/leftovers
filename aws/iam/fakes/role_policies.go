package fakes

type RolePolicies struct {
	DeleteCall struct {
		CallCount int
		Receives  struct {
			RoleName string
		}
		Returns struct {
			Error error
		}
	}
}

func (r *RolePolicies) Delete(roleName string) error {
	r.DeleteCall.CallCount++
	r.DeleteCall.Receives.RoleName = roleName

	return r.DeleteCall.Returns.Error
}
