package iam

import (
	"fmt"
	"strings"

	awsiam "github.com/aws/aws-sdk-go/service/iam"
)

type serverCertificatesClient interface {
	ListServerCertificates(*awsiam.ListServerCertificatesInput) (*awsiam.ListServerCertificatesOutput, error)
	DeleteServerCertificate(*awsiam.DeleteServerCertificateInput) (*awsiam.DeleteServerCertificateOutput, error)
}

type ServerCertificates struct {
	client serverCertificatesClient
	logger logger
}

func NewServerCertificates(client serverCertificatesClient, logger logger) ServerCertificates {
	return ServerCertificates{
		client: client,
		logger: logger,
	}
}

func (s ServerCertificates) Delete(filter string) error {
	certificates, err := s.client.ListServerCertificates(&awsiam.ListServerCertificatesInput{})
	if err != nil {
		return fmt.Errorf("Listing server certificates: %s", err)
	}

	for _, c := range certificates.ServerCertificateMetadataList {
		n := *c.ServerCertificateName

		if !strings.Contains(n, filter) {
			continue
		}

		proceed := s.logger.Prompt(fmt.Sprintf("Are you sure you want to delete server certificate %s?", n))
		if !proceed {
			continue
		}

		_, err := s.client.DeleteServerCertificate(&awsiam.DeleteServerCertificateInput{ServerCertificateName: c.ServerCertificateName})
		if err == nil {
			s.logger.Printf("SUCCESS deleting server certificate %s\n", n)
		} else {
			s.logger.Printf("ERROR deleting server certificate %s: %s\n", n, err)
		}
	}

	return nil
}
