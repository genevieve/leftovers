package compute

import (
	"fmt"

	"github.com/genevieve/leftovers/common"
	gcpcompute "google.golang.org/api/compute/v1"
)

//go:generate faux --interface sslCertificatesClient --output fakes/ssl_certificates_client.go
type sslCertificatesClient interface {
	ListSslCertificates() ([]*gcpcompute.SslCertificate, error)
	DeleteSslCertificate(certificate string) error
}

type SslCertificates struct {
	client sslCertificatesClient
	logger logger
}

func NewSslCertificates(client sslCertificatesClient, logger logger) SslCertificates {
	return SslCertificates{
		client: client,
		logger: logger,
	}
}

func (s SslCertificates) List(filter string, regex bool) ([]common.Deletable, error) {
	s.logger.Debugln("Listing SSL Certificates...")
	sslCertificates, err := s.client.ListSslCertificates()
	if err != nil {
		return nil, fmt.Errorf("List Ssl Certificates: %s", err)
	}

	var resources []common.Deletable
	for _, cert := range sslCertificates {
		resource := NewSslCertificate(s.client, cert.Name)

		if !common.ResourceMatches(resource.Name(), filter, regex) {
			continue
		}

		proceed := s.logger.PromptWithDetails(resource.Type(), resource.Name())
		if !proceed {
			continue
		}

		resources = append(resources, resource)
	}

	return resources, nil
}

func (s SslCertificates) Type() string {
	return "compute-ssl-certificate"
}
