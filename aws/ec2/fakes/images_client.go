package fakes

import (
	"sync"

	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
)

type ImagesClient struct {
	DeregisterImageCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DeregisterImageInput *awsec2.DeregisterImageInput
		}
		Returns struct {
			DeregisterImageOutput *awsec2.DeregisterImageOutput
			Error                 error
		}
		Stub func(*awsec2.DeregisterImageInput) (*awsec2.DeregisterImageOutput, error)
	}
	DescribeImagesCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DescribeImagesInput *awsec2.DescribeImagesInput
		}
		Returns struct {
			DescribeImagesOutput *awsec2.DescribeImagesOutput
			Error                error
		}
		Stub func(*awsec2.DescribeImagesInput) (*awsec2.DescribeImagesOutput, error)
	}
}

func (f *ImagesClient) DeregisterImage(param1 *awsec2.DeregisterImageInput) (*awsec2.DeregisterImageOutput, error) {
	f.DeregisterImageCall.Lock()
	defer f.DeregisterImageCall.Unlock()
	f.DeregisterImageCall.CallCount++
	f.DeregisterImageCall.Receives.DeregisterImageInput = param1
	if f.DeregisterImageCall.Stub != nil {
		return f.DeregisterImageCall.Stub(param1)
	}
	return f.DeregisterImageCall.Returns.DeregisterImageOutput, f.DeregisterImageCall.Returns.Error
}
func (f *ImagesClient) DescribeImages(param1 *awsec2.DescribeImagesInput) (*awsec2.DescribeImagesOutput, error) {
	f.DescribeImagesCall.Lock()
	defer f.DescribeImagesCall.Unlock()
	f.DescribeImagesCall.CallCount++
	f.DescribeImagesCall.Receives.DescribeImagesInput = param1
	if f.DescribeImagesCall.Stub != nil {
		return f.DescribeImagesCall.Stub(param1)
	}
	return f.DescribeImagesCall.Returns.DescribeImagesOutput, f.DescribeImagesCall.Returns.Error
}
