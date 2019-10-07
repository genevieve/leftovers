package fakes

import (
	"sync"

	gcpstorage "google.golang.org/api/storage/v1"
)

type BucketsClient struct {
	DeleteBucketCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Bucket string
		}
		Returns struct {
			Error error
		}
		Stub func(string) error
	}
	DeleteObjectCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Bucket     string
			Object     string
			Generation int64
		}
		Returns struct {
			Error error
		}
		Stub func(string, string, int64) error
	}
	ListBucketsCall struct {
		sync.Mutex
		CallCount int
		Returns   struct {
			Buckets *gcpstorage.Buckets
			Error   error
		}
		Stub func() (*gcpstorage.Buckets, error)
	}
	ListObjectsCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Bucket string
		}
		Returns struct {
			Objects *gcpstorage.Objects
			Error   error
		}
		Stub func(string) (*gcpstorage.Objects, error)
	}
}

func (f *BucketsClient) DeleteBucket(param1 string) error {
	f.DeleteBucketCall.Lock()
	defer f.DeleteBucketCall.Unlock()
	f.DeleteBucketCall.CallCount++
	f.DeleteBucketCall.Receives.Bucket = param1
	if f.DeleteBucketCall.Stub != nil {
		return f.DeleteBucketCall.Stub(param1)
	}
	return f.DeleteBucketCall.Returns.Error
}
func (f *BucketsClient) DeleteObject(param1 string, param2 string, param3 int64) error {
	f.DeleteObjectCall.Lock()
	defer f.DeleteObjectCall.Unlock()
	f.DeleteObjectCall.CallCount++
	f.DeleteObjectCall.Receives.Bucket = param1
	f.DeleteObjectCall.Receives.Object = param2
	f.DeleteObjectCall.Receives.Generation = param3
	if f.DeleteObjectCall.Stub != nil {
		return f.DeleteObjectCall.Stub(param1, param2, param3)
	}
	return f.DeleteObjectCall.Returns.Error
}
func (f *BucketsClient) ListBuckets() (*gcpstorage.Buckets, error) {
	f.ListBucketsCall.Lock()
	defer f.ListBucketsCall.Unlock()
	f.ListBucketsCall.CallCount++
	if f.ListBucketsCall.Stub != nil {
		return f.ListBucketsCall.Stub()
	}
	return f.ListBucketsCall.Returns.Buckets, f.ListBucketsCall.Returns.Error
}
func (f *BucketsClient) ListObjects(param1 string) (*gcpstorage.Objects, error) {
	f.ListObjectsCall.Lock()
	defer f.ListObjectsCall.Unlock()
	f.ListObjectsCall.CallCount++
	f.ListObjectsCall.Receives.Bucket = param1
	if f.ListObjectsCall.Stub != nil {
		return f.ListObjectsCall.Stub(param1)
	}
	return f.ListObjectsCall.Returns.Objects, f.ListObjectsCall.Returns.Error
}
