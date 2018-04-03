package fakes

import (
	awsroute53 "github.com/aws/aws-sdk-go/service/route53"
)

type HealthChecksClient struct {
	ListHealthChecksCall struct {
		CallCount int
		Receives  struct {
			Input *awsroute53.ListHealthChecksInput
		}
		Returns struct {
			Output *awsroute53.ListHealthChecksOutput
			Error  error
		}
	}

	DeleteHealthCheckCall struct {
		CallCount int
		Receives  struct {
			Input *awsroute53.DeleteHealthCheckInput
		}
		Returns struct {
			Output *awsroute53.DeleteHealthCheckOutput
			Error  error
		}
	}
}

func (h *HealthChecksClient) ListHealthChecks(input *awsroute53.ListHealthChecksInput) (*awsroute53.ListHealthChecksOutput, error) {
	h.ListHealthChecksCall.CallCount++
	h.ListHealthChecksCall.Receives.Input = input

	return h.ListHealthChecksCall.Returns.Output, h.ListHealthChecksCall.Returns.Error
}

func (h *HealthChecksClient) DeleteHealthCheck(input *awsroute53.DeleteHealthCheckInput) (*awsroute53.DeleteHealthCheckOutput, error) {
	h.DeleteHealthCheckCall.CallCount++
	h.DeleteHealthCheckCall.Receives.Input = input

	return h.DeleteHealthCheckCall.Returns.Output, h.DeleteHealthCheckCall.Returns.Error
}
