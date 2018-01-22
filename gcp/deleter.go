package gcp

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/genevievelesperance/leftovers/gcp/compute"
	"golang.org/x/oauth2/google"
	gcpcompute "google.golang.org/api/compute/v1"
)

type logger interface {
	Printf(message string, a ...interface{})
	Println(message string)
	Prompt(message string) bool
}

type resource interface {
	List(filter string) (map[string]string, error)
	Delete(resources map[string]string)
}

type Deleter struct {
	resources []resource
}

func (d Deleter) Delete(filter string) error {
	for _, r := range d.resources {
		list, err := r.List(filter)
		if err != nil {
			return err
		}

		r.Delete(list)
	}
	return nil
}

func NewDeleter(logger logger, keyPath string) (Deleter, error) {
	if keyPath == "" {
		return Deleter{}, errors.New("Missing service account key path.")
	}

	key, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return Deleter{}, fmt.Errorf("Reading service account key path %s: %s", keyPath, err)
	}

	p := struct {
		ProjectId string `json:"project_id"`
	}{}
	if err := json.Unmarshal(key, &p); err != nil {
		return Deleter{}, fmt.Errorf("Unmarshalling account key for project id: %s", err)
	}

	logger.Println(fmt.Sprintf("Cleaning gcp project: %s.", p.ProjectId))

	config, err := google.JWTConfigFromJSON(key, gcpcompute.ComputeScope)
	if err != nil {
		return Deleter{}, fmt.Errorf("Creating jwt config: %s", err)
	}

	service, err := gcpcompute.New(config.Client(context.Background()))
	if err != nil {
		return Deleter{}, fmt.Errorf("Creating gcp client: %s", err)
	}

	client := compute.NewClient(p.ProjectId, service, logger)

	regions, err := client.ListRegions()
	if err != nil {
		return Deleter{}, fmt.Errorf("Listing regions: %s", err)
	}

	zones, err := client.ListZones()
	if err != nil {
		return Deleter{}, fmt.Errorf("Listing zones: %s", err)
	}

	return Deleter{
		resources: []resource{
			compute.NewForwardingRules(client, logger, regions),
			compute.NewGlobalForwardingRules(client, logger),
			compute.NewFirewalls(client, logger),
			compute.NewTargetHttpProxies(client, logger),
			compute.NewTargetHttpsProxies(client, logger),
			compute.NewUrlMaps(client, logger),
			compute.NewTargetPools(client, logger, regions),
			compute.NewBackendServices(client, logger),
			compute.NewInstances(client, logger, zones),
			compute.NewInstanceGroups(client, logger, zones),
			compute.NewGlobalHealthChecks(client, logger),
			compute.NewHttpHealthChecks(client, logger),
			compute.NewHttpsHealthChecks(client, logger),
			compute.NewDisks(client, logger, zones),
			compute.NewNetworks(client, logger),
			compute.NewAddresses(client, logger, regions),
		},
	}, nil
}
