package fakes

import "github.com/aws/aws-sdk-go/service/kms"

type AliasesClient struct {
	ListAliasesCall struct {
		CallCount int
		Receives  struct {
			Input *kms.ListAliasesInput
		}
		Returns struct {
			Output *kms.ListAliasesOutput
			Error  error
		}
	}

	DeleteAliasCall struct {
		CallCount int
		Receives  struct {
			Input *kms.DeleteAliasInput
		}
		Returns struct {
			Output *kms.DeleteAliasOutput
			Error  error
		}
	}
}

func (a *AliasesClient) ListAliases(input *kms.ListAliasesInput) (*kms.ListAliasesOutput, error) {
	a.ListAliasesCall.CallCount++
	a.ListAliasesCall.Receives.Input = input

	return a.ListAliasesCall.Returns.Output, a.ListAliasesCall.Returns.Error
}

func (a *AliasesClient) DeleteAlias(input *kms.DeleteAliasInput) (*kms.DeleteAliasOutput, error) {
	a.DeleteAliasCall.CallCount++
	a.DeleteAliasCall.Receives.Input = input

	return a.DeleteAliasCall.Returns.Output, a.DeleteAliasCall.Returns.Error
}
