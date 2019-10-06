package fakes

import (
	"sync"

	awskms "github.com/aws/aws-sdk-go/service/kms"
)

type AliasesClient struct {
	DeleteAliasCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DeleteAliasInput *awskms.DeleteAliasInput
		}
		Returns struct {
			DeleteAliasOutput *awskms.DeleteAliasOutput
			Error             error
		}
		Stub func(*awskms.DeleteAliasInput) (*awskms.DeleteAliasOutput, error)
	}
	ListAliasesCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			ListAliasesInput *awskms.ListAliasesInput
		}
		Returns struct {
			ListAliasesOutput *awskms.ListAliasesOutput
			Error             error
		}
		Stub func(*awskms.ListAliasesInput) (*awskms.ListAliasesOutput, error)
	}
}

func (f *AliasesClient) DeleteAlias(param1 *awskms.DeleteAliasInput) (*awskms.DeleteAliasOutput, error) {
	f.DeleteAliasCall.Lock()
	defer f.DeleteAliasCall.Unlock()
	f.DeleteAliasCall.CallCount++
	f.DeleteAliasCall.Receives.DeleteAliasInput = param1
	if f.DeleteAliasCall.Stub != nil {
		return f.DeleteAliasCall.Stub(param1)
	}
	return f.DeleteAliasCall.Returns.DeleteAliasOutput, f.DeleteAliasCall.Returns.Error
}
func (f *AliasesClient) ListAliases(param1 *awskms.ListAliasesInput) (*awskms.ListAliasesOutput, error) {
	f.ListAliasesCall.Lock()
	defer f.ListAliasesCall.Unlock()
	f.ListAliasesCall.CallCount++
	f.ListAliasesCall.Receives.ListAliasesInput = param1
	if f.ListAliasesCall.Stub != nil {
		return f.ListAliasesCall.Stub(param1)
	}
	return f.ListAliasesCall.Returns.ListAliasesOutput, f.ListAliasesCall.Returns.Error
}
