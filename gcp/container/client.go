package container

import (
	"fmt"

	gcpcontainer "google.golang.org/api/container/v1"
	"google.golang.org/api/googleapi"
)

type client struct {
	project string
	logger  logger

	service    *gcpcontainer.Service
	containers *gcpcontainer.ProjectsLocationsClustersService
}

func NewClient(project string, service *gcpcontainer.Service, logger logger) client {
	return client{
		project: project,
		logger:  logger,

		service:    service,
		containers: service.Projects.Locations.Clusters,
	}
}

func (c client) ListClusters() (*gcpcontainer.ListClustersResponse, error) {
	parent := fmt.Sprintf("projects/%v/locations/-", c.project)
	return c.containers.List(parent).Do()
}

func (c client) DeleteCluster(zone string, cluster string) error {
	name := fmt.Sprintf("projects/%v/locations/%v/clusters/%v", c.project, zone, cluster)
	return c.wait(c.containers.Delete(name))
}

type request interface {
	Do(...googleapi.CallOption) (*gcpcontainer.Operation, error)
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
