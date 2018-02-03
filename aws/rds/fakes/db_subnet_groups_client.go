package fakes

import awsrds "github.com/aws/aws-sdk-go/service/rds"

type DBSubnetGroupsClient struct {
	DescribeDBSubnetGroupsCall struct {
		CallCount int
		Returns   struct {
			Output *awsrds.DescribeDBSubnetGroupsOutput
			Error  error
		}
	}
	DeleteDBSubnetGroupCall struct {
		CallCount int
		Receives  struct {
			Input *awsrds.DeleteDBSubnetGroupInput
		}
		Returns struct {
			Output *awsrds.DeleteDBSubnetGroupOutput
			Error  error
		}
	}
}

func (d *DBSubnetGroupsClient) DeleteDBSubnetGroup(input *awsrds.DeleteDBSubnetGroupInput) (*awsrds.DeleteDBSubnetGroupOutput, error) {
	d.DeleteDBSubnetGroupCall.CallCount++
	d.DeleteDBSubnetGroupCall.Receives.Input = input

	return d.DeleteDBSubnetGroupCall.Returns.Output, d.DeleteDBSubnetGroupCall.Returns.Error
}

func (d *DBSubnetGroupsClient) DescribeDBSubnetGroups(input *awsrds.DescribeDBSubnetGroupsInput) (*awsrds.DescribeDBSubnetGroupsOutput, error) {
	d.DescribeDBSubnetGroupsCall.CallCount++

	return d.DescribeDBSubnetGroupsCall.Returns.Output, d.DescribeDBSubnetGroupsCall.Returns.Error
}
