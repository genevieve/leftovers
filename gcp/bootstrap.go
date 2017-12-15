package gcp

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/oauth2/google"

	compute "google.golang.org/api/compute/v1"
)

type logger interface {
	Printf(m string, a ...interface{})
	Prompt(m string) bool
}

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

	projectId := struct {
		ProjectId string `json:"project_id"`
	}{}
	json.Unmarshal(key, &projectId)

	config, err := google.JWTConfigFromJSON(key, compute.ComputeScope)
	if err != nil {
		log.Fatalf("Creating JWT config from GCP_CREDENTIALS: %s", err)
	}

	client, err := compute.New(config.Client(context.Background()))
	if err != nil {
		log.Fatalf("Creating GCP client: %s", err)
	}

	networks, err := client.Networks.List(projectId.ProjectId).Do()
	if err != nil {
		log.Fatalf("Listing networks: %s", err)
	}
	for _, e := range networks.Items {
		fmt.Println(e.Name)
	}
}
