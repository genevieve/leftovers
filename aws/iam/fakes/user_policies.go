package fakes

type UserPolicies struct {
	DeleteCall struct {
		CallCount int
		Receives  struct {
			UserName string
		}
		Returns struct {
			Error error
		}
	}
}

func (u *UserPolicies) Delete(userName string) error {
	u.DeleteCall.CallCount++
	u.DeleteCall.Receives.UserName = userName

	return u.DeleteCall.Returns.Error
}
