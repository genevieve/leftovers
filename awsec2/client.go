package awsec2

import (
	"github.com/aws/aws-sdk-go/service/ec2"
)

type ec2Client interface {
	DescribeVolumes(*ec2.DescribeVolumesInput) (*ec2.DescribeVolumesOutput, error)
	DeleteVolume(*ec2.DeleteVolumeInput) (*ec2.DeleteVolumeOutput, error)
}
