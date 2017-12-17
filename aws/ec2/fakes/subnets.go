package fakes

type Subnets struct {
	DeleteCall struct {
		CallCount int
		Receives  struct {
			VpcId string
		}
		Returns struct {
			Error error
		}
	}
}

func (s *Subnets) Delete(vpcId string) error {
	s.DeleteCall.CallCount++
	s.DeleteCall.Receives.VpcId = vpcId

	return s.DeleteCall.Returns.Error
}
