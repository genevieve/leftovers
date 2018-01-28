package s3

import (
	"fmt"
	"strings"

	awss3 "github.com/aws/aws-sdk-go/service/s3"
)

type bucketsClient interface {
	ListBuckets(*awss3.ListBucketsInput) (*awss3.ListBucketsOutput, error)
	DeleteBucket(*awss3.DeleteBucketInput) (*awss3.DeleteBucketOutput, error)

	ListObjectVersions(*awss3.ListObjectVersionsInput) (*awss3.ListObjectVersionsOutput, error)
	DeleteObjects(*awss3.DeleteObjectsInput) (*awss3.DeleteObjectsOutput, error)
}

type Buckets struct {
	client  bucketsClient
	logger  logger
	manager bucketManager
}

func NewBuckets(client bucketsClient, logger logger, manager bucketManager) Buckets {
	return Buckets{
		client:  client,
		logger:  logger,
		manager: manager,
	}
}

func (b Buckets) List(filter string) (map[string]string, error) {
	buckets, err := b.list(filter)
	if err != nil {
		return nil, err
	}

	delete := map[string]string{}
	for _, bucket := range buckets {
		delete[bucket.identifier] = ""
	}

	return delete, nil
}

func (b Buckets) list(filter string) ([]Bucket, error) {
	buckets, err := b.client.ListBuckets(&awss3.ListBucketsInput{})
	if err != nil {
		return nil, fmt.Errorf("Listing buckets: %s", err)
	}

	var resources []Bucket
	for _, bucket := range buckets.Buckets {
		resource := NewBucket(b.client, bucket.Name)

		if !strings.Contains(resource.identifier, filter) {
			continue
		}

		if !b.manager.IsInRegion(resource.identifier) {
			continue
		}

		proceed := b.logger.Prompt(fmt.Sprintf("Are you sure you want to delete bucket %s?", resource.identifier))
		if !proceed {
			continue
		}

		resources = append(resources, resource)
	}

	return resources, nil
}

func (b Buckets) Delete(buckets map[string]string) error {
	var resources []Bucket
	for name, _ := range buckets {
		resources = append(resources, NewBucket(b.client, &name))
	}

	return b.cleanup(resources)
}

func (b Buckets) cleanup(resources []Bucket) error {
	for _, resource := range resources {
		err := resource.Delete()

		if err == nil {
			b.logger.Printf("SUCCESS deleting bucket %s\n", resource.identifier)
		} else {
			b.logger.Printf("ERROR deleting bucket %s: %s\n", resource.identifier, err)
		}
	}

	return nil
}
