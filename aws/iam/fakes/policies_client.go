package fakes

import awsiam "github.com/aws/aws-sdk-go/service/iam"

type PoliciesClient struct {
	ListPoliciesCall struct {
		CallCount int
		Receives  struct {
			Input *awsiam.ListPoliciesInput
		}
		Returns struct {
			Output *awsiam.ListPoliciesOutput
			Error  error
		}
	}

	DeletePolicyCall struct {
		CallCount int
		Receives  struct {
			Input *awsiam.DeletePolicyInput
		}
		Returns struct {
			Output *awsiam.DeletePolicyOutput
			Error  error
		}
	}
}

func (i *PoliciesClient) ListPolicies(input *awsiam.ListPoliciesInput) (*awsiam.ListPoliciesOutput, error) {
	i.ListPoliciesCall.CallCount++
	i.ListPoliciesCall.Receives.Input = input

	return i.ListPoliciesCall.Returns.Output, i.ListPoliciesCall.Returns.Error
}

func (i *PoliciesClient) DeletePolicy(input *awsiam.DeletePolicyInput) (*awsiam.DeletePolicyOutput, error) {
	i.DeletePolicyCall.CallCount++
	i.DeletePolicyCall.Receives.Input = input

	return i.DeletePolicyCall.Returns.Output, i.DeletePolicyCall.Returns.Error
}
