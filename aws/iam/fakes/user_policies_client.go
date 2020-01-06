package fakes

import (
	"sync"

	awsiam "github.com/aws/aws-sdk-go/service/iam"
)

type UserPoliciesClient struct {
	DeleteUserPolicyCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DeleteUserPolicyInput *awsiam.DeleteUserPolicyInput
		}
		Returns struct {
			DeleteUserPolicyOutput *awsiam.DeleteUserPolicyOutput
			Error                  error
		}
		Stub func(*awsiam.DeleteUserPolicyInput) (*awsiam.DeleteUserPolicyOutput, error)
	}
	DetachUserPolicyCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DetachUserPolicyInput *awsiam.DetachUserPolicyInput
		}
		Returns struct {
			DetachUserPolicyOutput *awsiam.DetachUserPolicyOutput
			Error                  error
		}
		Stub func(*awsiam.DetachUserPolicyInput) (*awsiam.DetachUserPolicyOutput, error)
	}
	ListAttachedUserPoliciesCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			ListAttachedUserPoliciesInput *awsiam.ListAttachedUserPoliciesInput
		}
		Returns struct {
			ListAttachedUserPoliciesOutput *awsiam.ListAttachedUserPoliciesOutput
			Error                          error
		}
		Stub func(*awsiam.ListAttachedUserPoliciesInput) (*awsiam.ListAttachedUserPoliciesOutput, error)
	}
	ListUserPoliciesCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			ListUserPoliciesInput *awsiam.ListUserPoliciesInput
		}
		Returns struct {
			ListUserPoliciesOutput *awsiam.ListUserPoliciesOutput
			Error                  error
		}
		Stub func(*awsiam.ListUserPoliciesInput) (*awsiam.ListUserPoliciesOutput, error)
	}
}

func (f *UserPoliciesClient) DeleteUserPolicy(param1 *awsiam.DeleteUserPolicyInput) (*awsiam.DeleteUserPolicyOutput, error) {
	f.DeleteUserPolicyCall.Lock()
	defer f.DeleteUserPolicyCall.Unlock()
	f.DeleteUserPolicyCall.CallCount++
	f.DeleteUserPolicyCall.Receives.DeleteUserPolicyInput = param1
	if f.DeleteUserPolicyCall.Stub != nil {
		return f.DeleteUserPolicyCall.Stub(param1)
	}
	return f.DeleteUserPolicyCall.Returns.DeleteUserPolicyOutput, f.DeleteUserPolicyCall.Returns.Error
}
func (f *UserPoliciesClient) DetachUserPolicy(param1 *awsiam.DetachUserPolicyInput) (*awsiam.DetachUserPolicyOutput, error) {
	f.DetachUserPolicyCall.Lock()
	defer f.DetachUserPolicyCall.Unlock()
	f.DetachUserPolicyCall.CallCount++
	f.DetachUserPolicyCall.Receives.DetachUserPolicyInput = param1
	if f.DetachUserPolicyCall.Stub != nil {
		return f.DetachUserPolicyCall.Stub(param1)
	}
	return f.DetachUserPolicyCall.Returns.DetachUserPolicyOutput, f.DetachUserPolicyCall.Returns.Error
}
func (f *UserPoliciesClient) ListAttachedUserPolicies(param1 *awsiam.ListAttachedUserPoliciesInput) (*awsiam.ListAttachedUserPoliciesOutput, error) {
	f.ListAttachedUserPoliciesCall.Lock()
	defer f.ListAttachedUserPoliciesCall.Unlock()
	f.ListAttachedUserPoliciesCall.CallCount++
	f.ListAttachedUserPoliciesCall.Receives.ListAttachedUserPoliciesInput = param1
	if f.ListAttachedUserPoliciesCall.Stub != nil {
		return f.ListAttachedUserPoliciesCall.Stub(param1)
	}
	return f.ListAttachedUserPoliciesCall.Returns.ListAttachedUserPoliciesOutput, f.ListAttachedUserPoliciesCall.Returns.Error
}
func (f *UserPoliciesClient) ListUserPolicies(param1 *awsiam.ListUserPoliciesInput) (*awsiam.ListUserPoliciesOutput, error) {
	f.ListUserPoliciesCall.Lock()
	defer f.ListUserPoliciesCall.Unlock()
	f.ListUserPoliciesCall.CallCount++
	f.ListUserPoliciesCall.Receives.ListUserPoliciesInput = param1
	if f.ListUserPoliciesCall.Stub != nil {
		return f.ListUserPoliciesCall.Stub(param1)
	}
	return f.ListUserPoliciesCall.Returns.ListUserPoliciesOutput, f.ListUserPoliciesCall.Returns.Error
}
