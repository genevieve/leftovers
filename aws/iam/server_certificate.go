package iam

import awsiam "github.com/aws/aws-sdk-go/service/iam"

type ServerCertificate struct {
	client     serverCertificatesClient
	name       *string
	identifier string
}

func NewServerCertificate(client serverCertificatesClient, name *string) ServerCertificate {
	return ServerCertificate{
		client:     client,
		name:       name,
		identifier: *name,
	}
}

func (s ServerCertificate) Delete() error {
	_, err := s.client.DeleteServerCertificate(&awsiam.DeleteServerCertificateInput{
		ServerCertificateName: s.name,
	})
	return err
}
