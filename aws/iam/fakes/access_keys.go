package fakes

type AccessKeys struct {
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

func (u *AccessKeys) Delete(userName string) error {
	u.DeleteCall.CallCount++
	u.DeleteCall.Receives.UserName = userName

	return u.DeleteCall.Returns.Error
}
