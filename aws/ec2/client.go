package ec2

import (
	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
)

type ec2Client interface {
	DescribeVolumes(*awsec2.DescribeVolumesInput) (*awsec2.DescribeVolumesOutput, error)
	DeleteVolume(*awsec2.DeleteVolumeInput) (*awsec2.DeleteVolumeOutput, error)

	DescribeTags(*awsec2.DescribeTagsInput) (*awsec2.DescribeTagsOutput, error)
	DeleteTags(*awsec2.DeleteTagsInput) (*awsec2.DeleteTagsOutput, error)

	DescribeKeyPairs(*awsec2.DescribeKeyPairsInput) (*awsec2.DescribeKeyPairsOutput, error)
	DeleteKeyPair(*awsec2.DeleteKeyPairInput) (*awsec2.DeleteKeyPairOutput, error)
}
