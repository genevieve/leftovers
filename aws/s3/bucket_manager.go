package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type bucketManager interface {
	IsInRegion(bucket, region string) bool
}

type BucketManager struct{}

func NewBucketManager() BucketManager {
	return BucketManager{}
}

func (u BucketManager) IsInRegion(bucket, region string) bool {
	sess := session.Must(session.NewSession())
	r, _ := s3manager.GetBucketRegion(aws.BackgroundContext(), sess, bucket, "us-west-1")
	return region == r
}
