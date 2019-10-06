package fakes

import (
	"sync"

	awskms "github.com/aws/aws-sdk-go/service/kms"
)

type KeysClient struct {
	DescribeKeyCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DescribeKeyInput *awskms.DescribeKeyInput
		}
		Returns struct {
			DescribeKeyOutput *awskms.DescribeKeyOutput
			Error             error
		}
		Stub func(*awskms.DescribeKeyInput) (*awskms.DescribeKeyOutput, error)
	}
	DisableKeyCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DisableKeyInput *awskms.DisableKeyInput
		}
		Returns struct {
			DisableKeyOutput *awskms.DisableKeyOutput
			Error            error
		}
		Stub func(*awskms.DisableKeyInput) (*awskms.DisableKeyOutput, error)
	}
	ListKeysCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			ListKeysInput *awskms.ListKeysInput
		}
		Returns struct {
			ListKeysOutput *awskms.ListKeysOutput
			Error          error
		}
		Stub func(*awskms.ListKeysInput) (*awskms.ListKeysOutput, error)
	}
	ListResourceTagsCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			ListResourceTagsInput *awskms.ListResourceTagsInput
		}
		Returns struct {
			ListResourceTagsOutput *awskms.ListResourceTagsOutput
			Error                  error
		}
		Stub func(*awskms.ListResourceTagsInput) (*awskms.ListResourceTagsOutput, error)
	}
	ScheduleKeyDeletionCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			ScheduleKeyDeletionInput *awskms.ScheduleKeyDeletionInput
		}
		Returns struct {
			ScheduleKeyDeletionOutput *awskms.ScheduleKeyDeletionOutput
			Error                     error
		}
		Stub func(*awskms.ScheduleKeyDeletionInput) (*awskms.ScheduleKeyDeletionOutput, error)
	}
}

func (f *KeysClient) DescribeKey(param1 *awskms.DescribeKeyInput) (*awskms.DescribeKeyOutput, error) {
	f.DescribeKeyCall.Lock()
	defer f.DescribeKeyCall.Unlock()
	f.DescribeKeyCall.CallCount++
	f.DescribeKeyCall.Receives.DescribeKeyInput = param1
	if f.DescribeKeyCall.Stub != nil {
		return f.DescribeKeyCall.Stub(param1)
	}
	return f.DescribeKeyCall.Returns.DescribeKeyOutput, f.DescribeKeyCall.Returns.Error
}
func (f *KeysClient) DisableKey(param1 *awskms.DisableKeyInput) (*awskms.DisableKeyOutput, error) {
	f.DisableKeyCall.Lock()
	defer f.DisableKeyCall.Unlock()
	f.DisableKeyCall.CallCount++
	f.DisableKeyCall.Receives.DisableKeyInput = param1
	if f.DisableKeyCall.Stub != nil {
		return f.DisableKeyCall.Stub(param1)
	}
	return f.DisableKeyCall.Returns.DisableKeyOutput, f.DisableKeyCall.Returns.Error
}
func (f *KeysClient) ListKeys(param1 *awskms.ListKeysInput) (*awskms.ListKeysOutput, error) {
	f.ListKeysCall.Lock()
	defer f.ListKeysCall.Unlock()
	f.ListKeysCall.CallCount++
	f.ListKeysCall.Receives.ListKeysInput = param1
	if f.ListKeysCall.Stub != nil {
		return f.ListKeysCall.Stub(param1)
	}
	return f.ListKeysCall.Returns.ListKeysOutput, f.ListKeysCall.Returns.Error
}
func (f *KeysClient) ListResourceTags(param1 *awskms.ListResourceTagsInput) (*awskms.ListResourceTagsOutput, error) {
	f.ListResourceTagsCall.Lock()
	defer f.ListResourceTagsCall.Unlock()
	f.ListResourceTagsCall.CallCount++
	f.ListResourceTagsCall.Receives.ListResourceTagsInput = param1
	if f.ListResourceTagsCall.Stub != nil {
		return f.ListResourceTagsCall.Stub(param1)
	}
	return f.ListResourceTagsCall.Returns.ListResourceTagsOutput, f.ListResourceTagsCall.Returns.Error
}
func (f *KeysClient) ScheduleKeyDeletion(param1 *awskms.ScheduleKeyDeletionInput) (*awskms.ScheduleKeyDeletionOutput, error) {
	f.ScheduleKeyDeletionCall.Lock()
	defer f.ScheduleKeyDeletionCall.Unlock()
	f.ScheduleKeyDeletionCall.CallCount++
	f.ScheduleKeyDeletionCall.Receives.ScheduleKeyDeletionInput = param1
	if f.ScheduleKeyDeletionCall.Stub != nil {
		return f.ScheduleKeyDeletionCall.Stub(param1)
	}
	return f.ScheduleKeyDeletionCall.Returns.ScheduleKeyDeletionOutput, f.ScheduleKeyDeletionCall.Returns.Error
}
