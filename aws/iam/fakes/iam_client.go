package fakes

import "github.com/aws/aws-sdk-go/service/iam"

type IAMClient struct {
	ListServerCertificatesCall struct {
		CallCount int
		Receives  struct {
			Input *iam.ListServerCertificatesInput
		}
		Returns struct {
			Output *iam.ListServerCertificatesOutput
			Error  error
		}
	}

	DeleteServerCertificateCall struct {
		CallCount int
		Receives  struct {
			Input *iam.DeleteServerCertificateInput
		}
		Returns struct {
			Output *iam.DeleteServerCertificateOutput
			Error  error
		}
	}

	ListRolesCall struct {
		CallCount int
		Receives  struct {
			Input *iam.ListRolesInput
		}
		Returns struct {
			Output *iam.ListRolesOutput
			Error  error
		}
	}

	DeleteRoleCall struct {
		CallCount int
		Receives  struct {
			Input *iam.DeleteRoleInput
		}
		Returns struct {
			Output *iam.DeleteRoleOutput
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

func (i *IAMClient) ListServerCertificates(input *iam.ListServerCertificatesInput) (*iam.ListServerCertificatesOutput, error) {
	i.ListServerCertificatesCall.CallCount++
	i.ListServerCertificatesCall.Receives.Input = input

	return i.ListServerCertificatesCall.Returns.Output, i.ListServerCertificatesCall.Returns.Error
}

func (i *IAMClient) DeleteServerCertificate(input *iam.DeleteServerCertificateInput) (*iam.DeleteServerCertificateOutput, error) {
	i.DeleteServerCertificateCall.CallCount++
	i.DeleteServerCertificateCall.Receives.Input = input

	return i.DeleteServerCertificateCall.Returns.Output, i.DeleteServerCertificateCall.Returns.Error
}

func (i *IAMClient) ListRoles(input *iam.ListRolesInput) (*iam.ListRolesOutput, error) {
	i.ListRolesCall.CallCount++
	i.ListRolesCall.Receives.Input = input

	return i.ListRolesCall.Returns.Output, i.ListRolesCall.Returns.Error
}

func (i *IAMClient) DeleteRole(input *iam.DeleteRoleInput) (*iam.DeleteRoleOutput, error) {
	i.DeleteRoleCall.CallCount++
	i.DeleteRoleCall.Receives.Input = input

	return i.DeleteRoleCall.Returns.Output, i.DeleteRoleCall.Returns.Error
}

func (i *IAMClient) ListRolePolicies(input *iam.ListRolePoliciesInput) (*iam.ListRolePoliciesOutput, error) {
	i.ListRolePoliciesCall.CallCount++
	i.ListRolePoliciesCall.Receives.Input = input

	return i.ListRolePoliciesCall.Returns.Output, i.ListRolePoliciesCall.Returns.Error
}

func (i *IAMClient) ListPolicies(input *iam.ListPoliciesInput) (*iam.ListPoliciesOutput, error) {
	i.ListPoliciesCall.CallCount++
	i.ListPoliciesCall.Receives.Input = input

	return i.ListPoliciesCall.Returns.Output, i.ListPoliciesCall.Returns.Error
}

func (i *IAMClient) DetachRolePolicy(input *iam.DetachRolePolicyInput) (*iam.DetachRolePolicyOutput, error) {
	i.DetachRolePolicyCall.CallCount++
	i.DetachRolePolicyCall.Receives.Input = input

	return i.DetachRolePolicyCall.Returns.Output, i.DetachRolePolicyCall.Returns.Error
}

func (i *IAMClient) DeleteRolePolicy(input *iam.DeleteRolePolicyInput) (*iam.DeleteRolePolicyOutput, error) {
	i.DeleteRolePolicyCall.CallCount++
	i.DeleteRolePolicyCall.Receives.Input = input

	return i.DeleteRolePolicyCall.Returns.Output, i.DeleteRolePolicyCall.Returns.Error
}
