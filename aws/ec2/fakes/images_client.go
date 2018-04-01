package fakes

import "github.com/aws/aws-sdk-go/service/ec2"

type ImagesClient struct {
	DescribeImagesCall struct {
		CallCount int
		Receives  struct {
			Input *ec2.DescribeImagesInput
		}
		Returns struct {
			Output *ec2.DescribeImagesOutput
			Error  error
		}
	}

	DeregisterImageCall struct {
		CallCount int
		Receives  struct {
			Input *ec2.DeregisterImageInput
		}
		Returns struct {
			Output *ec2.DeregisterImageOutput
			Error  error
		}
	}
}

func (i *ImagesClient) DescribeImages(input *ec2.DescribeImagesInput) (*ec2.DescribeImagesOutput, error) {
	i.DescribeImagesCall.CallCount++
	i.DescribeImagesCall.Receives.Input = input

	return i.DescribeImagesCall.Returns.Output, i.DescribeImagesCall.Returns.Error
}

func (i *ImagesClient) DeregisterImage(input *ec2.DeregisterImageInput) (*ec2.DeregisterImageOutput, error) {
	i.DeregisterImageCall.CallCount++
	i.DeregisterImageCall.Receives.Input = input

	return i.DeregisterImageCall.Returns.Output, i.DeregisterImageCall.Returns.Error
}
