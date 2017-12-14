package fakes

import "github.com/aws/aws-sdk-go/service/ec2"

type AddressesClient struct {
	DescribeAddressesCall struct {
		CallCount int
		Receives  struct {
			Input *ec2.DescribeAddressesInput
		}
		Returns struct {
			Output *ec2.DescribeAddressesOutput
			Error  error
		}
	}

	ReleaseAddressCall struct {
		CallCount int
		Receives  struct {
			Input *ec2.ReleaseAddressInput
		}
		Returns struct {
			Output *ec2.ReleaseAddressOutput
			Error  error
		}
	}
}

func (e *AddressesClient) DescribeAddresses(input *ec2.DescribeAddressesInput) (*ec2.DescribeAddressesOutput, error) {
	e.DescribeAddressesCall.CallCount++
	e.DescribeAddressesCall.Receives.Input = input

	return e.DescribeAddressesCall.Returns.Output, e.DescribeAddressesCall.Returns.Error
}

func (e *AddressesClient) ReleaseAddress(input *ec2.ReleaseAddressInput) (*ec2.ReleaseAddressOutput, error) {
	e.ReleaseAddressCall.CallCount++
	e.ReleaseAddressCall.Receives.Input = input

	return e.ReleaseAddressCall.Returns.Output, e.ReleaseAddressCall.Returns.Error
}
