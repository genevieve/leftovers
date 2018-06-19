package fakes

import gcpstorage "google.golang.org/api/storage/v1"

type BucketsClient struct {
	ListBucketsCall struct {
		CallCount int
		Returns   struct {
			Output *gcpstorage.Buckets
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

	ListObjectsCall struct {
		CallCount int
		Receives  struct {
			Bucket string
		}
		Returns struct {
			Objects *gcpstorage.Objects
			Error   error
		}
	}

	DeleteObjectCall struct {
		CallCount int
		Receives  struct {
			Bucket     string
			Object     string
			Generation int64
		}
		Returns struct {
			Error error
		}
	}
}

func (u *BucketsClient) ListBuckets() (*gcpstorage.Buckets, error) {
	u.ListBucketsCall.CallCount++

	return u.ListBucketsCall.Returns.Output, u.ListBucketsCall.Returns.Error
}

func (u *BucketsClient) DeleteBucket(bucket string) error {
	u.DeleteBucketCall.CallCount++
	u.DeleteBucketCall.Receives.Bucket = bucket

	return u.DeleteBucketCall.Returns.Error
}

func (b *BucketsClient) ListObjects(bucket string) (*gcpstorage.Objects, error) {
	b.ListObjectsCall.CallCount++
	b.ListObjectsCall.Receives.Bucket = bucket

	return b.ListObjectsCall.Returns.Objects, b.ListObjectsCall.Returns.Error
}

func (b *BucketsClient) DeleteObject(bucket, object string, generation int64) error {
	b.DeleteObjectCall.CallCount++
	b.DeleteObjectCall.Receives.Bucket = bucket
	b.DeleteObjectCall.Receives.Object = object
	b.DeleteObjectCall.Receives.Generation = generation

	return b.DeleteObjectCall.Returns.Error
}
