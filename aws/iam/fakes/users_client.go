package fakes

import "github.com/aws/aws-sdk-go/service/iam"

type UsersClient struct {
	ListUsersCall struct {
		CallCount int
		Receives  struct {
			Input *iam.ListUsersInput
		}
		Returns struct {
			Output *iam.ListUsersOutput
			Error  error
		}
	}

	DeleteUserCall struct {
		CallCount int
		Receives  struct {
			Input *iam.DeleteUserInput
		}
		Returns struct {
			Output *iam.DeleteUserOutput
			Error  error
		}
	}
}

func (i *UsersClient) ListUsers(input *iam.ListUsersInput) (*iam.ListUsersOutput, error) {
	i.ListUsersCall.CallCount++
	i.ListUsersCall.Receives.Input = input

	return i.ListUsersCall.Returns.Output, i.ListUsersCall.Returns.Error
}

func (i *UsersClient) DeleteUser(input *iam.DeleteUserInput) (*iam.DeleteUserOutput, error) {
	i.DeleteUserCall.CallCount++
	i.DeleteUserCall.Receives.Input = input

	return i.DeleteUserCall.Returns.Output, i.DeleteUserCall.Returns.Error
}
