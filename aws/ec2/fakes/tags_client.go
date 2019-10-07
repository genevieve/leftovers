package fakes

import (
	"sync"

	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
)

type TagsClient struct {
	DeleteTagsCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DeleteTagsInput *awsec2.DeleteTagsInput
		}
		Returns struct {
			DeleteTagsOutput *awsec2.DeleteTagsOutput
			Error            error
		}
		Stub func(*awsec2.DeleteTagsInput) (*awsec2.DeleteTagsOutput, error)
	}
	DescribeTagsCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DescribeTagsInput *awsec2.DescribeTagsInput
		}
		Returns struct {
			DescribeTagsOutput *awsec2.DescribeTagsOutput
			Error              error
		}
		Stub func(*awsec2.DescribeTagsInput) (*awsec2.DescribeTagsOutput, error)
	}
}

func (f *TagsClient) DeleteTags(param1 *awsec2.DeleteTagsInput) (*awsec2.DeleteTagsOutput, error) {
	f.DeleteTagsCall.Lock()
	defer f.DeleteTagsCall.Unlock()
	f.DeleteTagsCall.CallCount++
	f.DeleteTagsCall.Receives.DeleteTagsInput = param1
	if f.DeleteTagsCall.Stub != nil {
		return f.DeleteTagsCall.Stub(param1)
	}
	return f.DeleteTagsCall.Returns.DeleteTagsOutput, f.DeleteTagsCall.Returns.Error
}
func (f *TagsClient) DescribeTags(param1 *awsec2.DescribeTagsInput) (*awsec2.DescribeTagsOutput, error) {
	f.DescribeTagsCall.Lock()
	defer f.DescribeTagsCall.Unlock()
	f.DescribeTagsCall.CallCount++
	f.DescribeTagsCall.Receives.DescribeTagsInput = param1
	if f.DescribeTagsCall.Stub != nil {
		return f.DescribeTagsCall.Stub(param1)
	}
	return f.DescribeTagsCall.Returns.DescribeTagsOutput, f.DescribeTagsCall.Returns.Error
}
