package fakes

import "github.com/aws/aws-sdk-go/service/ec2"

type KeyPairsClient struct {
	DescribeKeyPairsCall struct {
		CallCount int
		Receives  struct {
			Input *ec2.DescribeKeyPairsInput
		}
		Returns struct {
			Output *ec2.DescribeKeyPairsOutput
			Error  error
		}
	}

	DeleteKeyPairCall struct {
		CallCount int
		Receives  struct {
			Input *ec2.DeleteKeyPairInput
		}
		Returns struct {
			Output *ec2.DeleteKeyPairOutput
			Error  error
		}
	}
}

func (e *KeyPairsClient) DescribeKeyPairs(input *ec2.DescribeKeyPairsInput) (*ec2.DescribeKeyPairsOutput, error) {
	e.DescribeKeyPairsCall.CallCount++
	e.DescribeKeyPairsCall.Receives.Input = input

	return e.DescribeKeyPairsCall.Returns.Output, e.DescribeKeyPairsCall.Returns.Error
}

func (e *KeyPairsClient) DeleteKeyPair(input *ec2.DeleteKeyPairInput) (*ec2.DeleteKeyPairOutput, error) {
	e.DeleteKeyPairCall.CallCount++
	e.DeleteKeyPairCall.Receives.Input = input

	return e.DeleteKeyPairCall.Returns.Output, e.DeleteKeyPairCall.Returns.Error
}
