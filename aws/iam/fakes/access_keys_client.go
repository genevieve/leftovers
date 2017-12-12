package fakes

import "github.com/aws/aws-sdk-go/service/iam"

type AccessKeysClient struct {
	ListAccessKeysCall struct {
		CallCount int
		Receives  struct {
			Input *iam.ListAccessKeysInput
		}
		Returns struct {
			Output *iam.ListAccessKeysOutput
			Error  error
		}
	}

	DeleteAccessKeyCall struct {
		CallCount int
		Receives  struct {
			Input *iam.DeleteAccessKeyInput
		}
		Returns struct {
			Output *iam.DeleteAccessKeyOutput
			Error  error
		}
	}
}

func (i *AccessKeysClient) ListAccessKeys(input *iam.ListAccessKeysInput) (*iam.ListAccessKeysOutput, error) {
	i.ListAccessKeysCall.CallCount++
	i.ListAccessKeysCall.Receives.Input = input

	return i.ListAccessKeysCall.Returns.Output, i.ListAccessKeysCall.Returns.Error
}

func (i *AccessKeysClient) DeleteAccessKey(input *iam.DeleteAccessKeyInput) (*iam.DeleteAccessKeyOutput, error) {
	i.DeleteAccessKeyCall.CallCount++
	i.DeleteAccessKeyCall.Receives.Input = input

	return i.DeleteAccessKeyCall.Returns.Output, i.DeleteAccessKeyCall.Returns.Error
}
