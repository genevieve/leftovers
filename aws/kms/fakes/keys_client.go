package fakes

import "github.com/aws/aws-sdk-go/service/kms"

type KeysClient struct {
	ListKeysCall struct {
		CallCount int
		Receives  struct {
			Input *kms.ListKeysInput
		}
		Returns struct {
			Output *kms.ListKeysOutput
			Error  error
		}
	}
	DescribeKeyCall struct {
		CallCount int
		Receives  struct {
			Input *kms.DescribeKeyInput
		}
		Returns struct {
			Output *kms.DescribeKeyOutput
			Error  error
		}
	}
	ListResourceTagsCall struct {
		CallCount int
		Receives  struct {
			Input *kms.ListResourceTagsInput
		}
		Returns struct {
			Output *kms.ListResourceTagsOutput
			Error  error
		}
	}
	DisableKeyCall struct {
		CallCount int
		Receives  struct {
			Input *kms.DisableKeyInput
		}
		Returns struct {
			Output *kms.DisableKeyOutput
			Error  error
		}
	}
	ScheduleKeyDeletionCall struct {
		CallCount int
		Receives  struct {
			Input *kms.ScheduleKeyDeletionInput
		}
		Returns struct {
			Output *kms.ScheduleKeyDeletionOutput
			Error  error
		}
	}
}

func (k *KeysClient) ListKeys(input *kms.ListKeysInput) (*kms.ListKeysOutput, error) {
	k.ListKeysCall.CallCount++
	k.ListKeysCall.Receives.Input = input

	return k.ListKeysCall.Returns.Output, k.ListKeysCall.Returns.Error
}

func (k *KeysClient) DescribeKey(input *kms.DescribeKeyInput) (*kms.DescribeKeyOutput, error) {
	k.DescribeKeyCall.CallCount++
	k.DescribeKeyCall.Receives.Input = input

	return k.DescribeKeyCall.Returns.Output, k.DescribeKeyCall.Returns.Error
}

func (k *KeysClient) ListResourceTags(input *kms.ListResourceTagsInput) (*kms.ListResourceTagsOutput, error) {
	k.ListResourceTagsCall.CallCount++
	k.ListResourceTagsCall.Receives.Input = input

	return k.ListResourceTagsCall.Returns.Output, k.ListResourceTagsCall.Returns.Error
}

func (k *KeysClient) DisableKey(input *kms.DisableKeyInput) (*kms.DisableKeyOutput, error) {
	k.DisableKeyCall.CallCount++
	k.DisableKeyCall.Receives.Input = input

	return k.DisableKeyCall.Returns.Output, k.DisableKeyCall.Returns.Error
}

func (k *KeysClient) ScheduleKeyDeletion(input *kms.ScheduleKeyDeletionInput) (*kms.ScheduleKeyDeletionOutput, error) {
	k.ScheduleKeyDeletionCall.CallCount++
	k.ScheduleKeyDeletionCall.Receives.Input = input

	return k.ScheduleKeyDeletionCall.Returns.Output, k.ScheduleKeyDeletionCall.Returns.Error
}
