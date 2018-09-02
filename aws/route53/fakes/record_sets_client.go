package fakes

import (
	awsroute53 "github.com/aws/aws-sdk-go/service/route53"
)

type RecordSetsClient struct {
	ListResourceRecordSetsCall struct {
		CallCount int
		Receives  []ListResourceRecordSetsCallReceive
		Returns   []ListResourceRecordSetsCallReturn
	}

	ChangeResourceRecordSetsCall struct {
		CallCount int
		Receives  struct {
			Input *awsroute53.ChangeResourceRecordSetsInput
		}
		Returns struct {
			Output *awsroute53.ChangeResourceRecordSetsOutput
			Error  error
		}
	}
}

type ListResourceRecordSetsCallReceive struct {
	Input *awsroute53.ListResourceRecordSetsInput
}

type ListResourceRecordSetsCallReturn struct {
	Output *awsroute53.ListResourceRecordSetsOutput
	Error  error
}

func (r *RecordSetsClient) ListResourceRecordSets(input *awsroute53.ListResourceRecordSetsInput) (*awsroute53.ListResourceRecordSetsOutput, error) {
	r.ListResourceRecordSetsCall.CallCount++
	r.ListResourceRecordSetsCall.Receives = append(r.ListResourceRecordSetsCall.Receives, ListResourceRecordSetsCallReceive{Input: input})

	if len(r.ListResourceRecordSetsCall.Returns) < r.ListResourceRecordSetsCall.CallCount {
		return nil, nil
	}

	return r.ListResourceRecordSetsCall.Returns[r.ListResourceRecordSetsCall.CallCount-1].Output, r.ListResourceRecordSetsCall.Returns[r.ListResourceRecordSetsCall.CallCount-1].Error
}

func (r *RecordSetsClient) ChangeResourceRecordSets(input *awsroute53.ChangeResourceRecordSetsInput) (*awsroute53.ChangeResourceRecordSetsOutput, error) {
	r.ChangeResourceRecordSetsCall.CallCount++
	r.ChangeResourceRecordSetsCall.Receives.Input = input

	return r.ChangeResourceRecordSetsCall.Returns.Output, r.ChangeResourceRecordSetsCall.Returns.Error
}
