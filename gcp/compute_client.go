package gcp

import (
	compute "google.golang.org/api/compute/v1"
)

type computeClient struct {
	networks *compute.NetworksService
}

func (c computeClient) List(project string) (*compute.NetworkList, error) {
	return c.networks.List(project).Do()
}

func (c computeClient) Delete(project, network string) (*compute.Operation, error) {
	return c.networks.Delete(project, network).Do()
}
