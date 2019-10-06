package fakes

import (
	"sync"

	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
)

type SecurityGroupsClient struct {
	DeleteSecurityGroupCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DeleteSecurityGroupInput *awsec2.DeleteSecurityGroupInput
		}
		Returns struct {
			DeleteSecurityGroupOutput *awsec2.DeleteSecurityGroupOutput
			Error                     error
		}
		Stub func(*awsec2.DeleteSecurityGroupInput) (*awsec2.DeleteSecurityGroupOutput, error)
	}
	DescribeSecurityGroupsCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DescribeSecurityGroupsInput *awsec2.DescribeSecurityGroupsInput
		}
		Returns struct {
			DescribeSecurityGroupsOutput *awsec2.DescribeSecurityGroupsOutput
			Error                        error
		}
		Stub func(*awsec2.DescribeSecurityGroupsInput) (*awsec2.DescribeSecurityGroupsOutput, error)
	}
	RevokeSecurityGroupEgressCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			RevokeSecurityGroupEgressInput *awsec2.RevokeSecurityGroupEgressInput
		}
		Returns struct {
			RevokeSecurityGroupEgressOutput *awsec2.RevokeSecurityGroupEgressOutput
			Error                           error
		}
		Stub func(*awsec2.RevokeSecurityGroupEgressInput) (*awsec2.RevokeSecurityGroupEgressOutput, error)
	}
	RevokeSecurityGroupIngressCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			RevokeSecurityGroupIngressInput *awsec2.RevokeSecurityGroupIngressInput
		}
		Returns struct {
			RevokeSecurityGroupIngressOutput *awsec2.RevokeSecurityGroupIngressOutput
			Error                            error
		}
		Stub func(*awsec2.RevokeSecurityGroupIngressInput) (*awsec2.RevokeSecurityGroupIngressOutput, error)
	}
}

func (f *SecurityGroupsClient) DeleteSecurityGroup(param1 *awsec2.DeleteSecurityGroupInput) (*awsec2.DeleteSecurityGroupOutput, error) {
	f.DeleteSecurityGroupCall.Lock()
	defer f.DeleteSecurityGroupCall.Unlock()
	f.DeleteSecurityGroupCall.CallCount++
	f.DeleteSecurityGroupCall.Receives.DeleteSecurityGroupInput = param1
	if f.DeleteSecurityGroupCall.Stub != nil {
		return f.DeleteSecurityGroupCall.Stub(param1)
	}
	return f.DeleteSecurityGroupCall.Returns.DeleteSecurityGroupOutput, f.DeleteSecurityGroupCall.Returns.Error
}
func (f *SecurityGroupsClient) DescribeSecurityGroups(param1 *awsec2.DescribeSecurityGroupsInput) (*awsec2.DescribeSecurityGroupsOutput, error) {
	f.DescribeSecurityGroupsCall.Lock()
	defer f.DescribeSecurityGroupsCall.Unlock()
	f.DescribeSecurityGroupsCall.CallCount++
	f.DescribeSecurityGroupsCall.Receives.DescribeSecurityGroupsInput = param1
	if f.DescribeSecurityGroupsCall.Stub != nil {
		return f.DescribeSecurityGroupsCall.Stub(param1)
	}
	return f.DescribeSecurityGroupsCall.Returns.DescribeSecurityGroupsOutput, f.DescribeSecurityGroupsCall.Returns.Error
}
func (f *SecurityGroupsClient) RevokeSecurityGroupEgress(param1 *awsec2.RevokeSecurityGroupEgressInput) (*awsec2.RevokeSecurityGroupEgressOutput, error) {
	f.RevokeSecurityGroupEgressCall.Lock()
	defer f.RevokeSecurityGroupEgressCall.Unlock()
	f.RevokeSecurityGroupEgressCall.CallCount++
	f.RevokeSecurityGroupEgressCall.Receives.RevokeSecurityGroupEgressInput = param1
	if f.RevokeSecurityGroupEgressCall.Stub != nil {
		return f.RevokeSecurityGroupEgressCall.Stub(param1)
	}
	return f.RevokeSecurityGroupEgressCall.Returns.RevokeSecurityGroupEgressOutput, f.RevokeSecurityGroupEgressCall.Returns.Error
}
func (f *SecurityGroupsClient) RevokeSecurityGroupIngress(param1 *awsec2.RevokeSecurityGroupIngressInput) (*awsec2.RevokeSecurityGroupIngressOutput, error) {
	f.RevokeSecurityGroupIngressCall.Lock()
	defer f.RevokeSecurityGroupIngressCall.Unlock()
	f.RevokeSecurityGroupIngressCall.CallCount++
	f.RevokeSecurityGroupIngressCall.Receives.RevokeSecurityGroupIngressInput = param1
	if f.RevokeSecurityGroupIngressCall.Stub != nil {
		return f.RevokeSecurityGroupIngressCall.Stub(param1)
	}
	return f.RevokeSecurityGroupIngressCall.Returns.RevokeSecurityGroupIngressOutput, f.RevokeSecurityGroupIngressCall.Returns.Error
}
