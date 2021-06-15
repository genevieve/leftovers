package sql

import (
	"fmt"
	"github.com/genevieve/leftovers/common"
	gcpsql "google.golang.org/api/sqladmin/v1beta4"
)

//go:generate faux --interface instancesClient --output fakes/instances_client.go
type instancesClient interface {
	ListInstances() (*gcpsql.InstancesListResponse, error)
	DeleteInstance(user string) error
}

type Instances struct {
	client instancesClient
	logger logger
}

func NewInstances(client instancesClient, logger logger) Instances {
	return Instances{
		client: client,
		logger: logger,
	}
}

func (i Instances) List(filter string, regex bool) ([]common.Deletable, error) {
	i.logger.Debugln("Listing SQL Instances...")
	instances, err := i.client.ListInstances()
	if err != nil {
		return nil, fmt.Errorf("List SQL Instances: %s", err)
	}

	var resources []common.Deletable
	for _, instance := range instances.Items {
		resource := NewInstance(i.client, instance.Name)

		if !common.ResourceMatches(resource.name, filter, regex) {
			continue
		}

		proceed := i.logger.PromptWithDetails(resource.Type(), resource.Name())
		if !proceed {
			continue
		}

		resources = append(resources, resource)
	}

	return resources, nil
}

func (i Instances) Type() string {
	return "sql-instance"
}
