package fakes

import (
	"sync"

	awsiam "github.com/aws/aws-sdk-go/service/iam"
)

type UsersClient struct {
	DeleteUserCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DeleteUserInput *awsiam.DeleteUserInput
		}
		Returns struct {
			DeleteUserOutput *awsiam.DeleteUserOutput
			Error            error
		}
		Stub func(*awsiam.DeleteUserInput) (*awsiam.DeleteUserOutput, error)
	}
	ListUsersCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			ListUsersInput *awsiam.ListUsersInput
		}
		Returns struct {
			ListUsersOutput *awsiam.ListUsersOutput
			Error           error
		}
		Stub func(*awsiam.ListUsersInput) (*awsiam.ListUsersOutput, error)
	}
}

func (f *UsersClient) DeleteUser(param1 *awsiam.DeleteUserInput) (*awsiam.DeleteUserOutput, error) {
	f.DeleteUserCall.Lock()
	defer f.DeleteUserCall.Unlock()
	f.DeleteUserCall.CallCount++
	f.DeleteUserCall.Receives.DeleteUserInput = param1
	if f.DeleteUserCall.Stub != nil {
		return f.DeleteUserCall.Stub(param1)
	}
	return f.DeleteUserCall.Returns.DeleteUserOutput, f.DeleteUserCall.Returns.Error
}
func (f *UsersClient) ListUsers(param1 *awsiam.ListUsersInput) (*awsiam.ListUsersOutput, error) {
	f.ListUsersCall.Lock()
	defer f.ListUsersCall.Unlock()
	f.ListUsersCall.CallCount++
	f.ListUsersCall.Receives.ListUsersInput = param1
	if f.ListUsersCall.Stub != nil {
		return f.ListUsersCall.Stub(param1)
	}
	return f.ListUsersCall.Returns.ListUsersOutput, f.ListUsersCall.Returns.Error
}
