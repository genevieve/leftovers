package ec2

import (
	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
)

type ec2Client interface {
	DescribeTags(*awsec2.DescribeTagsInput) (*awsec2.DescribeTagsOutput, error)
	DeleteTags(*awsec2.DeleteTagsInput) (*awsec2.DeleteTagsOutput, error)
}
