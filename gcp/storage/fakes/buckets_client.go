package fakes

import gcpsql "google.golang.org/api/storage/v1"

type BucketsClient struct {
	ListBucketsCall struct {
		CallCount int
		Returns   struct {
			Output *gcpsql.Buckets
			Error  error
		}
	}

	DeleteBucketCall struct {
		CallCount int
		Receives  struct {
			Bucket string
		}
		Returns struct {
			Error error
		}
	}
}

func (u *BucketsClient) ListBuckets() (*gcpsql.Buckets, error) {
	u.ListBucketsCall.CallCount++

	return u.ListBucketsCall.Returns.Output, u.ListBucketsCall.Returns.Error
}

func (u *BucketsClient) DeleteBucket(bucket string) error {
	u.DeleteBucketCall.CallCount++
	u.DeleteBucketCall.Receives.Bucket = bucket

	return u.DeleteBucketCall.Returns.Error
}
