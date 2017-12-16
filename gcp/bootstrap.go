package gcp

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	"golang.org/x/oauth2/google"

	compute "google.golang.org/api/compute/v1"
)

type resource interface {
	Delete() error
}

func Bootstrap(logger logger, serviceAccountKey string) {
	if serviceAccountKey == "" {
		log.Fatal("Missing GCP_SERVICE_ACCOUNT_KEY.")
	}

	key, err := ioutil.ReadFile(serviceAccountKey)
	if err != nil {
		log.Fatal("Reading GCP_SERVICE_ACCOUNT_KEY: %s", err)
	}

	p := struct {
		ProjectId string `json:"project_id"`
	}{}
	json.Unmarshal(key, &p)

	config, err := google.JWTConfigFromJSON(key, compute.ComputeScope)
	if err != nil {
		log.Fatalf("Creating JWT config from GCP_CREDENTIALS: %s", err)
	}

	service, err := compute.New(config.Client(context.Background()))
	if err != nil {
		log.Fatalf("Creating GCP client: %s", err)
	}

	client := NewComputeClient(p.ProjectId, service)

	zones, err := client.ListZones()
	if err != nil {
		log.Fatalf("Listing zones: %s", err)
	}

	ne := NewNetworks(client, logger)
	di := NewDisks(client, logger, zones)

	resources := []resource{ne, di}
	for _, r := range resources {
		if err := r.Delete(); err != nil {
			log.Fatalf("\n\n%s\n", err)
		}
	}
}
