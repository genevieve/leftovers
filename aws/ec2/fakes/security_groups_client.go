package fakes

import "github.com/aws/aws-sdk-go/service/ec2"

type SecurityGroupsClient struct {
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
}

func (e *SecurityGroupsClient) DescribeSecurityGroups(input *ec2.DescribeSecurityGroupsInput) (*ec2.DescribeSecurityGroupsOutput, error) {
	e.DescribeSecurityGroupsCall.CallCount++
	e.DescribeSecurityGroupsCall.Receives.Input = input

	return e.DescribeSecurityGroupsCall.Returns.Output, e.DescribeSecurityGroupsCall.Returns.Error
}

func (e *SecurityGroupsClient) RevokeSecurityGroupIngress(input *ec2.RevokeSecurityGroupIngressInput) (*ec2.RevokeSecurityGroupIngressOutput, error) {
	e.RevokeSecurityGroupIngressCall.CallCount++
	e.RevokeSecurityGroupIngressCall.Receives.Input = input

	return e.RevokeSecurityGroupIngressCall.Returns.Output, e.RevokeSecurityGroupIngressCall.Returns.Error
}

func (e *SecurityGroupsClient) RevokeSecurityGroupEgress(input *ec2.RevokeSecurityGroupEgressInput) (*ec2.RevokeSecurityGroupEgressOutput, error) {
	e.RevokeSecurityGroupEgressCall.CallCount++
	e.RevokeSecurityGroupEgressCall.Receives.Input = input

	return e.RevokeSecurityGroupEgressCall.Returns.Output, e.RevokeSecurityGroupEgressCall.Returns.Error
}

func (e *SecurityGroupsClient) DeleteSecurityGroup(input *ec2.DeleteSecurityGroupInput) (*ec2.DeleteSecurityGroupOutput, error) {
	e.DeleteSecurityGroupCall.CallCount++
	e.DeleteSecurityGroupCall.Receives.Input = input

	return e.DeleteSecurityGroupCall.Returns.Output, e.DeleteSecurityGroupCall.Returns.Error
}
