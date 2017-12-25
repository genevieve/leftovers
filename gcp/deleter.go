package gcp

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/genevievelesperance/leftovers/gcp/compute"
	"golang.org/x/oauth2/google"
	gcpcompute "google.golang.org/api/compute/v1"
)

type logger interface {
	Printf(m string, a ...interface{})
	Prompt(m string) bool
}

type resource interface {
	Delete() error
}

type Deleter struct {
	resources []resource
}

func (d Deleter) Delete() error {
	for _, r := range d.resources {
		if err := r.Delete(); err != nil {
			return err
		}
	}
	return nil
}

func NewDeleter(logger logger, serviceAccountKey string) Deleter {
	if serviceAccountKey == "" {
		log.Fatal("Missing BBL_GCP_SERVICE_ACCOUNT_KEY.")
	}

	key, err := ioutil.ReadFile(serviceAccountKey)
	if err != nil {
		log.Fatal("Reading %s: %s", serviceAccountKey, err)
	}

	p := struct {
		ProjectId string `json:"project_id"`
	}{}
	json.Unmarshal(key, &p)

	config, err := google.JWTConfigFromJSON(key, gcpcompute.ComputeScope)
	if err != nil {
		log.Fatalf("Creating jwt config: %s", err)
	}

	service, err := gcpcompute.New(config.Client(context.Background()))
	if err != nil {
		log.Fatalf("Creating gcp client: %s", err)
	}

	client := compute.NewClient(p.ProjectId, service)

	regions, err := client.ListRegions()
	if err != nil {
		log.Fatalf("Listing regions: %s", err)
	}

	zones, err := client.ListZones()
	if err != nil {
		log.Fatalf("Listing zones: %s", err)
	}

	ne := compute.NewNetworks(client, logger)
	fw := compute.NewForwardingRules(client, logger, regions)
	tp := compute.NewTargetPools(client, logger, regions)
	in := compute.NewInstances(client, logger, zones)
	ht := compute.NewHttpHealthChecks(client, logger)
	hs := compute.NewHttpsHealthChecks(client, logger)
	ba := compute.NewBackendServices(client, logger)
	di := compute.NewDisks(client, logger, zones)

	return Deleter{
		resources: []resource{ne, fw, tp, in, di, ht, hs, ba},
	}
}
