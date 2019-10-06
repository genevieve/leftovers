package fakes

import (
	"sync"

	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
)

type SnapshotsClient struct {
	DeleteSnapshotCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DeleteSnapshotInput *awsec2.DeleteSnapshotInput
		}
		Returns struct {
			DeleteSnapshotOutput *awsec2.DeleteSnapshotOutput
			Error                error
		}
		Stub func(*awsec2.DeleteSnapshotInput) (*awsec2.DeleteSnapshotOutput, error)
	}
	DescribeSnapshotsCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DescribeSnapshotsInput *awsec2.DescribeSnapshotsInput
		}
		Returns struct {
			DescribeSnapshotsOutput *awsec2.DescribeSnapshotsOutput
			Error                   error
		}
		Stub func(*awsec2.DescribeSnapshotsInput) (*awsec2.DescribeSnapshotsOutput, error)
	}
}

func (f *SnapshotsClient) DeleteSnapshot(param1 *awsec2.DeleteSnapshotInput) (*awsec2.DeleteSnapshotOutput, error) {
	f.DeleteSnapshotCall.Lock()
	defer f.DeleteSnapshotCall.Unlock()
	f.DeleteSnapshotCall.CallCount++
	f.DeleteSnapshotCall.Receives.DeleteSnapshotInput = param1
	if f.DeleteSnapshotCall.Stub != nil {
		return f.DeleteSnapshotCall.Stub(param1)
	}
	return f.DeleteSnapshotCall.Returns.DeleteSnapshotOutput, f.DeleteSnapshotCall.Returns.Error
}
func (f *SnapshotsClient) DescribeSnapshots(param1 *awsec2.DescribeSnapshotsInput) (*awsec2.DescribeSnapshotsOutput, error) {
	f.DescribeSnapshotsCall.Lock()
	defer f.DescribeSnapshotsCall.Unlock()
	f.DescribeSnapshotsCall.CallCount++
	f.DescribeSnapshotsCall.Receives.DescribeSnapshotsInput = param1
	if f.DescribeSnapshotsCall.Stub != nil {
		return f.DescribeSnapshotsCall.Stub(param1)
	}
	return f.DescribeSnapshotsCall.Returns.DescribeSnapshotsOutput, f.DescribeSnapshotsCall.Returns.Error
}
