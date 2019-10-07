package fakes

import (
	"sync"

	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
)

type VpcsClient struct {
	DeleteVpcCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DeleteVpcInput *awsec2.DeleteVpcInput
		}
		Returns struct {
			DeleteVpcOutput *awsec2.DeleteVpcOutput
			Error           error
		}
		Stub func(*awsec2.DeleteVpcInput) (*awsec2.DeleteVpcOutput, error)
	}
	DescribeVpcsCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DescribeVpcsInput *awsec2.DescribeVpcsInput
		}
		Returns struct {
			DescribeVpcsOutput *awsec2.DescribeVpcsOutput
			Error              error
		}
		Stub func(*awsec2.DescribeVpcsInput) (*awsec2.DescribeVpcsOutput, error)
	}
}

func (f *VpcsClient) DeleteVpc(param1 *awsec2.DeleteVpcInput) (*awsec2.DeleteVpcOutput, error) {
	f.DeleteVpcCall.Lock()
	defer f.DeleteVpcCall.Unlock()
	f.DeleteVpcCall.CallCount++
	f.DeleteVpcCall.Receives.DeleteVpcInput = param1
	if f.DeleteVpcCall.Stub != nil {
		return f.DeleteVpcCall.Stub(param1)
	}
	return f.DeleteVpcCall.Returns.DeleteVpcOutput, f.DeleteVpcCall.Returns.Error
}
func (f *VpcsClient) DescribeVpcs(param1 *awsec2.DescribeVpcsInput) (*awsec2.DescribeVpcsOutput, error) {
	f.DescribeVpcsCall.Lock()
	defer f.DescribeVpcsCall.Unlock()
	f.DescribeVpcsCall.CallCount++
	f.DescribeVpcsCall.Receives.DescribeVpcsInput = param1
	if f.DescribeVpcsCall.Stub != nil {
		return f.DescribeVpcsCall.Stub(param1)
	}
	return f.DescribeVpcsCall.Returns.DescribeVpcsOutput, f.DescribeVpcsCall.Returns.Error
}
