package fakes

import (
	"sync"

	gcpcompute "google.golang.org/api/compute/v1"
)

type SslCertificatesClient struct {
	DeleteSslCertificateCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Certificate string
		}
		Returns struct {
			Error error
		}
		Stub func(string) error
	}
	ListSslCertificatesCall struct {
		sync.Mutex
		CallCount int
		Returns   struct {
			SslCertificateSlice []*gcpcompute.SslCertificate
			Error               error
		}
		Stub func() ([]*gcpcompute.SslCertificate, error)
	}
}

func (f *SslCertificatesClient) DeleteSslCertificate(param1 string) error {
	f.DeleteSslCertificateCall.Lock()
	defer f.DeleteSslCertificateCall.Unlock()
	f.DeleteSslCertificateCall.CallCount++
	f.DeleteSslCertificateCall.Receives.Certificate = param1
	if f.DeleteSslCertificateCall.Stub != nil {
		return f.DeleteSslCertificateCall.Stub(param1)
	}
	return f.DeleteSslCertificateCall.Returns.Error
}
func (f *SslCertificatesClient) ListSslCertificates() ([]*gcpcompute.SslCertificate, error) {
	f.ListSslCertificatesCall.Lock()
	defer f.ListSslCertificatesCall.Unlock()
	f.ListSslCertificatesCall.CallCount++
	if f.ListSslCertificatesCall.Stub != nil {
		return f.ListSslCertificatesCall.Stub()
	}
	return f.ListSslCertificatesCall.Returns.SslCertificateSlice, f.ListSslCertificatesCall.Returns.Error
}
