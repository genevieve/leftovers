package fakes

import (
	"sync"

	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
)

type KeyPairsClient struct {
	DeleteKeyPairCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DeleteKeyPairInput *awsec2.DeleteKeyPairInput
		}
		Returns struct {
			DeleteKeyPairOutput *awsec2.DeleteKeyPairOutput
			Error               error
		}
		Stub func(*awsec2.DeleteKeyPairInput) (*awsec2.DeleteKeyPairOutput, error)
	}
	DescribeKeyPairsCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DescribeKeyPairsInput *awsec2.DescribeKeyPairsInput
		}
		Returns struct {
			DescribeKeyPairsOutput *awsec2.DescribeKeyPairsOutput
			Error                  error
		}
		Stub func(*awsec2.DescribeKeyPairsInput) (*awsec2.DescribeKeyPairsOutput, error)
	}
}

func (f *KeyPairsClient) DeleteKeyPair(param1 *awsec2.DeleteKeyPairInput) (*awsec2.DeleteKeyPairOutput, error) {
	f.DeleteKeyPairCall.Lock()
	defer f.DeleteKeyPairCall.Unlock()
	f.DeleteKeyPairCall.CallCount++
	f.DeleteKeyPairCall.Receives.DeleteKeyPairInput = param1
	if f.DeleteKeyPairCall.Stub != nil {
		return f.DeleteKeyPairCall.Stub(param1)
	}
	return f.DeleteKeyPairCall.Returns.DeleteKeyPairOutput, f.DeleteKeyPairCall.Returns.Error
}
func (f *KeyPairsClient) DescribeKeyPairs(param1 *awsec2.DescribeKeyPairsInput) (*awsec2.DescribeKeyPairsOutput, error) {
	f.DescribeKeyPairsCall.Lock()
	defer f.DescribeKeyPairsCall.Unlock()
	f.DescribeKeyPairsCall.CallCount++
	f.DescribeKeyPairsCall.Receives.DescribeKeyPairsInput = param1
	if f.DescribeKeyPairsCall.Stub != nil {
		return f.DescribeKeyPairsCall.Stub(param1)
	}
	return f.DescribeKeyPairsCall.Returns.DescribeKeyPairsOutput, f.DescribeKeyPairsCall.Returns.Error
}
