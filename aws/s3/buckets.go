package s3

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
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
	for name, _ := range buckets {
		err := b.emptyAndDelete(name)

		if err == nil {
			b.logger.Printf("SUCCESS deleting bucket %s\n", name)
		} else {
			b.logger.Printf("ERROR deleting bucket %s: %s\n", name, err)
		}
	}

	return nil
}

func (u Buckets) emptyAndDelete(name string) error {
	_, err := u.client.DeleteBucket(&awss3.DeleteBucketInput{Bucket: aws.String(name)})

	if err != nil {
		ec2err, ok := err.(awserr.Error)

		if ok && ec2err.Code() == "BucketNotEmpty" {
			resp, err := u.client.ListObjectVersions(&awss3.ListObjectVersionsInput{Bucket: aws.String(name)})
			if err != nil {
				return err
			}

			objects := make([]*awss3.ObjectIdentifier, 0)

			if len(resp.DeleteMarkers) != 0 {
				for _, v := range resp.DeleteMarkers {
					objects = append(objects, &awss3.ObjectIdentifier{
						Key:       v.Key,
						VersionId: v.VersionId,
					})
				}
			}

			if len(resp.Versions) != 0 {
				for _, v := range resp.Versions {
					objects = append(objects, &awss3.ObjectIdentifier{
						Key:       v.Key,
						VersionId: v.VersionId,
					})
				}
			}

			_, err = u.client.DeleteObjects(&awss3.DeleteObjectsInput{
				Bucket: aws.String(name),
				Delete: &awss3.Delete{Objects: objects},
			})
			if err != nil {
				return err
			}

			return u.emptyAndDelete(name)
		}

		return err
	}

	return nil
}
