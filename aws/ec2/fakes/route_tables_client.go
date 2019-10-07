package fakes

import (
	"sync"

	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
)

type RouteTablesClient struct {
	DeleteRouteTableCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DeleteRouteTableInput *awsec2.DeleteRouteTableInput
		}
		Returns struct {
			DeleteRouteTableOutput *awsec2.DeleteRouteTableOutput
			Error                  error
		}
		Stub func(*awsec2.DeleteRouteTableInput) (*awsec2.DeleteRouteTableOutput, error)
	}
	DescribeRouteTablesCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DescribeRouteTablesInput *awsec2.DescribeRouteTablesInput
		}
		Returns struct {
			DescribeRouteTablesOutput *awsec2.DescribeRouteTablesOutput
			Error                     error
		}
		Stub func(*awsec2.DescribeRouteTablesInput) (*awsec2.DescribeRouteTablesOutput, error)
	}
	DisassociateRouteTableCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DisassociateRouteTableInput *awsec2.DisassociateRouteTableInput
		}
		Returns struct {
			DisassociateRouteTableOutput *awsec2.DisassociateRouteTableOutput
			Error                        error
		}
		Stub func(*awsec2.DisassociateRouteTableInput) (*awsec2.DisassociateRouteTableOutput, error)
	}
}

func (f *RouteTablesClient) DeleteRouteTable(param1 *awsec2.DeleteRouteTableInput) (*awsec2.DeleteRouteTableOutput, error) {
	f.DeleteRouteTableCall.Lock()
	defer f.DeleteRouteTableCall.Unlock()
	f.DeleteRouteTableCall.CallCount++
	f.DeleteRouteTableCall.Receives.DeleteRouteTableInput = param1
	if f.DeleteRouteTableCall.Stub != nil {
		return f.DeleteRouteTableCall.Stub(param1)
	}
	return f.DeleteRouteTableCall.Returns.DeleteRouteTableOutput, f.DeleteRouteTableCall.Returns.Error
}
func (f *RouteTablesClient) DescribeRouteTables(param1 *awsec2.DescribeRouteTablesInput) (*awsec2.DescribeRouteTablesOutput, error) {
	f.DescribeRouteTablesCall.Lock()
	defer f.DescribeRouteTablesCall.Unlock()
	f.DescribeRouteTablesCall.CallCount++
	f.DescribeRouteTablesCall.Receives.DescribeRouteTablesInput = param1
	if f.DescribeRouteTablesCall.Stub != nil {
		return f.DescribeRouteTablesCall.Stub(param1)
	}
	return f.DescribeRouteTablesCall.Returns.DescribeRouteTablesOutput, f.DescribeRouteTablesCall.Returns.Error
}
func (f *RouteTablesClient) DisassociateRouteTable(param1 *awsec2.DisassociateRouteTableInput) (*awsec2.DisassociateRouteTableOutput, error) {
	f.DisassociateRouteTableCall.Lock()
	defer f.DisassociateRouteTableCall.Unlock()
	f.DisassociateRouteTableCall.CallCount++
	f.DisassociateRouteTableCall.Receives.DisassociateRouteTableInput = param1
	if f.DisassociateRouteTableCall.Stub != nil {
		return f.DisassociateRouteTableCall.Stub(param1)
	}
	return f.DisassociateRouteTableCall.Returns.DisassociateRouteTableOutput, f.DisassociateRouteTableCall.Returns.Error
}
