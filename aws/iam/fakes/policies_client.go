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

	ListPolicyVersionsCall struct {
		CallCount int
		Receives  struct {
			Input *awsiam.ListPolicyVersionsInput
		}
		Returns struct {
			Output *awsiam.ListPolicyVersionsOutput
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

	DeletePolicyVersionCall struct {
		CallCount int
		Receives  struct {
			Input *awsiam.DeletePolicyVersionInput
		}
		Returns struct {
			Output *awsiam.DeletePolicyVersionOutput
			Error  error
		}
	}
}

func (i *PoliciesClient) ListPolicies(input *awsiam.ListPoliciesInput) (*awsiam.ListPoliciesOutput, error) {
	i.ListPoliciesCall.CallCount++
	i.ListPoliciesCall.Receives.Input = input

	return i.ListPoliciesCall.Returns.Output, i.ListPoliciesCall.Returns.Error
}

func (i *PoliciesClient) ListPolicyVersions(input *awsiam.ListPolicyVersionsInput) (*awsiam.ListPolicyVersionsOutput, error) {
	i.ListPolicyVersionsCall.CallCount++
	i.ListPolicyVersionsCall.Receives.Input = input

	return i.ListPolicyVersionsCall.Returns.Output, i.ListPolicyVersionsCall.Returns.Error
}

func (i *PoliciesClient) DeletePolicy(input *awsiam.DeletePolicyInput) (*awsiam.DeletePolicyOutput, error) {
	i.DeletePolicyCall.CallCount++
	i.DeletePolicyCall.Receives.Input = input

	return i.DeletePolicyCall.Returns.Output, i.DeletePolicyCall.Returns.Error
}

func (i *PoliciesClient) DeletePolicyVersion(input *awsiam.DeletePolicyVersionInput) (*awsiam.DeletePolicyVersionOutput, error) {
	i.DeletePolicyVersionCall.CallCount++
	i.DeletePolicyVersionCall.Receives.Input = input

	return i.DeletePolicyVersionCall.Returns.Output, i.DeletePolicyVersionCall.Returns.Error
}
