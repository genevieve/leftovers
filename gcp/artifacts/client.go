package artifacts

import (
	"fmt"

	gcpartifact "google.golang.org/api/artifactregistry/v1"
	"google.golang.org/api/googleapi"
)

type client struct {
	project string
	logger  logger

	service      *gcpartifact.Service
	repositories *gcpartifact.ProjectsLocationsRepositoriesService
}

func NewClient(project string, service *gcpartifact.Service, logger logger) client {
	return client{
		project: project,
		logger:  logger,

		service:      service,
		repositories: service.Projects.Locations.Repositories,
	}
}

func (c client) ListRepositories(region string) (*gcpartifact.ListRepositoriesResponse, error) {
	parent := fmt.Sprintf("projects/%v/locations/%v", c.project, region)
	return c.repositories.List(parent).Do()
}

func (c client) DeleteRepository(name string) error {
	return c.wait(c.repositories.Delete(name))
}

type request interface {
	Do(...googleapi.CallOption) (*gcpartifact.Operation, error)
}

func (c client) wait(request request) error {
	op, err := request.Do()
	if err != nil {
		if gerr, ok := err.(*googleapi.Error); ok {
			if gerr.Code == 404 {
				return nil
			}
		}
		return fmt.Errorf("do request: %s", err)
	}

	waiter := NewOperationWaiter(op, c.service, c.project, c.logger)

	return waiter.Wait()
}
