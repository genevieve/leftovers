package fakes

import "github.com/aws/aws-sdk-go/service/s3"

type BucketsClient struct {
	ListBucketsCall struct {
		CallCount int
		Receives  struct {
			Input *s3.ListBucketsInput
		}
		Returns struct {
			Output *s3.ListBucketsOutput
			Error  error
		}
	}

	DeleteBucketCall struct {
		CallCount int
		Receives  struct {
			Input *s3.DeleteBucketInput
		}
		Returns struct {
			Output *s3.DeleteBucketOutput
			Error  error
		}
	}
}

func (i *BucketsClient) ListBuckets(input *s3.ListBucketsInput) (*s3.ListBucketsOutput, error) {
	i.ListBucketsCall.CallCount++
	i.ListBucketsCall.Receives.Input = input

	return i.ListBucketsCall.Returns.Output, i.ListBucketsCall.Returns.Error
}

func (i *BucketsClient) DeleteBucket(input *s3.DeleteBucketInput) (*s3.DeleteBucketOutput, error) {
	i.DeleteBucketCall.CallCount++
	i.DeleteBucketCall.Receives.Input = input

	return i.DeleteBucketCall.Returns.Output, i.DeleteBucketCall.Returns.Error
}
