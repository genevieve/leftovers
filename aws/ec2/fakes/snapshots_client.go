package fakes

import "github.com/aws/aws-sdk-go/service/ec2"

type SnapshotsClient struct {
	DescribeSnapshotsCall struct {
		CallCount int
		Receives  struct {
			Input *ec2.DescribeSnapshotsInput
		}
		Returns struct {
			Output *ec2.DescribeSnapshotsOutput
			Error  error
		}
	}

	DeleteSnapshotCall struct {
		CallCount int
		Receives  struct {
			Input *ec2.DeleteSnapshotInput
		}
		Returns struct {
			Output *ec2.DeleteSnapshotOutput
			Error  error
		}
	}
}

func (e *SnapshotsClient) DescribeSnapshots(input *ec2.DescribeSnapshotsInput) (*ec2.DescribeSnapshotsOutput, error) {
	e.DescribeSnapshotsCall.CallCount++
	e.DescribeSnapshotsCall.Receives.Input = input

	return e.DescribeSnapshotsCall.Returns.Output, e.DescribeSnapshotsCall.Returns.Error
}

func (e *SnapshotsClient) DeleteSnapshot(input *ec2.DeleteSnapshotInput) (*ec2.DeleteSnapshotOutput, error) {
	e.DeleteSnapshotCall.CallCount++
	e.DeleteSnapshotCall.Receives.Input = input

	return e.DeleteSnapshotCall.Returns.Output, e.DeleteSnapshotCall.Returns.Error
}
