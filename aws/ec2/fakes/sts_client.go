package fakes

import (
	"sync"

	awssts "github.com/aws/aws-sdk-go/service/sts"
)

type StsClient struct {
	GetCallerIdentityCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			GetCallerIdentityInput *awssts.GetCallerIdentityInput
		}
		Returns struct {
			GetCallerIdentityOutput *awssts.GetCallerIdentityOutput
			Error                   error
		}
		Stub func(*awssts.GetCallerIdentityInput) (*awssts.GetCallerIdentityOutput, error)
	}
}

func (f *StsClient) GetCallerIdentity(param1 *awssts.GetCallerIdentityInput) (*awssts.GetCallerIdentityOutput, error) {
	f.GetCallerIdentityCall.Lock()
	defer f.GetCallerIdentityCall.Unlock()
	f.GetCallerIdentityCall.CallCount++
	f.GetCallerIdentityCall.Receives.GetCallerIdentityInput = param1
	if f.GetCallerIdentityCall.Stub != nil {
		return f.GetCallerIdentityCall.Stub(param1)
	}
	return f.GetCallerIdentityCall.Returns.GetCallerIdentityOutput, f.GetCallerIdentityCall.Returns.Error
}
