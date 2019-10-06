package iam

import (
	"fmt"
	"strings"

	"github.com/genevieve/leftovers/common"
	gcpcrm "google.golang.org/api/cloudresourcemanager/v1"
	gcpiam "google.golang.org/api/iam/v1"
)

//go:generate faux --interface serviceAccountsClient --output fakes/service_accounts_client.go
type serviceAccountsClient interface {
	ListServiceAccounts() ([]*gcpiam.ServiceAccount, error)
	DeleteServiceAccount(account string) error

	GetProjectIamPolicy() (*gcpcrm.Policy, error)
	SetProjectIamPolicy(*gcpcrm.Policy) (*gcpcrm.Policy, error)
}

type ServiceAccounts struct {
	client        serviceAccountsClient
	projectName   string
	projectNumber string
	logger        logger
}

func NewServiceAccounts(client serviceAccountsClient, projectName string, projectNumber string, logger logger) ServiceAccounts {
	return ServiceAccounts{
		client:        client,
		projectName:   projectName,
		projectNumber: projectNumber,
		logger:        logger,
	}
}

func (s ServiceAccounts) List(filter string) ([]common.Deletable, error) {
	s.logger.Debugln("Listing IAM Service Accounts...")
	accounts, err := s.client.ListServiceAccounts()
	if err != nil {
		return nil, fmt.Errorf("List IAM Service Accounts: %s", err)
	}

	var resources []common.Deletable
	for _, account := range accounts {
		resource := NewServiceAccount(s.client, s.logger, account.Name, account.Email)

		if isDefault(s.projectName, s.projectNumber, account.Email) {
			continue
		}

		if !strings.Contains(resource.Name(), filter) {
			continue
		}

		proceed := s.logger.PromptWithDetails(resource.Type(), resource.Name())
		if !proceed {
			continue
		}

		resources = append(resources, resource)
	}

	return resources, nil
}

func (s ServiceAccounts) Type() string {
	return "service-account"
}

func isDefault(projectName, projectNumber, email string) bool {
	return email == fmt.Sprintf("%s@appspot.gserviceaccount.com", projectName) ||
		email == fmt.Sprintf("%s-compute@developer.gserviceaccount.com", projectNumber)
}
