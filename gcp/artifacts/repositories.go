package artifacts

import (
	"fmt"
	"strings"

	"github.com/genevieve/leftovers/common"
	gcpartifact "google.golang.org/api/artifactregistry/v1"
)

type Repositories struct {
	client  repositoriesClient
	logger  logger
	regions map[string]string
}

//go:generate faux --interface repositoriesClient --output fakes/repositories_client.go
type repositoriesClient interface {
	ListRepositories(region string) (*gcpartifact.ListRepositoriesResponse, error)
	DeleteRepository(cluster string) error
}

func NewRepositories(client repositoriesClient, logger logger, regions map[string]string) Repositories {
	return Repositories{
		client:  client,
		logger:  logger,
		regions: regions,
	}
}

func (c Repositories) List(filter string) ([]common.Deletable, error) {
	repositories := []*gcpartifact.Repository{}
	for _, region := range c.regions {
		c.logger.Debugf("Listing Repositories for region %s...\n", region)
		l, err := c.client.ListRepositories(region)
		if err != nil {
			return nil, fmt.Errorf("list repositories for region %v: %w", region, err)
		}

		repositories = append(repositories, l.Repositories...)
	}

	deletables := []common.Deletable{}
	for _, repository := range repositories {
		resource := NewRepository(c.client, repository.Name)

		if !strings.Contains(resource.Name(), filter) {
			continue
		}

		proceed := c.logger.PromptWithDetails(resource.Type(), resource.Name())
		if !proceed {
			continue
		}

		deletables = append(deletables, resource)
	}

	return deletables, nil
}

func (c Repositories) Type() string {
	return "repository"
}
