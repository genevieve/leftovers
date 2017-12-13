package s3

import (
	"fmt"

	awss3 "github.com/aws/aws-sdk-go/service/s3"
)

type bucketsClient interface {
	ListBuckets(*awss3.ListBucketsInput) (*awss3.ListBucketsOutput, error)
	DeleteBucket(*awss3.DeleteBucketInput) (*awss3.DeleteBucketOutput, error)
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

func (u Buckets) Delete() error {
	buckets, err := u.client.ListBuckets(&awss3.ListBucketsInput{})
	if err != nil {
		return fmt.Errorf("Listing buckets: %s", err)
	}

	for _, b := range buckets.Buckets {
		n := *b.Name

		if !u.manager.IsInRegion(n) {
			continue
		}

		proceed := u.logger.Prompt(fmt.Sprintf("Are you sure you want to delete bucket %s?", n))
		if !proceed {
			continue
		}

		_, err = u.client.DeleteBucket(&awss3.DeleteBucketInput{Bucket: b.Name})
		if err == nil {
			u.logger.Printf("SUCCESS deleting bucket %s\n", n)
		} else {
			u.logger.Printf("ERROR deleting bucket %s: %s\n", n, err)
		}
	}

	return nil
}
