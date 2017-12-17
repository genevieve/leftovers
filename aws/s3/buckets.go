package s3

import (
	"fmt"

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

		err = u.emptyAndDelete(b)
		if err == nil {
			u.logger.Printf("SUCCESS deleting bucket %s\n", n)
		} else {
			u.logger.Printf("ERROR deleting bucket %s: %s\n", n, err)
		}
	}

	return nil
}

func (u Buckets) emptyAndDelete(b *awss3.Bucket) error {
	_, err := u.client.DeleteBucket(&awss3.DeleteBucketInput{Bucket: b.Name})
	if err != nil {
		ec2err, ok := err.(awserr.Error)

		if ok && ec2err.Code() == "BucketNotEmpty" {
			resp, err := u.client.ListObjectVersions(&awss3.ListObjectVersionsInput{Bucket: b.Name})
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
				Bucket: b.Name,
				Delete: &awss3.Delete{Objects: objects},
			})
			if err != nil {
				return err
			}

			return u.emptyAndDelete(b)
		}

		return err
	}

	return nil
}
