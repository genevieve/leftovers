package fakes

import (
	"sync"

	awselbv2 "github.com/aws/aws-sdk-go/service/elbv2"
)

type TargetGroupsClient struct {
	DeleteTargetGroupCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DeleteTargetGroupInput *awselbv2.DeleteTargetGroupInput
		}
		Returns struct {
			DeleteTargetGroupOutput *awselbv2.DeleteTargetGroupOutput
			Error                   error
		}
		Stub func(*awselbv2.DeleteTargetGroupInput) (*awselbv2.DeleteTargetGroupOutput, error)
	}
	DescribeTargetGroupsCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DescribeTargetGroupsInput *awselbv2.DescribeTargetGroupsInput
		}
		Returns struct {
			DescribeTargetGroupsOutput *awselbv2.DescribeTargetGroupsOutput
			Error                      error
		}
		Stub func(*awselbv2.DescribeTargetGroupsInput) (*awselbv2.DescribeTargetGroupsOutput, error)
	}
}

func (f *TargetGroupsClient) DeleteTargetGroup(param1 *awselbv2.DeleteTargetGroupInput) (*awselbv2.DeleteTargetGroupOutput, error) {
	f.DeleteTargetGroupCall.Lock()
	defer f.DeleteTargetGroupCall.Unlock()
	f.DeleteTargetGroupCall.CallCount++
	f.DeleteTargetGroupCall.Receives.DeleteTargetGroupInput = param1
	if f.DeleteTargetGroupCall.Stub != nil {
		return f.DeleteTargetGroupCall.Stub(param1)
	}
	return f.DeleteTargetGroupCall.Returns.DeleteTargetGroupOutput, f.DeleteTargetGroupCall.Returns.Error
}
func (f *TargetGroupsClient) DescribeTargetGroups(param1 *awselbv2.DescribeTargetGroupsInput) (*awselbv2.DescribeTargetGroupsOutput, error) {
	f.DescribeTargetGroupsCall.Lock()
	defer f.DescribeTargetGroupsCall.Unlock()
	f.DescribeTargetGroupsCall.CallCount++
	f.DescribeTargetGroupsCall.Receives.DescribeTargetGroupsInput = param1
	if f.DescribeTargetGroupsCall.Stub != nil {
		return f.DescribeTargetGroupsCall.Stub(param1)
	}
	return f.DescribeTargetGroupsCall.Returns.DescribeTargetGroupsOutput, f.DescribeTargetGroupsCall.Returns.Error
}
