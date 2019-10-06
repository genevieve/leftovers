package fakes

import (
	"sync"

	awss3 "github.com/aws/aws-sdk-go/service/s3"
)

type BucketsClient struct {
	DeleteBucketCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DeleteBucketInput *awss3.DeleteBucketInput
		}
		Returns struct {
			DeleteBucketOutput *awss3.DeleteBucketOutput
			Error              error
		}
		Stub func(*awss3.DeleteBucketInput) (*awss3.DeleteBucketOutput, error)
	}
	DeleteObjectsCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DeleteObjectsInput *awss3.DeleteObjectsInput
		}
		Returns struct {
			DeleteObjectsOutput *awss3.DeleteObjectsOutput
			Error               error
		}
		Stub func(*awss3.DeleteObjectsInput) (*awss3.DeleteObjectsOutput, error)
	}
	ListBucketsCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			ListBucketsInput *awss3.ListBucketsInput
		}
		Returns struct {
			ListBucketsOutput *awss3.ListBucketsOutput
			Error             error
		}
		Stub func(*awss3.ListBucketsInput) (*awss3.ListBucketsOutput, error)
	}
	ListObjectVersionsCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			ListObjectVersionsInput *awss3.ListObjectVersionsInput
		}
		Returns struct {
			ListObjectVersionsOutput *awss3.ListObjectVersionsOutput
			Error                    error
		}
		Stub func(*awss3.ListObjectVersionsInput) (*awss3.ListObjectVersionsOutput, error)
	}
}

func (f *BucketsClient) DeleteBucket(param1 *awss3.DeleteBucketInput) (*awss3.DeleteBucketOutput, error) {
	f.DeleteBucketCall.Lock()
	defer f.DeleteBucketCall.Unlock()
	f.DeleteBucketCall.CallCount++
	f.DeleteBucketCall.Receives.DeleteBucketInput = param1
	if f.DeleteBucketCall.Stub != nil {
		return f.DeleteBucketCall.Stub(param1)
	}
	return f.DeleteBucketCall.Returns.DeleteBucketOutput, f.DeleteBucketCall.Returns.Error
}
func (f *BucketsClient) DeleteObjects(param1 *awss3.DeleteObjectsInput) (*awss3.DeleteObjectsOutput, error) {
	f.DeleteObjectsCall.Lock()
	defer f.DeleteObjectsCall.Unlock()
	f.DeleteObjectsCall.CallCount++
	f.DeleteObjectsCall.Receives.DeleteObjectsInput = param1
	if f.DeleteObjectsCall.Stub != nil {
		return f.DeleteObjectsCall.Stub(param1)
	}
	return f.DeleteObjectsCall.Returns.DeleteObjectsOutput, f.DeleteObjectsCall.Returns.Error
}
func (f *BucketsClient) ListBuckets(param1 *awss3.ListBucketsInput) (*awss3.ListBucketsOutput, error) {
	f.ListBucketsCall.Lock()
	defer f.ListBucketsCall.Unlock()
	f.ListBucketsCall.CallCount++
	f.ListBucketsCall.Receives.ListBucketsInput = param1
	if f.ListBucketsCall.Stub != nil {
		return f.ListBucketsCall.Stub(param1)
	}
	return f.ListBucketsCall.Returns.ListBucketsOutput, f.ListBucketsCall.Returns.Error
}
func (f *BucketsClient) ListObjectVersions(param1 *awss3.ListObjectVersionsInput) (*awss3.ListObjectVersionsOutput, error) {
	f.ListObjectVersionsCall.Lock()
	defer f.ListObjectVersionsCall.Unlock()
	f.ListObjectVersionsCall.CallCount++
	f.ListObjectVersionsCall.Receives.ListObjectVersionsInput = param1
	if f.ListObjectVersionsCall.Stub != nil {
		return f.ListObjectVersionsCall.Stub(param1)
	}
	return f.ListObjectVersionsCall.Returns.ListObjectVersionsOutput, f.ListObjectVersionsCall.Returns.Error
}
