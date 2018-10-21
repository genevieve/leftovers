package fakes

import (
	gcpcrm "google.golang.org/api/cloudresourcemanager/v1"
	gcpiam "google.golang.org/api/iam/v1"
)

type ServiceAccountsClient struct {
	GetProjectIamPolicyCall struct {
		CallCount int
		Returns   struct {
			Output *gcpcrm.Policy
			Error  error
		}
	}

	SetProjectIamPolicyCall struct {
		CallCount int
		Receives  struct {
			Input *gcpcrm.Policy
		}
		Returns struct {
			Output *gcpcrm.Policy
			Error  error
		}
	}

	ListServiceAccountsCall struct {
		CallCount int
		Returns   struct {
			Output []*gcpiam.ServiceAccount
			Error  error
		}
	}

	DeleteServiceAccountCall struct {
		CallCount int
		Receives  struct {
			ServiceAccount string
		}
		Returns struct {
			Error error
		}
	}
}

func (s *ServiceAccountsClient) SetProjectIamPolicy(input *gcpcrm.Policy) (*gcpcrm.Policy, error) {
	s.SetProjectIamPolicyCall.CallCount++
	s.SetProjectIamPolicyCall.Receives.Input = input

	return s.SetProjectIamPolicyCall.Returns.Output, s.SetProjectIamPolicyCall.Returns.Error
}

func (s *ServiceAccountsClient) GetProjectIamPolicy() (*gcpcrm.Policy, error) {
	s.GetProjectIamPolicyCall.CallCount++

	return s.GetProjectIamPolicyCall.Returns.Output, s.GetProjectIamPolicyCall.Returns.Error
}

func (s *ServiceAccountsClient) ListServiceAccounts() ([]*gcpiam.ServiceAccount, error) {
	s.ListServiceAccountsCall.CallCount++

	return s.ListServiceAccountsCall.Returns.Output, s.ListServiceAccountsCall.Returns.Error
}

func (s *ServiceAccountsClient) DeleteServiceAccount(account string) error {
	s.DeleteServiceAccountCall.CallCount++
	s.DeleteServiceAccountCall.Receives.ServiceAccount = account

	return s.DeleteServiceAccountCall.Returns.Error
}
