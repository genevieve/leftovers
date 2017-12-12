package fakes

import (
	"github.com/aws/aws-sdk-go/service/ec2"
)

type EC2Client struct {
	DescribeTagsCall struct {
		CallCount int
		Receives  struct {
			Input *ec2.DescribeTagsInput
		}
		Returns struct {
			Output *ec2.DescribeTagsOutput
			Error  error
		}
	}

	DeleteTagsCall struct {
		CallCount int
		Receives  struct {
			Input *ec2.DeleteTagsInput
		}
		Returns struct {
			Output *ec2.DeleteTagsOutput
			Error  error
		}
	}

}

func (e *EC2Client) DescribeTags(input *ec2.DescribeTagsInput) (*ec2.DescribeTagsOutput, error) {
	e.DescribeTagsCall.CallCount++
	e.DescribeTagsCall.Receives.Input = input

	return e.DescribeTagsCall.Returns.Output, e.DescribeTagsCall.Returns.Error
}

func (e *EC2Client) DeleteTags(input *ec2.DeleteTagsInput) (*ec2.DeleteTagsOutput, error) {
	e.DeleteTagsCall.CallCount++
	e.DeleteTagsCall.Receives.Input = input

	return e.DeleteTagsCall.Returns.Output, e.DeleteTagsCall.Returns.Error
}
