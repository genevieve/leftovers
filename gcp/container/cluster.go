package container

import "fmt"

type Cluster struct {
	name     string
	location string
	client   clustersClient
}

func NewCluster(client clustersClient, location string, name string) Cluster {
	return Cluster{
		name:     name,
		location: location,
		client:   client,
	}
}

func (c Cluster) Delete() error {
	err := c.client.DeleteCluster(c.location, c.name)
	if err != nil {
		return fmt.Errorf("Delete: %s", err)
	}
	return nil
}

func (c Cluster) Name() string {
	return c.name
}

func (c Cluster) Type() string {
	return "Container Cluster"
}
