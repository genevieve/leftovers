package fakes

import gcpcompute "google.golang.org/api/compute/v1"

type SslCertificatesClient struct {
	ListSslCertificatesCall struct {
		CallCount int
		Returns   struct {
			Output []*gcpcompute.SslCertificate
			Error  error
		}
	}

	DeleteSslCertificateCall struct {
		CallCount int
		Receives  struct {
			SslCertificate string
		}
		Returns struct {
			Error error
		}
	}
}

func (n *SslCertificatesClient) ListSslCertificates() ([]*gcpcompute.SslCertificate, error) {
	n.ListSslCertificatesCall.CallCount++

	return n.ListSslCertificatesCall.Returns.Output, n.ListSslCertificatesCall.Returns.Error
}

func (n *SslCertificatesClient) DeleteSslCertificate(sslCertificate string) error {
	n.DeleteSslCertificateCall.CallCount++
	n.DeleteSslCertificateCall.Receives.SslCertificate = sslCertificate

	return n.DeleteSslCertificateCall.Returns.Error
}
