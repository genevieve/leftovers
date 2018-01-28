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

func (s ServerCertificates) List(filter string) (map[string]string, error) {
	certificates, err := s.list(filter)
	if err != nil {
		return nil, err
	}

	delete := map[string]string{}
	for _, c := range certificates {
		delete[*c.name] = ""
	}

	return delete, nil
}

func (s ServerCertificates) list(filter string) ([]ServerCertificate, error) {
	certificates, err := s.client.ListServerCertificates(&awsiam.ListServerCertificatesInput{})
	if err != nil {
		return nil, fmt.Errorf("Listing server certificates: %s", err)
	}

	var resources []ServerCertificate
	for _, c := range certificates.ServerCertificateMetadataList {
		resource := NewServerCertificate(s.client, c.ServerCertificateName)

		if !strings.Contains(resource.identifier, filter) {
			continue
		}

		proceed := s.logger.Prompt(fmt.Sprintf("Are you sure you want to delete server certificate %s?", resource.identifier))
		if !proceed {
			continue
		}

		resources = append(resources, resource)
	}

	return resources, nil
}

func (s ServerCertificates) Delete(serverCertificates map[string]string) error {
	var resources []ServerCertificate
	for name, _ := range serverCertificates {
		resources = append(resources, NewServerCertificate(s.client, &name))
	}

	return s.cleanup(resources)
}

func (s ServerCertificates) cleanup(resources []ServerCertificate) error {
	for _, resource := range resources {
		err := resource.Delete()

		if err == nil {
			s.logger.Printf("SUCCESS deleting server certificate %s\n", resource.identifier)
		} else {
			s.logger.Printf("ERROR deleting server certificate %s: %s\n", resource.identifier, err)
		}
	}

	return nil
}
