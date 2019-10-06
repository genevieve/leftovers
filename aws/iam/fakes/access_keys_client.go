package fakes

import (
	"sync"

	awsiam "github.com/aws/aws-sdk-go/service/iam"
)

type AccessKeysClient struct {
	DeleteAccessKeyCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DeleteAccessKeyInput *awsiam.DeleteAccessKeyInput
		}
		Returns struct {
			DeleteAccessKeyOutput *awsiam.DeleteAccessKeyOutput
			Error                 error
		}
		Stub func(*awsiam.DeleteAccessKeyInput) (*awsiam.DeleteAccessKeyOutput, error)
	}
	ListAccessKeysCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			ListAccessKeysInput *awsiam.ListAccessKeysInput
		}
		Returns struct {
			ListAccessKeysOutput *awsiam.ListAccessKeysOutput
			Error                error
		}
		Stub func(*awsiam.ListAccessKeysInput) (*awsiam.ListAccessKeysOutput, error)
	}
}

func (f *AccessKeysClient) DeleteAccessKey(param1 *awsiam.DeleteAccessKeyInput) (*awsiam.DeleteAccessKeyOutput, error) {
	f.DeleteAccessKeyCall.Lock()
	defer f.DeleteAccessKeyCall.Unlock()
	f.DeleteAccessKeyCall.CallCount++
	f.DeleteAccessKeyCall.Receives.DeleteAccessKeyInput = param1
	if f.DeleteAccessKeyCall.Stub != nil {
		return f.DeleteAccessKeyCall.Stub(param1)
	}
	return f.DeleteAccessKeyCall.Returns.DeleteAccessKeyOutput, f.DeleteAccessKeyCall.Returns.Error
}
func (f *AccessKeysClient) ListAccessKeys(param1 *awsiam.ListAccessKeysInput) (*awsiam.ListAccessKeysOutput, error) {
	f.ListAccessKeysCall.Lock()
	defer f.ListAccessKeysCall.Unlock()
	f.ListAccessKeysCall.CallCount++
	f.ListAccessKeysCall.Receives.ListAccessKeysInput = param1
	if f.ListAccessKeysCall.Stub != nil {
		return f.ListAccessKeysCall.Stub(param1)
	}
	return f.ListAccessKeysCall.Returns.ListAccessKeysOutput, f.ListAccessKeysCall.Returns.Error
}
