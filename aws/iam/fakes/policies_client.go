package fakes

import (
	"sync"

	awsiam "github.com/aws/aws-sdk-go/service/iam"
)

type PoliciesClient struct {
	DeletePolicyCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DeletePolicyInput *awsiam.DeletePolicyInput
		}
		Returns struct {
			DeletePolicyOutput *awsiam.DeletePolicyOutput
			Error              error
		}
		Stub func(*awsiam.DeletePolicyInput) (*awsiam.DeletePolicyOutput, error)
	}
	DeletePolicyVersionCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DeletePolicyVersionInput *awsiam.DeletePolicyVersionInput
		}
		Returns struct {
			DeletePolicyVersionOutput *awsiam.DeletePolicyVersionOutput
			Error                     error
		}
		Stub func(*awsiam.DeletePolicyVersionInput) (*awsiam.DeletePolicyVersionOutput, error)
	}
	ListPoliciesCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			ListPoliciesInput *awsiam.ListPoliciesInput
		}
		Returns struct {
			ListPoliciesOutput *awsiam.ListPoliciesOutput
			Error              error
		}
		Stub func(*awsiam.ListPoliciesInput) (*awsiam.ListPoliciesOutput, error)
	}
	ListPolicyVersionsCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			ListPolicyVersionsInput *awsiam.ListPolicyVersionsInput
		}
		Returns struct {
			ListPolicyVersionsOutput *awsiam.ListPolicyVersionsOutput
			Error                    error
		}
		Stub func(*awsiam.ListPolicyVersionsInput) (*awsiam.ListPolicyVersionsOutput, error)
	}
}

func (f *PoliciesClient) DeletePolicy(param1 *awsiam.DeletePolicyInput) (*awsiam.DeletePolicyOutput, error) {
	f.DeletePolicyCall.Lock()
	defer f.DeletePolicyCall.Unlock()
	f.DeletePolicyCall.CallCount++
	f.DeletePolicyCall.Receives.DeletePolicyInput = param1
	if f.DeletePolicyCall.Stub != nil {
		return f.DeletePolicyCall.Stub(param1)
	}
	return f.DeletePolicyCall.Returns.DeletePolicyOutput, f.DeletePolicyCall.Returns.Error
}
func (f *PoliciesClient) DeletePolicyVersion(param1 *awsiam.DeletePolicyVersionInput) (*awsiam.DeletePolicyVersionOutput, error) {
	f.DeletePolicyVersionCall.Lock()
	defer f.DeletePolicyVersionCall.Unlock()
	f.DeletePolicyVersionCall.CallCount++
	f.DeletePolicyVersionCall.Receives.DeletePolicyVersionInput = param1
	if f.DeletePolicyVersionCall.Stub != nil {
		return f.DeletePolicyVersionCall.Stub(param1)
	}
	return f.DeletePolicyVersionCall.Returns.DeletePolicyVersionOutput, f.DeletePolicyVersionCall.Returns.Error
}
func (f *PoliciesClient) ListPolicies(param1 *awsiam.ListPoliciesInput) (*awsiam.ListPoliciesOutput, error) {
	f.ListPoliciesCall.Lock()
	defer f.ListPoliciesCall.Unlock()
	f.ListPoliciesCall.CallCount++
	f.ListPoliciesCall.Receives.ListPoliciesInput = param1
	if f.ListPoliciesCall.Stub != nil {
		return f.ListPoliciesCall.Stub(param1)
	}
	return f.ListPoliciesCall.Returns.ListPoliciesOutput, f.ListPoliciesCall.Returns.Error
}
func (f *PoliciesClient) ListPolicyVersions(param1 *awsiam.ListPolicyVersionsInput) (*awsiam.ListPolicyVersionsOutput, error) {
	f.ListPolicyVersionsCall.Lock()
	defer f.ListPolicyVersionsCall.Unlock()
	f.ListPolicyVersionsCall.CallCount++
	f.ListPolicyVersionsCall.Receives.ListPolicyVersionsInput = param1
	if f.ListPolicyVersionsCall.Stub != nil {
		return f.ListPolicyVersionsCall.Stub(param1)
	}
	return f.ListPolicyVersionsCall.Returns.ListPolicyVersionsOutput, f.ListPolicyVersionsCall.Returns.Error
}
