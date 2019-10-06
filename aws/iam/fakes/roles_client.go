package fakes

import (
	"sync"

	awsiam "github.com/aws/aws-sdk-go/service/iam"
)

type RolesClient struct {
	DeleteRoleCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DeleteRoleInput *awsiam.DeleteRoleInput
		}
		Returns struct {
			DeleteRoleOutput *awsiam.DeleteRoleOutput
			Error            error
		}
		Stub func(*awsiam.DeleteRoleInput) (*awsiam.DeleteRoleOutput, error)
	}
	ListRolesCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			ListRolesInput *awsiam.ListRolesInput
		}
		Returns struct {
			ListRolesOutput *awsiam.ListRolesOutput
			Error           error
		}
		Stub func(*awsiam.ListRolesInput) (*awsiam.ListRolesOutput, error)
	}
}

func (f *RolesClient) DeleteRole(param1 *awsiam.DeleteRoleInput) (*awsiam.DeleteRoleOutput, error) {
	f.DeleteRoleCall.Lock()
	defer f.DeleteRoleCall.Unlock()
	f.DeleteRoleCall.CallCount++
	f.DeleteRoleCall.Receives.DeleteRoleInput = param1
	if f.DeleteRoleCall.Stub != nil {
		return f.DeleteRoleCall.Stub(param1)
	}
	return f.DeleteRoleCall.Returns.DeleteRoleOutput, f.DeleteRoleCall.Returns.Error
}
func (f *RolesClient) ListRoles(param1 *awsiam.ListRolesInput) (*awsiam.ListRolesOutput, error) {
	f.ListRolesCall.Lock()
	defer f.ListRolesCall.Unlock()
	f.ListRolesCall.CallCount++
	f.ListRolesCall.Receives.ListRolesInput = param1
	if f.ListRolesCall.Stub != nil {
		return f.ListRolesCall.Stub(param1)
	}
	return f.ListRolesCall.Returns.ListRolesOutput, f.ListRolesCall.Returns.Error
}
