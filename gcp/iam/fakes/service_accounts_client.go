package fakes

import gcpiam "google.golang.org/api/iam/v1"

type ServiceAccountsClient struct {
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

func (u *ServiceAccountsClient) ListServiceAccounts() ([]*gcpiam.ServiceAccount, error) {
	u.ListServiceAccountsCall.CallCount++

	return u.ListServiceAccountsCall.Returns.Output, u.ListServiceAccountsCall.Returns.Error
}

func (u *ServiceAccountsClient) DeleteServiceAccount(account string) error {
	u.DeleteServiceAccountCall.CallCount++
	u.DeleteServiceAccountCall.Receives.ServiceAccount = account

	return u.DeleteServiceAccountCall.Returns.Error
}
