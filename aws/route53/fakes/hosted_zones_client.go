package fakes

import (
	awsroute53 "github.com/aws/aws-sdk-go/service/route53"
)

type HostedZonesClient struct {
	ListHostedZonesCall struct {
		CallCount int
		Receives  struct {
			Input *awsroute53.ListHostedZonesInput
		}
		Returns struct {
			Output *awsroute53.ListHostedZonesOutput
			Error  error
		}
	}

	DeleteHostedZoneCall struct {
		CallCount int
		Receives  struct {
			Input *awsroute53.DeleteHostedZoneInput
		}
		Returns struct {
			Output *awsroute53.DeleteHostedZoneOutput
			Error  error
		}
	}

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

func (h *HostedZonesClient) ListHostedZones(input *awsroute53.ListHostedZonesInput) (*awsroute53.ListHostedZonesOutput, error) {
	h.ListHostedZonesCall.CallCount++
	h.ListHostedZonesCall.Receives.Input = input

	return h.ListHostedZonesCall.Returns.Output, h.ListHostedZonesCall.Returns.Error
}

func (h *HostedZonesClient) DeleteHostedZone(input *awsroute53.DeleteHostedZoneInput) (*awsroute53.DeleteHostedZoneOutput, error) {
	h.DeleteHostedZoneCall.CallCount++
	h.DeleteHostedZoneCall.Receives.Input = input

	return h.DeleteHostedZoneCall.Returns.Output, h.DeleteHostedZoneCall.Returns.Error
}

func (h *HostedZonesClient) ListResourceRecordSets(input *awsroute53.ListResourceRecordSetsInput) (*awsroute53.ListResourceRecordSetsOutput, error) {
	h.ListResourceRecordSetsCall.CallCount++
	h.ListResourceRecordSetsCall.Receives = append(h.ListResourceRecordSetsCall.Receives, ListResourceRecordSetsCallReceive{Input: input})

	if len(h.ListResourceRecordSetsCall.Returns) < h.ListResourceRecordSetsCall.CallCount {
		return nil, nil
	}

	return h.ListResourceRecordSetsCall.Returns[h.ListResourceRecordSetsCall.CallCount-1].Output, h.ListResourceRecordSetsCall.Returns[h.ListResourceRecordSetsCall.CallCount-1].Error
}

func (h *HostedZonesClient) ChangeResourceRecordSets(input *awsroute53.ChangeResourceRecordSetsInput) (*awsroute53.ChangeResourceRecordSetsOutput, error) {
	h.ChangeResourceRecordSetsCall.CallCount++
	h.ChangeResourceRecordSetsCall.Receives.Input = input

	return h.ChangeResourceRecordSetsCall.Returns.Output, h.ChangeResourceRecordSetsCall.Returns.Error
}
