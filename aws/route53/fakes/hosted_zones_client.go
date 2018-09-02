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
