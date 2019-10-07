package fakes

import (
	"sync"

	awsiam "github.com/aws/aws-sdk-go/service/iam"
)

type ServerCertificatesClient struct {
	DeleteServerCertificateCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			DeleteServerCertificateInput *awsiam.DeleteServerCertificateInput
		}
		Returns struct {
			DeleteServerCertificateOutput *awsiam.DeleteServerCertificateOutput
			Error                         error
		}
		Stub func(*awsiam.DeleteServerCertificateInput) (*awsiam.DeleteServerCertificateOutput, error)
	}
	ListServerCertificatesCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			ListServerCertificatesInput *awsiam.ListServerCertificatesInput
		}
		Returns struct {
			ListServerCertificatesOutput *awsiam.ListServerCertificatesOutput
			Error                        error
		}
		Stub func(*awsiam.ListServerCertificatesInput) (*awsiam.ListServerCertificatesOutput, error)
	}
}

func (f *ServerCertificatesClient) DeleteServerCertificate(param1 *awsiam.DeleteServerCertificateInput) (*awsiam.DeleteServerCertificateOutput, error) {
	f.DeleteServerCertificateCall.Lock()
	defer f.DeleteServerCertificateCall.Unlock()
	f.DeleteServerCertificateCall.CallCount++
	f.DeleteServerCertificateCall.Receives.DeleteServerCertificateInput = param1
	if f.DeleteServerCertificateCall.Stub != nil {
		return f.DeleteServerCertificateCall.Stub(param1)
	}
	return f.DeleteServerCertificateCall.Returns.DeleteServerCertificateOutput, f.DeleteServerCertificateCall.Returns.Error
}
func (f *ServerCertificatesClient) ListServerCertificates(param1 *awsiam.ListServerCertificatesInput) (*awsiam.ListServerCertificatesOutput, error) {
	f.ListServerCertificatesCall.Lock()
	defer f.ListServerCertificatesCall.Unlock()
	f.ListServerCertificatesCall.CallCount++
	f.ListServerCertificatesCall.Receives.ListServerCertificatesInput = param1
	if f.ListServerCertificatesCall.Stub != nil {
		return f.ListServerCertificatesCall.Stub(param1)
	}
	return f.ListServerCertificatesCall.Returns.ListServerCertificatesOutput, f.ListServerCertificatesCall.Returns.Error
}
