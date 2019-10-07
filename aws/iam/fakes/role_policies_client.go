package fakes

import (
	"sync"

	awsiam "github.com/aws/aws-sdk-go/service/iam"
)

type RolePoliciesClient struct {
	DeleteRolePolicyCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DeleteRolePolicyInput *awsiam.DeleteRolePolicyInput
		}
		Returns struct {
			DeleteRolePolicyOutput *awsiam.DeleteRolePolicyOutput
			Error                  error
		}
		Stub func(*awsiam.DeleteRolePolicyInput) (*awsiam.DeleteRolePolicyOutput, error)
	}
	DetachRolePolicyCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DetachRolePolicyInput *awsiam.DetachRolePolicyInput
		}
		Returns struct {
			DetachRolePolicyOutput *awsiam.DetachRolePolicyOutput
			Error                  error
		}
		Stub func(*awsiam.DetachRolePolicyInput) (*awsiam.DetachRolePolicyOutput, error)
	}
	ListAttachedRolePoliciesCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			ListAttachedRolePoliciesInput *awsiam.ListAttachedRolePoliciesInput
		}
		Returns struct {
			ListAttachedRolePoliciesOutput *awsiam.ListAttachedRolePoliciesOutput
			Error                          error
		}
		Stub func(*awsiam.ListAttachedRolePoliciesInput) (*awsiam.ListAttachedRolePoliciesOutput, error)
	}
	ListRolePoliciesCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			ListRolePoliciesInput *awsiam.ListRolePoliciesInput
		}
		Returns struct {
			ListRolePoliciesOutput *awsiam.ListRolePoliciesOutput
			Error                  error
		}
		Stub func(*awsiam.ListRolePoliciesInput) (*awsiam.ListRolePoliciesOutput, error)
	}
}

func (f *RolePoliciesClient) DeleteRolePolicy(param1 *awsiam.DeleteRolePolicyInput) (*awsiam.DeleteRolePolicyOutput, error) {
	f.DeleteRolePolicyCall.Lock()
	defer f.DeleteRolePolicyCall.Unlock()
	f.DeleteRolePolicyCall.CallCount++
	f.DeleteRolePolicyCall.Receives.DeleteRolePolicyInput = param1
	if f.DeleteRolePolicyCall.Stub != nil {
		return f.DeleteRolePolicyCall.Stub(param1)
	}
	return f.DeleteRolePolicyCall.Returns.DeleteRolePolicyOutput, f.DeleteRolePolicyCall.Returns.Error
}
func (f *RolePoliciesClient) DetachRolePolicy(param1 *awsiam.DetachRolePolicyInput) (*awsiam.DetachRolePolicyOutput, error) {
	f.DetachRolePolicyCall.Lock()
	defer f.DetachRolePolicyCall.Unlock()
	f.DetachRolePolicyCall.CallCount++
	f.DetachRolePolicyCall.Receives.DetachRolePolicyInput = param1
	if f.DetachRolePolicyCall.Stub != nil {
		return f.DetachRolePolicyCall.Stub(param1)
	}
	return f.DetachRolePolicyCall.Returns.DetachRolePolicyOutput, f.DetachRolePolicyCall.Returns.Error
}
func (f *RolePoliciesClient) ListAttachedRolePolicies(param1 *awsiam.ListAttachedRolePoliciesInput) (*awsiam.ListAttachedRolePoliciesOutput, error) {
	f.ListAttachedRolePoliciesCall.Lock()
	defer f.ListAttachedRolePoliciesCall.Unlock()
	f.ListAttachedRolePoliciesCall.CallCount++
	f.ListAttachedRolePoliciesCall.Receives.ListAttachedRolePoliciesInput = param1
	if f.ListAttachedRolePoliciesCall.Stub != nil {
		return f.ListAttachedRolePoliciesCall.Stub(param1)
	}
	return f.ListAttachedRolePoliciesCall.Returns.ListAttachedRolePoliciesOutput, f.ListAttachedRolePoliciesCall.Returns.Error
}
func (f *RolePoliciesClient) ListRolePolicies(param1 *awsiam.ListRolePoliciesInput) (*awsiam.ListRolePoliciesOutput, error) {
	f.ListRolePoliciesCall.Lock()
	defer f.ListRolePoliciesCall.Unlock()
	f.ListRolePoliciesCall.CallCount++
	f.ListRolePoliciesCall.Receives.ListRolePoliciesInput = param1
	if f.ListRolePoliciesCall.Stub != nil {
		return f.ListRolePoliciesCall.Stub(param1)
	}
	return f.ListRolePoliciesCall.Returns.ListRolePoliciesOutput, f.ListRolePoliciesCall.Returns.Error
}
