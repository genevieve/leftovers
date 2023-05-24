package artifacts

import "fmt"

type Repository struct {
	name   string
	client repositoriesClient
}

func NewRepository(client repositoriesClient, name string) Repository {
	return Repository{
		name:   name,
		client: client,
	}
}

func (c Repository) Delete() error {
	err := c.client.DeleteRepository(c.name)
	if err != nil {
		return fmt.Errorf("Delete: %s", err)
	}
	return nil
}

func (c Repository) Name() string {
	return c.name
}

func (c Repository) Type() string {
	return "Artifact Repository"
}
