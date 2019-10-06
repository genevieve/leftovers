package fakes

import (
	"sync"

	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
)

type SubnetsClient struct {
	DeleteSubnetCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DeleteSubnetInput *awsec2.DeleteSubnetInput
		}
		Returns struct {
			DeleteSubnetOutput *awsec2.DeleteSubnetOutput
			Error              error
		}
		Stub func(*awsec2.DeleteSubnetInput) (*awsec2.DeleteSubnetOutput, error)
	}
	DescribeSubnetsCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DescribeSubnetsInput *awsec2.DescribeSubnetsInput
		}
		Returns struct {
			DescribeSubnetsOutput *awsec2.DescribeSubnetsOutput
			Error                 error
		}
		Stub func(*awsec2.DescribeSubnetsInput) (*awsec2.DescribeSubnetsOutput, error)
	}
}

func (f *SubnetsClient) DeleteSubnet(param1 *awsec2.DeleteSubnetInput) (*awsec2.DeleteSubnetOutput, error) {
	f.DeleteSubnetCall.Lock()
	defer f.DeleteSubnetCall.Unlock()
	f.DeleteSubnetCall.CallCount++
	f.DeleteSubnetCall.Receives.DeleteSubnetInput = param1
	if f.DeleteSubnetCall.Stub != nil {
		return f.DeleteSubnetCall.Stub(param1)
	}
	return f.DeleteSubnetCall.Returns.DeleteSubnetOutput, f.DeleteSubnetCall.Returns.Error
}
func (f *SubnetsClient) DescribeSubnets(param1 *awsec2.DescribeSubnetsInput) (*awsec2.DescribeSubnetsOutput, error) {
	f.DescribeSubnetsCall.Lock()
	defer f.DescribeSubnetsCall.Unlock()
	f.DescribeSubnetsCall.CallCount++
	f.DescribeSubnetsCall.Receives.DescribeSubnetsInput = param1
	if f.DescribeSubnetsCall.Stub != nil {
		return f.DescribeSubnetsCall.Stub(param1)
	}
	return f.DescribeSubnetsCall.Returns.DescribeSubnetsOutput, f.DescribeSubnetsCall.Returns.Error
}
