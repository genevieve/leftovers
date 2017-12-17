package fakes

import "github.com/aws/aws-sdk-go/service/ec2"

type RouteTablesClient struct {
	DescribeRouteTablesCall struct {
		CallCount int
		Receives  struct {
			Input *ec2.DescribeRouteTablesInput
		}
		Returns struct {
			Output *ec2.DescribeRouteTablesOutput
			Error  error
		}
	}

	DeleteRouteTableCall struct {
		CallCount int
		Receives  struct {
			Input *ec2.DeleteRouteTableInput
		}
		Returns struct {
			Output *ec2.DeleteRouteTableOutput
			Error  error
		}
	}
}

func (i *RouteTablesClient) DescribeRouteTables(input *ec2.DescribeRouteTablesInput) (*ec2.DescribeRouteTablesOutput, error) {
	i.DescribeRouteTablesCall.CallCount++
	i.DescribeRouteTablesCall.Receives.Input = input

	return i.DescribeRouteTablesCall.Returns.Output, i.DescribeRouteTablesCall.Returns.Error
}

func (i *RouteTablesClient) DeleteRouteTable(input *ec2.DeleteRouteTableInput) (*ec2.DeleteRouteTableOutput, error) {
	i.DeleteRouteTableCall.CallCount++
	i.DeleteRouteTableCall.Receives.Input = input

	return i.DeleteRouteTableCall.Returns.Output, i.DeleteRouteTableCall.Returns.Error
}
