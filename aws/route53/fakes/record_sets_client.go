package fakes

import (
	"sync"

	awsroute53 "github.com/aws/aws-sdk-go/service/route53"
)

type RecordSetsClient struct {
	ChangeResourceRecordSetsCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			ChangeResourceRecordSetsInput *awsroute53.ChangeResourceRecordSetsInput
		}
		Returns struct {
			ChangeResourceRecordSetsOutput *awsroute53.ChangeResourceRecordSetsOutput
			Error                          error
		}
		Stub func(*awsroute53.ChangeResourceRecordSetsInput) (*awsroute53.ChangeResourceRecordSetsOutput, error)
	}
	ListResourceRecordSetsCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			ListResourceRecordSetsInput *awsroute53.ListResourceRecordSetsInput
		}
		Returns struct {
			ListResourceRecordSetsOutput *awsroute53.ListResourceRecordSetsOutput
			Error                        error
		}
		Stub func(*awsroute53.ListResourceRecordSetsInput) (*awsroute53.ListResourceRecordSetsOutput, error)
	}
}

func (f *RecordSetsClient) ChangeResourceRecordSets(param1 *awsroute53.ChangeResourceRecordSetsInput) (*awsroute53.ChangeResourceRecordSetsOutput, error) {
	f.ChangeResourceRecordSetsCall.Lock()
	defer f.ChangeResourceRecordSetsCall.Unlock()
	f.ChangeResourceRecordSetsCall.CallCount++
	f.ChangeResourceRecordSetsCall.Receives.ChangeResourceRecordSetsInput = param1
	if f.ChangeResourceRecordSetsCall.Stub != nil {
		return f.ChangeResourceRecordSetsCall.Stub(param1)
	}
	return f.ChangeResourceRecordSetsCall.Returns.ChangeResourceRecordSetsOutput, f.ChangeResourceRecordSetsCall.Returns.Error
}
func (f *RecordSetsClient) ListResourceRecordSets(param1 *awsroute53.ListResourceRecordSetsInput) (*awsroute53.ListResourceRecordSetsOutput, error) {
	f.ListResourceRecordSetsCall.Lock()
	defer f.ListResourceRecordSetsCall.Unlock()
	f.ListResourceRecordSetsCall.CallCount++
	f.ListResourceRecordSetsCall.Receives.ListResourceRecordSetsInput = param1
	if f.ListResourceRecordSetsCall.Stub != nil {
		return f.ListResourceRecordSetsCall.Stub(param1)
	}
	return f.ListResourceRecordSetsCall.Returns.ListResourceRecordSetsOutput, f.ListResourceRecordSetsCall.Returns.Error
}
