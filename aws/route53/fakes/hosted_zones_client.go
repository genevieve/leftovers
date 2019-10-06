package fakes

import (
	"sync"

	awsroute53 "github.com/aws/aws-sdk-go/service/route53"
)

type HostedZonesClient struct {
	DeleteHostedZoneCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DeleteHostedZoneInput *awsroute53.DeleteHostedZoneInput
		}
		Returns struct {
			DeleteHostedZoneOutput *awsroute53.DeleteHostedZoneOutput
			Error                  error
		}
		Stub func(*awsroute53.DeleteHostedZoneInput) (*awsroute53.DeleteHostedZoneOutput, error)
	}
	ListHostedZonesCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			ListHostedZonesInput *awsroute53.ListHostedZonesInput
		}
		Returns struct {
			ListHostedZonesOutput *awsroute53.ListHostedZonesOutput
			Error                 error
		}
		Stub func(*awsroute53.ListHostedZonesInput) (*awsroute53.ListHostedZonesOutput, error)
	}
}

func (f *HostedZonesClient) DeleteHostedZone(param1 *awsroute53.DeleteHostedZoneInput) (*awsroute53.DeleteHostedZoneOutput, error) {
	f.DeleteHostedZoneCall.Lock()
	defer f.DeleteHostedZoneCall.Unlock()
	f.DeleteHostedZoneCall.CallCount++
	f.DeleteHostedZoneCall.Receives.DeleteHostedZoneInput = param1
	if f.DeleteHostedZoneCall.Stub != nil {
		return f.DeleteHostedZoneCall.Stub(param1)
	}
	return f.DeleteHostedZoneCall.Returns.DeleteHostedZoneOutput, f.DeleteHostedZoneCall.Returns.Error
}
func (f *HostedZonesClient) ListHostedZones(param1 *awsroute53.ListHostedZonesInput) (*awsroute53.ListHostedZonesOutput, error) {
	f.ListHostedZonesCall.Lock()
	defer f.ListHostedZonesCall.Unlock()
	f.ListHostedZonesCall.CallCount++
	f.ListHostedZonesCall.Receives.ListHostedZonesInput = param1
	if f.ListHostedZonesCall.Stub != nil {
		return f.ListHostedZonesCall.Stub(param1)
	}
	return f.ListHostedZonesCall.Returns.ListHostedZonesOutput, f.ListHostedZonesCall.Returns.Error
}
