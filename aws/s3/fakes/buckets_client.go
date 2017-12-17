package fakes

import awss3 "github.com/aws/aws-sdk-go/service/s3"

type BucketsClient struct {
	ListBucketsCall struct {
		CallCount int
		Receives  struct {
			Input *awss3.ListBucketsInput
		}
		Returns struct {
			Output *awss3.ListBucketsOutput
			Error  error
		}
	}

	DeleteBucketCall struct {
		CallCount int
		Receives  struct {
			Input *awss3.DeleteBucketInput
		}
		Returns struct {
			Output *awss3.DeleteBucketOutput
			Error  error
		}
	}

	ListObjectVersionsCall struct {
		CallCount int
		Receives  struct {
			Input *awss3.ListObjectVersionsInput
		}
		Returns struct {
			Output *awss3.ListObjectVersionsOutput
			Error  error
		}
	}

	DeleteObjectsCall struct {
		CallCount int
		Receives  struct {
			Input *awss3.DeleteObjectsInput
		}
		Returns struct {
			Output *awss3.DeleteObjectsOutput
			Error  error
		}
	}
}

func (i *BucketsClient) ListBuckets(input *awss3.ListBucketsInput) (*awss3.ListBucketsOutput, error) {
	i.ListBucketsCall.CallCount++
	i.ListBucketsCall.Receives.Input = input

	return i.ListBucketsCall.Returns.Output, i.ListBucketsCall.Returns.Error
}

func (i *BucketsClient) DeleteBucket(input *awss3.DeleteBucketInput) (*awss3.DeleteBucketOutput, error) {
	i.DeleteBucketCall.CallCount++
	i.DeleteBucketCall.Receives.Input = input

	return i.DeleteBucketCall.Returns.Output, i.DeleteBucketCall.Returns.Error
}

func (i *BucketsClient) ListObjectVersions(input *awss3.ListObjectVersionsInput) (*awss3.ListObjectVersionsOutput, error) {
	i.ListObjectVersionsCall.CallCount++
	i.ListObjectVersionsCall.Receives.Input = input

	return i.ListObjectVersionsCall.Returns.Output, i.ListObjectVersionsCall.Returns.Error
}

func (i *BucketsClient) DeleteObjects(input *awss3.DeleteObjectsInput) (*awss3.DeleteObjectsOutput, error) {
	i.DeleteObjectsCall.CallCount++
	i.DeleteObjectsCall.Receives.Input = input

	return i.DeleteObjectsCall.Returns.Output, i.DeleteObjectsCall.Returns.Error
}
