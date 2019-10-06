package fakes

import (
	"sync"

	awsiam "github.com/aws/aws-sdk-go/service/iam"
)

type InstanceProfilesClient struct {
	DeleteInstanceProfileCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DeleteInstanceProfileInput *awsiam.DeleteInstanceProfileInput
		}
		Returns struct {
			DeleteInstanceProfileOutput *awsiam.DeleteInstanceProfileOutput
			Error                       error
		}
		Stub func(*awsiam.DeleteInstanceProfileInput) (*awsiam.DeleteInstanceProfileOutput, error)
	}
	ListInstanceProfilesCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			ListInstanceProfilesInput *awsiam.ListInstanceProfilesInput
		}
		Returns struct {
			ListInstanceProfilesOutput *awsiam.ListInstanceProfilesOutput
			Error                      error
		}
		Stub func(*awsiam.ListInstanceProfilesInput) (*awsiam.ListInstanceProfilesOutput, error)
	}
	RemoveRoleFromInstanceProfileCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			RemoveRoleFromInstanceProfileInput *awsiam.RemoveRoleFromInstanceProfileInput
		}
		Returns struct {
			RemoveRoleFromInstanceProfileOutput *awsiam.RemoveRoleFromInstanceProfileOutput
			Error                               error
		}
		Stub func(*awsiam.RemoveRoleFromInstanceProfileInput) (*awsiam.RemoveRoleFromInstanceProfileOutput, error)
	}
}

func (f *InstanceProfilesClient) DeleteInstanceProfile(param1 *awsiam.DeleteInstanceProfileInput) (*awsiam.DeleteInstanceProfileOutput, error) {
	f.DeleteInstanceProfileCall.Lock()
	defer f.DeleteInstanceProfileCall.Unlock()
	f.DeleteInstanceProfileCall.CallCount++
	f.DeleteInstanceProfileCall.Receives.DeleteInstanceProfileInput = param1
	if f.DeleteInstanceProfileCall.Stub != nil {
		return f.DeleteInstanceProfileCall.Stub(param1)
	}
	return f.DeleteInstanceProfileCall.Returns.DeleteInstanceProfileOutput, f.DeleteInstanceProfileCall.Returns.Error
}
func (f *InstanceProfilesClient) ListInstanceProfiles(param1 *awsiam.ListInstanceProfilesInput) (*awsiam.ListInstanceProfilesOutput, error) {
	f.ListInstanceProfilesCall.Lock()
	defer f.ListInstanceProfilesCall.Unlock()
	f.ListInstanceProfilesCall.CallCount++
	f.ListInstanceProfilesCall.Receives.ListInstanceProfilesInput = param1
	if f.ListInstanceProfilesCall.Stub != nil {
		return f.ListInstanceProfilesCall.Stub(param1)
	}
	return f.ListInstanceProfilesCall.Returns.ListInstanceProfilesOutput, f.ListInstanceProfilesCall.Returns.Error
}
func (f *InstanceProfilesClient) RemoveRoleFromInstanceProfile(param1 *awsiam.RemoveRoleFromInstanceProfileInput) (*awsiam.RemoveRoleFromInstanceProfileOutput, error) {
	f.RemoveRoleFromInstanceProfileCall.Lock()
	defer f.RemoveRoleFromInstanceProfileCall.Unlock()
	f.RemoveRoleFromInstanceProfileCall.CallCount++
	f.RemoveRoleFromInstanceProfileCall.Receives.RemoveRoleFromInstanceProfileInput = param1
	if f.RemoveRoleFromInstanceProfileCall.Stub != nil {
		return f.RemoveRoleFromInstanceProfileCall.Stub(param1)
	}
	return f.RemoveRoleFromInstanceProfileCall.Returns.RemoveRoleFromInstanceProfileOutput, f.RemoveRoleFromInstanceProfileCall.Returns.Error
}
