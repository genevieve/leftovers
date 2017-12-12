package fakes

import "github.com/aws/aws-sdk-go/service/iam"

type UserPoliciesClient struct {
	ListUserPoliciesCall struct {
		CallCount int
		Receives  struct {
			Input *iam.ListUserPoliciesInput
		}
		Returns struct {
			Output *iam.ListUserPoliciesOutput
			Error  error
		}
	}

	ListPoliciesCall struct {
		CallCount int
		Receives  struct {
			Input *iam.ListPoliciesInput
		}
		Returns struct {
			Output *iam.ListPoliciesOutput
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

func (i *UserPoliciesClient) ListUserPolicies(input *iam.ListUserPoliciesInput) (*iam.ListUserPoliciesOutput, error) {
	i.ListUserPoliciesCall.CallCount++
	i.ListUserPoliciesCall.Receives.Input = input

	return i.ListUserPoliciesCall.Returns.Output, i.ListUserPoliciesCall.Returns.Error
}

func (i *UserPoliciesClient) ListPolicies(input *iam.ListPoliciesInput) (*iam.ListPoliciesOutput, error) {
	i.ListPoliciesCall.CallCount++
	i.ListPoliciesCall.Receives.Input = input

	return i.ListPoliciesCall.Returns.Output, i.ListPoliciesCall.Returns.Error
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
