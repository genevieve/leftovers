package fakes

import "github.com/aws/aws-sdk-go/service/elbv2"

type TargetGroupsClient struct {
	DescribeTargetGroupsCall struct {
		CallCount int
		Receives  struct {
			Input *elbv2.DescribeTargetGroupsInput
		}
		Returns struct {
			Output *elbv2.DescribeTargetGroupsOutput
			Error  error
		}
	}

	DeleteTargetGroupCall struct {
		CallCount int
		Receives  struct {
			Input *elbv2.DeleteTargetGroupInput
		}
		Returns struct {
			Output *elbv2.DeleteTargetGroupOutput
			Error  error
		}
	}
}

func (e *TargetGroupsClient) DescribeTargetGroups(input *elbv2.DescribeTargetGroupsInput) (*elbv2.DescribeTargetGroupsOutput, error) {
	e.DescribeTargetGroupsCall.CallCount++
	e.DescribeTargetGroupsCall.Receives.Input = input

	return e.DescribeTargetGroupsCall.Returns.Output, e.DescribeTargetGroupsCall.Returns.Error
}

func (e *TargetGroupsClient) DeleteTargetGroup(input *elbv2.DeleteTargetGroupInput) (*elbv2.DeleteTargetGroupOutput, error) {
	e.DeleteTargetGroupCall.CallCount++
	e.DeleteTargetGroupCall.Receives.Input = input

	return e.DeleteTargetGroupCall.Returns.Output, e.DeleteTargetGroupCall.Returns.Error
}
