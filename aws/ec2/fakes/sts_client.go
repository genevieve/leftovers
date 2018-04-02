package fakes

import "github.com/aws/aws-sdk-go/service/sts"

type StsClient struct {
	GetCallerIdentityCall struct {
		CallCount int
		Receives  struct {
			Input *sts.GetCallerIdentityInput
		}
		Returns struct {
			Output *sts.GetCallerIdentityOutput
			Error  error
		}
	}
}

func (s *StsClient) GetCallerIdentity(input *sts.GetCallerIdentityInput) (*sts.GetCallerIdentityOutput, error) {
	s.GetCallerIdentityCall.CallCount++
	s.GetCallerIdentityCall.Receives.Input = input

	return s.GetCallerIdentityCall.Returns.Output, s.GetCallerIdentityCall.Returns.Error
}
