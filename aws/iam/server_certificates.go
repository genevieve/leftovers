package iam

import (
	"fmt"

	awsiam "github.com/aws/aws-sdk-go/service/iam"
)

type ServerCertificates struct {
	client iamClient
	logger logger
}

func NewServerCertificates(client iamClient, logger logger) ServerCertificates {
	return ServerCertificates{
		client: client,
		logger: logger,
	}
}

func (s ServerCertificates) Delete() error {
	certificates, err := s.client.ListServerCertificates(&awsiam.ListServerCertificatesInput{})
	if err != nil {
		return fmt.Errorf("Listing server certificates: %s", err)
	}

	for _, c := range certificates.ServerCertificateMetadataList {
		n := c.ServerCertificateName

		proceed := s.logger.Prompt(fmt.Sprintf("Are you sure you want to delete server certificate %s?", *n))
		if !proceed {
			continue
		}

		_, err := s.client.DeleteServerCertificate(&awsiam.DeleteServerCertificateInput{ServerCertificateName: n})
		if err == nil {
			s.logger.Printf("SUCCESS deleting server certificate %s\n", *n)
		} else {
			s.logger.Printf("ERROR deleting server certificate %s: %s\n", *n, err)
		}
	}

	return nil
}
