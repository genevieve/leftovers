package fakes

import "github.com/aws/aws-sdk-go/service/iam"

type InstanceProfilesClient struct {
	ListInstanceProfilesCall struct {
		CallCount int
		Receives  struct {
			Input *iam.ListInstanceProfilesInput
		}
		Returns struct {
			Output *iam.ListInstanceProfilesOutput
			Error  error
		}
	}

	DeleteInstanceProfileCall struct {
		CallCount int
		Receives  struct {
			Input *iam.DeleteInstanceProfileInput
		}
		Returns struct {
			Output *iam.DeleteInstanceProfileOutput
			Error  error
		}
	}

	RemoveRoleFromInstanceProfileCall struct {
		CallCount int
		Receives  struct {
			Input *iam.RemoveRoleFromInstanceProfileInput
		}
		Returns struct {
			Output *iam.RemoveRoleFromInstanceProfileOutput
			Error  error
		}
	}
}

func (i *InstanceProfilesClient) ListInstanceProfiles(input *iam.ListInstanceProfilesInput) (*iam.ListInstanceProfilesOutput, error) {
	i.ListInstanceProfilesCall.CallCount++
	i.ListInstanceProfilesCall.Receives.Input = input

	return i.ListInstanceProfilesCall.Returns.Output, i.ListInstanceProfilesCall.Returns.Error
}

func (i *InstanceProfilesClient) DeleteInstanceProfile(input *iam.DeleteInstanceProfileInput) (*iam.DeleteInstanceProfileOutput, error) {
	i.DeleteInstanceProfileCall.CallCount++
	i.DeleteInstanceProfileCall.Receives.Input = input

	return i.DeleteInstanceProfileCall.Returns.Output, i.DeleteInstanceProfileCall.Returns.Error
}

func (i *InstanceProfilesClient) RemoveRoleFromInstanceProfile(input *iam.RemoveRoleFromInstanceProfileInput) (*iam.RemoveRoleFromInstanceProfileOutput, error) {
	i.RemoveRoleFromInstanceProfileCall.CallCount++
	i.RemoveRoleFromInstanceProfileCall.Receives.Input = input

	return i.RemoveRoleFromInstanceProfileCall.Returns.Output, i.RemoveRoleFromInstanceProfileCall.Returns.Error
}
