package fakes

import (
	"sync"

	awsroute53 "github.com/aws/aws-sdk-go/service/route53"
)

type HealthChecksClient struct {
	DeleteHealthCheckCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DeleteHealthCheckInput *awsroute53.DeleteHealthCheckInput
		}
		Returns struct {
			DeleteHealthCheckOutput *awsroute53.DeleteHealthCheckOutput
			Error                   error
		}
		Stub func(*awsroute53.DeleteHealthCheckInput) (*awsroute53.DeleteHealthCheckOutput, error)
	}
	ListHealthChecksCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			ListHealthChecksInput *awsroute53.ListHealthChecksInput
		}
		Returns struct {
			ListHealthChecksOutput *awsroute53.ListHealthChecksOutput
			Error                  error
		}
		Stub func(*awsroute53.ListHealthChecksInput) (*awsroute53.ListHealthChecksOutput, error)
	}
}

func (f *HealthChecksClient) DeleteHealthCheck(param1 *awsroute53.DeleteHealthCheckInput) (*awsroute53.DeleteHealthCheckOutput, error) {
	f.DeleteHealthCheckCall.Lock()
	defer f.DeleteHealthCheckCall.Unlock()
	f.DeleteHealthCheckCall.CallCount++
	f.DeleteHealthCheckCall.Receives.DeleteHealthCheckInput = param1
	if f.DeleteHealthCheckCall.Stub != nil {
		return f.DeleteHealthCheckCall.Stub(param1)
	}
	return f.DeleteHealthCheckCall.Returns.DeleteHealthCheckOutput, f.DeleteHealthCheckCall.Returns.Error
}
func (f *HealthChecksClient) ListHealthChecks(param1 *awsroute53.ListHealthChecksInput) (*awsroute53.ListHealthChecksOutput, error) {
	f.ListHealthChecksCall.Lock()
	defer f.ListHealthChecksCall.Unlock()
	f.ListHealthChecksCall.CallCount++
	f.ListHealthChecksCall.Receives.ListHealthChecksInput = param1
	if f.ListHealthChecksCall.Stub != nil {
		return f.ListHealthChecksCall.Stub(param1)
	}
	return f.ListHealthChecksCall.Returns.ListHealthChecksOutput, f.ListHealthChecksCall.Returns.Error
}
