package awsiam

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/iam"
)

type ServerCertificates struct {
	client iamClient
}

func NewServerCertificates(client iamClient) ServerCertificates {
	return ServerCertificates{
		client: client,
	}
}

func (s ServerCertificates) Delete() {
	certificates, err := s.client.ListServerCertificates(&iam.ListServerCertificatesInput{})
	if err != nil {
		fmt.Errorf("ERROR listing server certificates: %s", err)
	}

	for _, c := range certificates.ServerCertificateMetadataList {
		n := c.ServerCertificateName
		_, err := s.client.DeleteServerCertificate(&iam.DeleteServerCertificateInput{ServerCertificateName: n})
		if err == nil {
			fmt.Printf("SUCCESS deleting server certificate %s\n", *n)
		} else {
			fmt.Printf("ERROR deleting server certificate %s: %s\n", *n, err)
		}
	}
}
