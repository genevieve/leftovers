package fakes

import "github.com/aws/aws-sdk-go/service/iam"

type RolePoliciesClient struct {
	ListAttachedRolePoliciesCall struct {
		CallCount int
		Receives  struct {
			Input *iam.ListAttachedRolePoliciesInput
		}
		Returns struct {
			Output *iam.ListAttachedRolePoliciesOutput
			Error  error
		}
	}

	ListRolePoliciesCall struct {
		CallCount int
		Receives  struct {
			Input *iam.ListRolePoliciesInput
		}
		Returns struct {
			Output *iam.ListRolePoliciesOutput
			Error  error
		}
	}

	DetachRolePolicyCall struct {
		CallCount int
		Receives  struct {
			Input *iam.DetachRolePolicyInput
		}
		Returns struct {
			Output *iam.DetachRolePolicyOutput
			Error  error
		}
	}

	DeleteRolePolicyCall struct {
		CallCount int
		Receives  struct {
			Input *iam.DeleteRolePolicyInput
		}
		Returns struct {
			Output *iam.DeleteRolePolicyOutput
			Error  error
		}
	}
}

func (i *RolePoliciesClient) ListAttachedRolePolicies(input *iam.ListAttachedRolePoliciesInput) (*iam.ListAttachedRolePoliciesOutput, error) {
	i.ListAttachedRolePoliciesCall.CallCount++
	i.ListAttachedRolePoliciesCall.Receives.Input = input

	return i.ListAttachedRolePoliciesCall.Returns.Output, i.ListAttachedRolePoliciesCall.Returns.Error
}

func (i *RolePoliciesClient) ListRolePolicies(input *iam.ListRolePoliciesInput) (*iam.ListRolePoliciesOutput, error) {
	i.ListRolePoliciesCall.CallCount++
	i.ListRolePoliciesCall.Receives.Input = input

	return i.ListRolePoliciesCall.Returns.Output, i.ListRolePoliciesCall.Returns.Error
}

func (i *RolePoliciesClient) DetachRolePolicy(input *iam.DetachRolePolicyInput) (*iam.DetachRolePolicyOutput, error) {
	i.DetachRolePolicyCall.CallCount++
	i.DetachRolePolicyCall.Receives.Input = input

	return i.DetachRolePolicyCall.Returns.Output, i.DetachRolePolicyCall.Returns.Error
}

func (i *RolePoliciesClient) DeleteRolePolicy(input *iam.DeleteRolePolicyInput) (*iam.DeleteRolePolicyOutput, error) {
	i.DeleteRolePolicyCall.CallCount++
	i.DeleteRolePolicyCall.Receives.Input = input

	return i.DeleteRolePolicyCall.Returns.Output, i.DeleteRolePolicyCall.Returns.Error
}
