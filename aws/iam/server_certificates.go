package iam

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
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

func (s ServerCertificates) List(filter string) (map[string]string, error) {
	delete := map[string]string{}

	certificates, err := s.client.ListServerCertificates(&awsiam.ListServerCertificatesInput{})
	if err != nil {
		return delete, fmt.Errorf("Listing server certificates: %s", err)
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

		delete[n] = ""
	}

	return delete, nil
}

func (s ServerCertificates) Delete(serverCertificates map[string]string) error {
	for name, _ := range serverCertificates {
		_, err := s.client.DeleteServerCertificate(&awsiam.DeleteServerCertificateInput{
			ServerCertificateName: aws.String(name),
		})

		if err == nil {
			s.logger.Printf("SUCCESS deleting server certificate %s\n", name)
		} else {
			s.logger.Printf("ERROR deleting server certificate %s: %s\n", name, err)
		}
	}

	return nil
}
