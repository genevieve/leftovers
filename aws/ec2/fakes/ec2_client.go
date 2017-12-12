package fakes

import (
	"github.com/aws/aws-sdk-go/service/ec2"
)

type EC2Client struct {
	DescribeSecurityGroupsCall struct {
		CallCount int
		Receives  struct {
			Input *ec2.DescribeSecurityGroupsInput
		}
		Returns struct {
			Output *ec2.DescribeSecurityGroupsOutput
			Error  error
		}
	}

	RevokeSecurityGroupIngressCall struct {
		CallCount int
		Receives  struct {
			Input *ec2.RevokeSecurityGroupIngressInput
		}
		Returns struct {
			Output *ec2.RevokeSecurityGroupIngressOutput
			Error  error
		}
	}

	RevokeSecurityGroupEgressCall struct {
		CallCount int
		Receives  struct {
			Input *ec2.RevokeSecurityGroupEgressInput
		}
		Returns struct {
			Output *ec2.RevokeSecurityGroupEgressOutput
			Error  error
		}
	}

	DeleteSecurityGroupCall struct {
		CallCount int
		Receives  struct {
			Input *ec2.DeleteSecurityGroupInput
		}
		Returns struct {
			Output *ec2.DeleteSecurityGroupOutput
			Error  error
		}
	}

	DescribeTagsCall struct {
		CallCount int
		Receives  struct {
			Input *ec2.DescribeTagsInput
		}
		Returns struct {
			Output *ec2.DescribeTagsOutput
			Error  error
		}
	}

	DeleteTagsCall struct {
		CallCount int
		Receives  struct {
			Input *ec2.DeleteTagsInput
		}
		Returns struct {
			Output *ec2.DeleteTagsOutput
			Error  error
		}
	}

	DescribeVolumesCall struct {
		CallCount int
		Receives  struct {
			Input *ec2.DescribeVolumesInput
		}
		Returns struct {
			Output *ec2.DescribeVolumesOutput
			Error  error
		}
	}

	DeleteVolumeCall struct {
		CallCount int
		Receives  struct {
			Input *ec2.DeleteVolumeInput
		}
		Returns struct {
			Output *ec2.DeleteVolumeOutput
			Error  error
		}
	}
}

func (e *EC2Client) DescribeSecurityGroups(input *ec2.DescribeSecurityGroupsInput) (*ec2.DescribeSecurityGroupsOutput, error) {
	e.DescribeSecurityGroupsCall.CallCount++
	e.DescribeSecurityGroupsCall.Receives.Input = input

	return e.DescribeSecurityGroupsCall.Returns.Output, e.DescribeSecurityGroupsCall.Returns.Error
}

func (e *EC2Client) RevokeSecurityGroupIngress(input *ec2.RevokeSecurityGroupIngressInput) (*ec2.RevokeSecurityGroupIngressOutput, error) {
	e.RevokeSecurityGroupIngressCall.CallCount++
	e.RevokeSecurityGroupIngressCall.Receives.Input = input

	return e.RevokeSecurityGroupIngressCall.Returns.Output, e.RevokeSecurityGroupIngressCall.Returns.Error
}

func (e *EC2Client) RevokeSecurityGroupEgress(input *ec2.RevokeSecurityGroupEgressInput) (*ec2.RevokeSecurityGroupEgressOutput, error) {
	e.RevokeSecurityGroupEgressCall.CallCount++
	e.RevokeSecurityGroupEgressCall.Receives.Input = input

	return e.RevokeSecurityGroupEgressCall.Returns.Output, e.RevokeSecurityGroupEgressCall.Returns.Error
}

func (e *EC2Client) DeleteSecurityGroup(input *ec2.DeleteSecurityGroupInput) (*ec2.DeleteSecurityGroupOutput, error) {
	e.DeleteSecurityGroupCall.CallCount++
	e.DeleteSecurityGroupCall.Receives.Input = input

	return e.DeleteSecurityGroupCall.Returns.Output, e.DeleteSecurityGroupCall.Returns.Error
}

func (e *EC2Client) DescribeTags(input *ec2.DescribeTagsInput) (*ec2.DescribeTagsOutput, error) {
	e.DescribeTagsCall.CallCount++
	e.DescribeTagsCall.Receives.Input = input

	return e.DescribeTagsCall.Returns.Output, e.DescribeTagsCall.Returns.Error
}

func (e *EC2Client) DeleteTags(input *ec2.DeleteTagsInput) (*ec2.DeleteTagsOutput, error) {
	e.DeleteTagsCall.CallCount++
	e.DeleteTagsCall.Receives.Input = input

	return e.DeleteTagsCall.Returns.Output, e.DeleteTagsCall.Returns.Error
}

func (e *EC2Client) DescribeVolumes(input *ec2.DescribeVolumesInput) (*ec2.DescribeVolumesOutput, error) {
	e.DescribeVolumesCall.CallCount++
	e.DescribeVolumesCall.Receives.Input = input

	return e.DescribeVolumesCall.Returns.Output, e.DescribeVolumesCall.Returns.Error
}

func (e *EC2Client) DeleteVolume(input *ec2.DeleteVolumeInput) (*ec2.DeleteVolumeOutput, error) {
	e.DeleteVolumeCall.CallCount++
	e.DeleteVolumeCall.Receives.Input = input

	return e.DeleteVolumeCall.Returns.Output, e.DeleteVolumeCall.Returns.Error
}
