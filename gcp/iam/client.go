package iam

import (
	"fmt"
	"time"

	"google.golang.org/api/googleapi"
	gcpiam "google.golang.org/api/iam/v1"
)

type client struct {
	project string

	service         *gcpiam.Service
	serviceAccounts *gcpiam.ProjectsServiceAccountsService
}

func NewClient(project string, service *gcpiam.Service) client {
	return client{
		project:         project,
		service:         service,
		serviceAccounts: service.Projects.ServiceAccounts,
	}
}

// ListServiceAccounts will loop over every page of results
// and return the full list of service accounts. To prevent
// backend errors from repeated calls, there is a 2s delay.
func (c client) ListServiceAccounts() ([]*gcpiam.ServiceAccount, error) {
	path := fmt.Sprintf("projects/%s", c.project)

	resp, err := c.serviceAccounts.List(path).Do()
	if err != nil {
		return []*gcpiam.ServiceAccount{}, err
	}

	serviceAccounts := resp.Accounts

	for resp.NextPageToken != "" {
		resp, err = c.serviceAccounts.List(path).PageToken(resp.NextPageToken).Do()
		if err != nil {
			return []*gcpiam.ServiceAccount{}, err
		}

		serviceAccounts = append(serviceAccounts, resp.Accounts...)

		time.Sleep(2 * time.Second)
	}

	return serviceAccounts, nil
}

func (c client) DeleteServiceAccount(account string) error {
	_, err := c.serviceAccounts.Delete(account).Do()
	if err != nil {
		gerr, ok := err.(*googleapi.Error)
		if ok && gerr != nil && gerr.Code == 404 {
			return nil
		}

		return err
	}
	return nil
}
