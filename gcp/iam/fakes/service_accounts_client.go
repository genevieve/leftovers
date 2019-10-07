package fakes

import (
	"sync"

	gcpcrm "google.golang.org/api/cloudresourcemanager/v1"
	gcpiam "google.golang.org/api/iam/v1"
)

type ServiceAccountsClient struct {
	DeleteServiceAccountCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Account string
		}
		Returns struct {
			Error error
		}
		Stub func(string) error
	}
	GetProjectIamPolicyCall struct {
		sync.Mutex
		CallCount int
		Returns   struct {
			Policy *gcpcrm.Policy
			Error  error
		}
		Stub func() (*gcpcrm.Policy, error)
	}
	ListServiceAccountsCall struct {
		sync.Mutex
		CallCount int
		Returns   struct {
			ServiceAccountSlice []*gcpiam.ServiceAccount
			Error               error
		}
		Stub func() ([]*gcpiam.ServiceAccount, error)
	}
	SetProjectIamPolicyCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Policy *gcpcrm.Policy
		}
		Returns struct {
			Policy *gcpcrm.Policy
			Error  error
		}
		Stub func(*gcpcrm.Policy) (*gcpcrm.Policy, error)
	}
}

func (f *ServiceAccountsClient) DeleteServiceAccount(param1 string) error {
	f.DeleteServiceAccountCall.Lock()
	defer f.DeleteServiceAccountCall.Unlock()
	f.DeleteServiceAccountCall.CallCount++
	f.DeleteServiceAccountCall.Receives.Account = param1
	if f.DeleteServiceAccountCall.Stub != nil {
		return f.DeleteServiceAccountCall.Stub(param1)
	}
	return f.DeleteServiceAccountCall.Returns.Error
}
func (f *ServiceAccountsClient) GetProjectIamPolicy() (*gcpcrm.Policy, error) {
	f.GetProjectIamPolicyCall.Lock()
	defer f.GetProjectIamPolicyCall.Unlock()
	f.GetProjectIamPolicyCall.CallCount++
	if f.GetProjectIamPolicyCall.Stub != nil {
		return f.GetProjectIamPolicyCall.Stub()
	}
	return f.GetProjectIamPolicyCall.Returns.Policy, f.GetProjectIamPolicyCall.Returns.Error
}
func (f *ServiceAccountsClient) ListServiceAccounts() ([]*gcpiam.ServiceAccount, error) {
	f.ListServiceAccountsCall.Lock()
	defer f.ListServiceAccountsCall.Unlock()
	f.ListServiceAccountsCall.CallCount++
	if f.ListServiceAccountsCall.Stub != nil {
		return f.ListServiceAccountsCall.Stub()
	}
	return f.ListServiceAccountsCall.Returns.ServiceAccountSlice, f.ListServiceAccountsCall.Returns.Error
}
func (f *ServiceAccountsClient) SetProjectIamPolicy(param1 *gcpcrm.Policy) (*gcpcrm.Policy, error) {
	f.SetProjectIamPolicyCall.Lock()
	defer f.SetProjectIamPolicyCall.Unlock()
	f.SetProjectIamPolicyCall.CallCount++
	f.SetProjectIamPolicyCall.Receives.Policy = param1
	if f.SetProjectIamPolicyCall.Stub != nil {
		return f.SetProjectIamPolicyCall.Stub(param1)
	}
	return f.SetProjectIamPolicyCall.Returns.Policy, f.SetProjectIamPolicyCall.Returns.Error
}
