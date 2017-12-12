package fakes

import "github.com/aws/aws-sdk-go/service/iam"

type UserPoliciesClient struct {
	ListAttachedUserPoliciesCall struct {
		CallCount int
		Receives  struct {
			Input *iam.ListAttachedUserPoliciesInput
		}
		Returns struct {
			Output *iam.ListAttachedUserPoliciesOutput
			Error  error
		}
	}

	DetachUserPolicyCall struct {
		CallCount int
		Receives  struct {
			Input *iam.DetachUserPolicyInput
		}
		Returns struct {
			Output *iam.DetachUserPolicyOutput
			Error  error
		}
	}

	DeleteUserPolicyCall struct {
		CallCount int
		Receives  struct {
			Input *iam.DeleteUserPolicyInput
		}
		Returns struct {
			Output *iam.DeleteUserPolicyOutput
			Error  error
		}
	}
}

func (i *UserPoliciesClient) ListAttachedUserPolicies(input *iam.ListAttachedUserPoliciesInput) (*iam.ListAttachedUserPoliciesOutput, error) {
	i.ListAttachedUserPoliciesCall.CallCount++
	i.ListAttachedUserPoliciesCall.Receives.Input = input

	return i.ListAttachedUserPoliciesCall.Returns.Output, i.ListAttachedUserPoliciesCall.Returns.Error
}

func (i *UserPoliciesClient) DetachUserPolicy(input *iam.DetachUserPolicyInput) (*iam.DetachUserPolicyOutput, error) {
	i.DetachUserPolicyCall.CallCount++
	i.DetachUserPolicyCall.Receives.Input = input

	return i.DetachUserPolicyCall.Returns.Output, i.DetachUserPolicyCall.Returns.Error
}

func (i *UserPoliciesClient) DeleteUserPolicy(input *iam.DeleteUserPolicyInput) (*iam.DeleteUserPolicyOutput, error) {
	i.DeleteUserPolicyCall.CallCount++
	i.DeleteUserPolicyCall.Receives.Input = input

	return i.DeleteUserPolicyCall.Returns.Output, i.DeleteUserPolicyCall.Returns.Error
}
