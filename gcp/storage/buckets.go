package storage

import (
	"fmt"

	"github.com/genevieve/leftovers/common"
	gcpstorage "google.golang.org/api/storage/v1"
)

//go:generate faux --interface bucketsClient --output fakes/buckets_client.go
type bucketsClient interface {
	ListBuckets() (*gcpstorage.Buckets, error)
	DeleteBucket(bucket string) error

	ListObjects(bucket string) (*gcpstorage.Objects, error)
	DeleteObject(bucket, object string, generation int64) error
}

type Buckets struct {
	client bucketsClient
	logger logger
}

func NewBuckets(client bucketsClient, logger logger) Buckets {
	return Buckets{
		client: client,
		logger: logger,
	}
}

func (i Buckets) List(filter string, regex bool) ([]common.Deletable, error) {
	i.logger.Debugln("Listing Storage Buckets...")
	buckets, err := i.client.ListBuckets()
	if err != nil {
		return nil, fmt.Errorf("List Storage Buckets: %s", err)
	}

	var resources []common.Deletable
	for _, bucket := range buckets.Items {
		resource := NewBucket(i.client, bucket.Name)

		if !common.ResourceMatches(resource.Name(), filter, regex) {
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

func (b Buckets) Type() string {
	return "bucket"
}
