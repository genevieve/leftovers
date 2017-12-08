package ec2

import (
	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
)

type ec2Client interface {
	DescribeVolumes(*awsec2.DescribeVolumesInput) (*awsec2.DescribeVolumesOutput, error)
	DeleteVolume(*awsec2.DeleteVolumeInput) (*awsec2.DeleteVolumeOutput, error)
}
