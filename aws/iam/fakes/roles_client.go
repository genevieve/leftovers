package fakes

import "github.com/aws/aws-sdk-go/service/iam"

type RolesClient struct {
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
}

func (i *RolesClient) ListRoles(input *iam.ListRolesInput) (*iam.ListRolesOutput, error) {
	i.ListRolesCall.CallCount++
	i.ListRolesCall.Receives.Input = input

	return i.ListRolesCall.Returns.Output, i.ListRolesCall.Returns.Error
}

func (i *RolesClient) DeleteRole(input *iam.DeleteRoleInput) (*iam.DeleteRoleOutput, error) {
	i.DeleteRoleCall.CallCount++
	i.DeleteRoleCall.Receives.Input = input

	return i.DeleteRoleCall.Returns.Output, i.DeleteRoleCall.Returns.Error
}
