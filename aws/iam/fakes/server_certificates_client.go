package fakes

import "github.com/aws/aws-sdk-go/service/iam"

type ServerCertificatesClient struct {
	ListServerCertificatesCall struct {
		CallCount int
		Receives  struct {
			Input *iam.ListServerCertificatesInput
		}
		Returns struct {
			Output *iam.ListServerCertificatesOutput
			Error  error
		}
	}

	DeleteServerCertificateCall struct {
		CallCount int
		Receives  struct {
			Input *iam.DeleteServerCertificateInput
		}
		Returns struct {
			Output *iam.DeleteServerCertificateOutput
			Error  error
		}
	}
}

func (i *ServerCertificatesClient) ListServerCertificates(input *iam.ListServerCertificatesInput) (*iam.ListServerCertificatesOutput, error) {
	i.ListServerCertificatesCall.CallCount++
	i.ListServerCertificatesCall.Receives.Input = input

	return i.ListServerCertificatesCall.Returns.Output, i.ListServerCertificatesCall.Returns.Error
}

func (i *ServerCertificatesClient) DeleteServerCertificate(input *iam.DeleteServerCertificateInput) (*iam.DeleteServerCertificateOutput, error) {
	i.DeleteServerCertificateCall.CallCount++
	i.DeleteServerCertificateCall.Receives.Input = input

	return i.DeleteServerCertificateCall.Returns.Output, i.DeleteServerCertificateCall.Returns.Error
}
