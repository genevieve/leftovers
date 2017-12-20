package fakes

import awsiam "github.com/aws/aws-sdk-go/service/iam"

type RolePoliciesClient struct {
	ListAttachedRolePoliciesCall struct {
		CallCount int
		Receives  struct {
			Input *awsiam.ListAttachedRolePoliciesInput
		}
		Returns struct {
			Output *awsiam.ListAttachedRolePoliciesOutput
			Error  error
		}
	}

	ListRolePoliciesCall struct {
		CallCount int
		Receives  struct {
			Input *awsiam.ListRolePoliciesInput
		}
		Returns struct {
			Output *awsiam.ListRolePoliciesOutput
			Error  error
		}
	}

	DetachRolePolicyCall struct {
		CallCount int
		Receives  struct {
			Input *awsiam.DetachRolePolicyInput
		}
		Returns struct {
			Output *awsiam.DetachRolePolicyOutput
			Error  error
		}
	}

	DeleteRolePolicyCall struct {
		CallCount int
		Receives  struct {
			Input *awsiam.DeleteRolePolicyInput
		}
		Returns struct {
			Output *awsiam.DeleteRolePolicyOutput
			Error  error
		}
	}
}

func (i *RolePoliciesClient) ListAttachedRolePolicies(input *awsiam.ListAttachedRolePoliciesInput) (*awsiam.ListAttachedRolePoliciesOutput, error) {
	i.ListAttachedRolePoliciesCall.CallCount++
	i.ListAttachedRolePoliciesCall.Receives.Input = input

	return i.ListAttachedRolePoliciesCall.Returns.Output, i.ListAttachedRolePoliciesCall.Returns.Error
}

func (i *RolePoliciesClient) ListRolePolicies(input *awsiam.ListRolePoliciesInput) (*awsiam.ListRolePoliciesOutput, error) {
	i.ListRolePoliciesCall.CallCount++
	i.ListRolePoliciesCall.Receives.Input = input

	return i.ListRolePoliciesCall.Returns.Output, i.ListRolePoliciesCall.Returns.Error
}

func (i *RolePoliciesClient) DetachRolePolicy(input *awsiam.DetachRolePolicyInput) (*awsiam.DetachRolePolicyOutput, error) {
	i.DetachRolePolicyCall.CallCount++
	i.DetachRolePolicyCall.Receives.Input = input

	return i.DetachRolePolicyCall.Returns.Output, i.DetachRolePolicyCall.Returns.Error
}

func (i *RolePoliciesClient) DeleteRolePolicy(input *awsiam.DeleteRolePolicyInput) (*awsiam.DeleteRolePolicyOutput, error) {
	i.DeleteRolePolicyCall.CallCount++
	i.DeleteRolePolicyCall.Receives.Input = input

	return i.DeleteRolePolicyCall.Returns.Output, i.DeleteRolePolicyCall.Returns.Error
}
