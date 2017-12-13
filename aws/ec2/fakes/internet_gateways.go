package fakes

type InternetGateways struct {
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

func (r *InternetGateways) Delete(vpcId string) error {
	r.DeleteCall.CallCount++
	r.DeleteCall.Receives.VpcId = vpcId

	return r.DeleteCall.Returns.Error
}
